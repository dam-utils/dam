package filesystem

import (
	"archive/tar"
	"compress/gzip"
	"dam/driver/logger"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Untar(source string) string {
	dst := source + "_tar"

	f, err := os.Open(source)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open tar file '%s' with error: %s", source, err)
	}

	tr := tar.NewReader(f)

	for {
		header, err := tr.Next()

		switch {
		case err == io.EOF:
			return dst
		case err != nil:
			logger.Fatal("Cannot find file %s in tar with error: %s", dst, err)
			return dst
		case header == nil:
			continue
		}

		target := filepath.Join(dst, header.Name)
		MkDir(filepath.Dir(target))

		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					logger.Fatal("Cannot create directory %s with error: %s", target, err)
					return dst
				}
			}
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			defer func() {
				if f != nil {
					f.Close()
				}
			}()
			if err != nil {
				logger.Fatal("Cannot open file %s with error: %s", target, err)
				return dst
			}
			if _, err := io.Copy(f, tr); err != nil {
				logger.Fatal("Cannot copy file: %s with error: %s", target, err)
				return dst
			}
			err = f.Sync()
			if err != nil {
				logger.Fatal("Cannot sync file '%s' with error: %s", target, err)
			}
		case tar.TypeSymlink:
			err = os.Symlink(header.Linkname, target)
			if err != nil {
				logger.Fatal("Cannot create symlink '%s' with old name '%s' with error: %s", target, header.Linkname, err)
			}
		}
	}
}

func Gunzip(source string) string {
	target := source + "_zip"

	f, err := os.Open(source)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open gzip file '%s' with error: %s", source, err)
	}

	gr, err := gzip.NewReader(f)
	defer func() {
		if gr != nil {
			gr.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot create gzip reader for file '%s' with error: %s", source, err)
	}

	target = filepath.Join(target, gr.Name)
	writer, err := os.Create(target)
	if err != nil {
		logger.Fatal("Cannot read archive '%s' with error: '%s'", source, err)
	}
	defer writer.Close()

	_, err = io.Copy(writer, gr)
	if err != nil {
		logger.Fatal("Cannot copy archive from '%s' to '%s' with error: '%s'", source, target, err)
	}

	return target
}

func Gzip(source, target string, onlyTar bool) {
	f, err := os.Create(target)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot create file '%s' with error: %s", target, err)
	}

	var tw *tar.Writer
	if !onlyTar {
		gw := gzip.NewWriter(f)
		defer func() {
			if gw != nil {
				gw.Close()
			}
		}()

		tw = tar.NewWriter(gw)
		defer func() {
			if tw != nil {
				tw.Close()
			}
		}()
	} else {
		tw = tar.NewWriter(f)
		defer func() {
			if tw != nil {
				tw.Close()
			}
		}()
	}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		var link string
		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			if link, err = os.Readlink(path); err != nil {
				return err
			}
		}

		header, err := tar.FileInfoHeader(info, link)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, source)
		if err = tw.WriteHeader(header); err != nil {
			return err
		}

		if !info.Mode().IsRegular() { //nothing more to do for non-regular
			return nil
		}

		fh, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fh.Close()

		if _, err = io.Copy(tw, fh); err != nil {
			return err
		}
		return nil
	}

	if err = filepath.Walk(source, walkFn); err != nil {
		logger.Fatal("Cannot create archive: %v", err)
	}
}
