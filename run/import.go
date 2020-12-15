package run

import (
	"bufio"
	"dam/driver/engine"
	"dam/run/internal"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"dam/config"
	"dam/driver/db"
	"dam/driver/decorate"
	fs "dam/driver/filesystem"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/driver/logger/color"
	"dam/driver/structures"
)

type ImportSettings struct {
	Yes bool
	Restore bool
}

var ImportFlags = new(ImportSettings)

func Import(arg string) {
	var appImportList []*structures.ImportApp

	var appSkipList []*structures.ImportApp
	var appDeleteList []*structures.ImportApp
	var appInstallList []*structures.ImportApp

	flag.ValidateFilePath(arg)
	logger.Debug("Flags validated with success")

	logger.Debug("Getting appAllList ...")
	appAllList := getListFromApps(db.ADriver.GetApps())

	if fs.IsTar(arg) {
		logger.Debug("Preparing app images from archive ...")
		appImportList = loadAppsFromArchive(arg)
	} else {
		logger.Debug("Getting the import list of apps ...")
		appImportList = appsFromFile(arg)
		loadAppFromRegistry(appImportList)
	}

	validateExistingImages(appImportList)

	logger.Debug("Preparing change list ...")
	if ImportFlags.Restore {
		appDeleteList = appAllList
		appInstallList = appImportList
	} else {
		appDeleteList, appInstallList, appSkipList = matchLists(appAllList, appImportList)
	}

	decorate.PrintAppList("Skip apps:\n", appSkipList, color.Yellow)
	decorate.PrintAppList("Install apps:\n", appInstallList, color.Green)
	decorate.PrintAppList("Delete apps:\n", appDeleteList, color.Red)

	if !ImportFlags.Yes {
		answer := questionYesNo()
		if answer == false {
			logger.Success("Stop import.")
			os.Exit(0)
		}
	}

	if len(appDeleteList) != 0 {
		for _, app := range appDeleteList {
			RemoveApp(app.Name)
		}
	}

	if len(appInstallList) != 0 {
		for _, app := range appInstallList {
			InstallApp(app.CurrentName())
		}
	}

	logger.Success("Import was successful.")
}

func validateExistingImages(apps []*structures.ImportApp) {
	for _, a := range apps {
		_ = engine.VDriver.GetImageID(internal.GetPrefixRepo()+a.CurrentName())
	}
}

func loadAppFromRegistry(apps []*structures.ImportApp) {
	repo := db.RDriver.GetDefaultRepo()
	for _, a := range apps {
		engine.VDriver.Pull(internal.GetPrefixRepo()+a.CurrentName(), repo)
	}
}

func loadAppsFromArchive(arch string) []*structures.ImportApp {
	tmpDir := fs.Untar(arch)
	defer fs.Remove(tmpDir)

	appList := getAppFilesList(tmpDir)
	for _, a := range appList {
		validateCheckSumArch(a)
		engine.VDriver.LoadImage(a)
	}

	apps := appsFromFile(tmpDir+string(filepath.Separator)+config.EXPORT_APPS_FILE_NAME)

	return apps
}

func getAppFilesList(tmpDir string) []string {
	resultFiles := make([]string,0)

	for _, f := range fs.Ls(tmpDir) {
		if strings.HasSuffix(f, config.SAVE_FILE_POSTFIX) {
			resultFiles = append(resultFiles, f)
		}
	}

	return resultFiles
}

func validateCheckSumArch(appFile string) {
	hash := fs.HashFileCRC32(appFile)
	size := fs.FileSize(appFile)

	result1 := strings.TrimRight(appFile, config.SAVE_FILE_POSTFIX)
	arrWithSize := strings.Split(result1, config.SAVE_FILE_SEPARATOR)
	fileSize := arrWithSize[len(arrWithSize)-1]
	if fileSize != size {
		logger.Fatal("File size '%s' not equal size '%s' in file name '%s'", size, fileSize, appFile)
	}
	result2 := strings.TrimRight(result1, fileSize)
	result3 := strings.TrimRight(result2, config.SAVE_FILE_SEPARATOR)
	arrWithHash := strings.Split(result3, config.SAVE_OPTIONAL_SEPARATOR)
	fileHash := arrWithHash[len(arrWithHash)-1]
	if fileHash != hash {
		logger.Fatal("File hash '%s' not equal hash '%s' file name '%s'", hash, fileHash, appFile)
	}
}

func getListFromApps(apps []*structures.App) []*structures.ImportApp {
	result := make([]*structures.ImportApp, 0)

	for _, a := range apps {
		result = append(result, &structures.ImportApp{
			Name:    a.ImageName,
			Version: a.ImageVersion,
		})
	}

	return result
}

func appsFromFile(path string) []*structures.ImportApp {
	result := make([]*structures.ImportApp, 0)

	f, err := os.Open(path)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open file '%s' with error: %s", path, err)
	}

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		newLine := fileScanner.Text()
		result = append(result, str2importApp(newLine))
	}

	return result
}

func str2importApp(str string) *structures.ImportApp {
	strArray := strings.Split(str, ":")
	if len(strArray) != 2 {
		logger.Fatal("Import file is bad. String '%s' not equal '<app>:<version>'", str)
	}

	return &structures.ImportApp{
		Name:    strArray[0],
		Version: strArray[1],
	}
}

func matchLists(allApps, importApps []*structures.ImportApp) (appDeleteList, appInstallList, appSkipList []*structures.ImportApp) {
	for _, iApp := range importApps {
		var flagExist = false
		for _, aApp := range allApps {
			if iApp.CurrentName() == aApp.CurrentName() {
				appSkipList = append(appSkipList, iApp)
				flagExist = true
			}
		}
		if flagExist == false {
			appInstallList = append(appInstallList, iApp)
		}
	}

	for _, aApp := range allApps {
		var flagExist = false
		for _, iApp := range importApps {
			if aApp.CurrentName() == iApp.CurrentName() {
				flagExist = true
			}
		}
		if flagExist == false {
			appDeleteList = append(appDeleteList, aApp)
		}
	}

	return
}

func questionYesNo() bool {
	fmt.Print("Are you sure you want to apply the changes to the system? [yY/nN]:")

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		logger.Fatal("Cannot read answer rune in Yes/No question.")
	}

	fmt.Println()

	switch char {
	case 'Y','y':
		return true
	case 'N', 'n':
		return false
	default:
		logger.Warn("Unknown symbol '%v'", char)
	}

	return false
}