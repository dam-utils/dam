package registry

import (
	"dam/driver/engine"
	"dam/driver/structures"
	"strings"

	"dam/driver/conf/option"
	"dam/driver/logger"
	registry_official "dam/driver/registry/official"
	registry_v2 "dam/driver/registry/v2"
)

func CheckRepository(repo *structures.Repo) {
	if repo.Id ==1 {
		return
	}

	for _, protocol := range option.Config.Search.GetProtocolStrategy() {
		err := registry_v2.CheckRepo(repo, protocol)
		if err != nil {
			logger.Debug("Cannot connect to default registry '%s' for '%s' protocol with error: %s", repo.Name, protocol, err)
		} else {
			return
		}
	}
	logger.Fatal("Cannot connect to default registry '%s'", repo.Name)
}

func GetAppNamesByMask(repo *structures.Repo, mask string) *[]string {
	if repo.Id == 1 {
		return engine.VDriver.
			SearchAppNames(mask)
	}
	names := registry_v2.GetAppNames(repo)
	return filterAppNamesByMask(names, mask)
}

func filterAppNamesByMask(names *[]string, mask string) *[]string {
	var res []string
	for _, name := range *names {
		if strings.Contains(name, mask) {
			res = append(res, name)
		}
	}
	return &res
}

func GetVersions(repo *structures.Repo, appName string) *[]string {
	var versions *[]string
	if repo.Id == 1 {
		versions = registry_official.GetAppVersions(appName)
	} else {
		versions = registry_v2.GetAppVersions(repo, appName)
	}
	return versions
}