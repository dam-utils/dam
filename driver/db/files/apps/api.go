package apps

import "os"

type provider struct {
	//GetApps() []*storage.App
	//NewApp(app *storage.App)
	//GetAppById(id int) *storage.App
	//ExistFamily(family string) bool
	//RemoveAppById(id int)

	client *os.File
}

func NewProvider() *provider {
	return &provider{}
}
