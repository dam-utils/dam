package docker

import (
	"context"

	"dam/driver/conf/option"
	"dam/driver/logger"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
)

func (p *provider) SearchAppNames(mask string) *[]string {
	if len(mask) <2 || len(mask)> 100 {
		logger.Fatal("Search mask for official registry not valid. It must be 2-24 symbols")
	}

	p.connect()
	defer p.close()

	searchOpts := types.ImageSearchOptions {}
	searchOpts.Limit = option.Config.Search.GetOfficialRepoAppsLimit()
	var results []registry.SearchResult
	results, err := p.client.ImageSearch(context.Background(), mask,  searchOpts)
	if err != nil {
		logger.Fatal("Cannot get results of docker search with error: %s", err)
	}

	var appNames []string
	for _, result := range results {
		appNames = append(appNames, result.Name)
	}
	return &appNames
}
