package meta

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/filesystem/env"
	"dam/driver/logger"
)

func PrepareExpFiles(metaDir string, envs map[string]string) {
	files := make([]string, 0)
	err := filepath.Walk(metaDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if fs.IsExistFile(path) {
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		logger.Fatal("Cannot walk in meta directory with error: %s", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file, config.EXPAND_META_FILE){
			prepareExpFile(file, envs)
		}
	}
}

func prepareExpFile(path string, envs map[string]string) {
	newPath := strings.TrimSuffix(path, config.EXPAND_META_FILE)

	f, err := os.Open(path)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open file '%s' with error: %s", path, err)
	}

	newf, err := os.Create(newPath)
	defer func() {
		if newf != nil {
			f.Close()
		}
	}()
	if err != nil {
		newf.Close()
		logger.Fatal("Cannot create file '%s' with error: %s", newPath, err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		newString := env.PrepareExpString(scanner.Text(), envs)
		_, err = newf.WriteString(newString+"\n")
		if err != nil {
			logger.Fatal("Cannot write string to file '%s' with error: %s", newPath, err)
		}
	}

	err = newf.Sync()
	if err != nil {
		logger.Fatal("Cannot sync of writable file '%s' with error: %s", newPath, err)
	}

	err = newf.Close()
	if err != nil {
		logger.Fatal("Cannot close of writable file '%s' with error: %s", newPath, err)
	}
}

func PrepareExecFiles(meta string) {
	fs.Chmod777(filepath.Join(meta, config.INSTALL_FILE_NAME))
	fs.Chmod777(filepath.Join(meta, config.UNINSTALL_FILE_NAME))
}
