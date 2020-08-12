// Copyright 2020 The Docker Applications Manager Authors
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
package filesystem

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"dam/driver/logger"
)

// COPY FROM: https://medium.com/@skdomino/taring-untaring-files-in-go-6b07cf56bc07
func Untar(source, dst string) string {
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
			logger.Fatal("Cannot find file %s in tar with error: %s", dst, err.Error())
			return dst

		case header == nil:
			continue
		}

		target := filepath.Join(dst, header.Name)

		switch header.Typeflag {

		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					logger.Fatal("Cannot create directory %s with error: %s", target, err.Error())
					return dst
				}
			}

		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				logger.Fatal("Cannot open file %s with error: %s", target, err.Error())
				return dst
			}

			if _, err := io.Copy(f, tr); err != nil {
				logger.Fatal("Cannot copy file: %s with error: %s", target, err.Error())
				return dst
			}

			err = f.Sync()
			if err != nil {
				logger.Fatal("Cannot sync file '%s' with error: %s", target, err)
			}
			err = f.Close()
			if err != nil {
				logger.Fatal("Cannot close file '%s' with error: %s", target, err)
			}
		}
	}
}


func Gzip(source, target string) {
	f, err := os.Create(target)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot create file '%s' with error: %s", target, err)
	}

	gw := gzip.NewWriter(f)
	defer func() {
		if gw != nil {
			gw.Close()
		}
	}()

	tw := tar.NewWriter(gw)
	defer func() {
		if tw != nil {
			tw.Close()
		}
	}()

	walkFn := func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsDir() {
			return nil
		}

		newPath := path[len(source):]
		if len(newPath) == 0 {
			return nil
		}
		pathFile, err := os.Open(path)
		defer func() {
			if pathFile != nil {
				pathFile.Close()
			}
		}()
		if err != nil {
			return err
		}

		if h, err := tar.FileInfoHeader(info, newPath); err != nil {
			logger.Fatal("Cannot create file info header for '%s' with error: %s", newPath, err)
		} else {
			h.Name = newPath
			if err = tw.WriteHeader(h); err != nil {
				logger.Fatal("Cannot write header for '%s' with error: %s", newPath, err)
			}
		}
		if _ , err := io.Copy(tw, pathFile); err != nil {
			logger.Fatal("Cannot write file '%s' to archive with error: %s", pathFile, err)
		}

		return nil
	}

	if err = filepath.Walk(source, walkFn); err != nil {
		fmt.Println(err)
	}


}