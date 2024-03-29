package run

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"

	"dam/driver/conf/option"
	"dam/driver/db"
	"dam/driver/decorate"
	"dam/driver/engine"
	fs "dam/driver/filesystem"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/driver/logger/color"
	"dam/driver/structures"
	"dam/run/internal"
	"dam/run/internal/archive/app_name"
)

type ImportSettings struct {
	Yes     bool
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

	if emptyAppLists(appSkipList, appInstallList, appDeleteList) {
		logger.Info("App list is empty.")
	} else {
		decorate.PrintAppList("Skip apps:\n", appSkipList, color.Yellow)
		decorate.PrintAppList("Install apps:\n", appInstallList, color.Green)
		decorate.PrintAppList("Delete apps:\n", appDeleteList, color.Red)
	}

	if !ImportFlags.Yes {
		answer := questionYesNo()
		if !answer {
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
		tag := internal.GetPrefixRepo() + a.CurrentName()
		id := engine.VDriver.GetImageID(tag)
		if id == "" {
			logger.Fatal("Image with tag '%s' not exist in the system", tag)
		}
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

	filesList := getAppFilesList(tmpDir)
	for _, appFile := range filesList {
		logger.Debug("Validating %s", appFile)
		validateCheckSumArch(appFile)

		engine.VDriver.LoadImage(appFile)

		logger.Debug("Preparing servers label ...")
		tag := getTagFromArchiveManifest(appFile)
		serversLabel := internal.GetServersByTag(tag)
		logger.Debug("Servers labels '%v' for tag '%s'", serversLabel, tag)
		createTagImages(tag, serversLabel)
	}

	return appsFromFile(path.Join(tmpDir, option.Config.Export.GetAppsFileName()))
}

func getAppFilesList(tmpDir string) []string {
	resultFiles := make([]string, 0)

	for _, f := range fs.Ls(tmpDir) {
		if strings.HasSuffix(f, option.Config.Save.GetFilePostfix()) {
			resultFiles = append(resultFiles, f)
		}
	}

	return resultFiles
}

func validateCheckSumArch(appFile string) {
	fileInfo := app_name.NewInfo()
	err := fileInfo.FromString(appFile)
	if err != nil {
		logger.Fatal("Cannot validate checksum of archive: parsing archive name '%s' error: %s", appFile, err)
	}

	fileHash := fs.HashFileCRC32(appFile)
	if fileInfo.Hash() != fileHash {
		logger.Fatal("File hash '%s' not equal hash '%s' file name '%s'", fileHash, fileInfo.Hash(), appFile)
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

	fileInfo, err := f.Stat()
	if err != nil {
		logger.Fatal("Cannot get stats file '%s' with error: %s", path, err)
	}

	if fileInfo.Size() == 0 {
		logger.Debug("File '%s' is empty", path)
		return result
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
		if !flagExist {
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
		if !flagExist {
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
	case 'Y', 'y':
		return true
	case 'N', 'n':
		return false
	default:
		logger.Warn("Unknown symbol '%v'", char)
	}

	return false
}

func emptyAppLists(appSkipList, appInstallList, appDeleteList []*structures.ImportApp) bool {
	result := make([]*structures.ImportApp, 0)
	result = append(result, appSkipList...)
	result = append(result, appInstallList...)
	result = append(result, appDeleteList...)
	return len(result) == 0
}