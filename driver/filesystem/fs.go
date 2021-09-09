package filesystem

import (
	"archive/tar"
	"bytes"
	"encoding/hex"
	"hash/crc32"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"dam/driver/conf/option"
	"dam/driver/logger"
)

func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		logger.Fatal("Cannot get current dir with error: %s", err)
	}
	return dir
}

func IsExistDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		logger.Debug("Cannot check directory '%s' with error: %s", path, err)
		return false
	}
	return info.IsDir()
}

func IsExistFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func Ls(dir string) []string {
	files, err := filepath.Glob(dir + string(os.PathSeparator) + "*")
	if err != nil {
		logger.Fatal("Cannot check files in path '%s' with error: %s", dir, err)
	}
	return files
}

func MkDir(dir string) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		logger.Fatal("Cannot create directory '%s/' with error: %s", dir, err)
	}
}

func GetDir(filePath string) string {
	return filepath.Dir(filePath)
}

func Remove(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		logger.Warn("Cannot check path '%s' with error: %s", path, err)
	}

	err = os.RemoveAll(path)
	if err != nil {
		logger.Warn("Cannot remove path '%s' with error: %s", path, err)
	}
	return false
}

func MoveFile(oldLocation, newLocation string) {
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		logger.Fatal("Cannot move file '%s' to '%s' with error: %s", oldLocation, newLocation, err)
	}
}

func GetAbsolutePath(path string) string {
	p, err := filepath.Abs(path)
	{
		if err != nil {
			logger.Fatal("Cannot get absolute path for '%s' with error: %s", path, err)
		}
	}
	return p
}

// TODO заменить на ChmodPlusX()
func Chmod777(path string) {
	if err := os.Chmod(path, 0777); err != nil {
		logger.Fatal("Cannot chmod 777 '%s' with error: %s", path, err)
	}
}

func Chdir(path string) {
	err := os.Chdir(path)
	if err != nil {
		logger.Fatal("Cannot change home dir to '%s' with error: %s", path, err)
	}
}

// https://stackoverflow.com/questions/40670228/how-to-run-binary-files-inside-golang-program
func RunFile(runFile string) {
	pwd := GetCurrentDir()
	defer Chdir(pwd)

	homeDir := filepath.Dir(runFile)
	Chdir(homeDir)

	c := exec.Command(runFile)
	c.Dir = homeDir   //TODO delete?
	// set var to get the output
	var outb, errb bytes.Buffer

	// set the output to our variable
	c.Stdout = &outb
	c.Stderr = &errb
	err := c.Run()
	if err != nil {
		logger.Warn(errb.String())
		logger.Fatal("Cannot execute file '%s' with error: %s", runFile, err)
	}
	logger.Info(outb.String())
}

func Touch(file string) {
	if !IsExistFile(file) {
		emptyFile, err := os.Create(file)
		defer func() {
			if emptyFile != nil {
				emptyFile.Close()
			}
		}()
		if err != nil {
			logger.Fatal("Cannot create file '%s' with error: %s", file, err)
		}
	}
}

func HashFileCRC32(filePath string) string {
	f, err := os.Open(filePath)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open file '%s' with error: %s", filePath, err)
	}

	tablePolynomial := crc32.MakeTable(option.Config.Save.GetPolynomialCksum())
	hash := crc32.New(tablePolynomial)
	if _, err := io.Copy(hash, f); err != nil {
		logger.Fatal("Cannot check hash file '%s' with error: %s", filePath, err)
	}
	hashInBytes := hash.Sum(nil)[:]
	return hex.EncodeToString(hashInBytes)
}

func IsTar(path string) bool {
	f, err := os.Open(path)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open file '%s' with error: %s", path, err)
	}

	tr := tar.NewReader(f)
	_, err = tr.Next()

	return err == nil
}