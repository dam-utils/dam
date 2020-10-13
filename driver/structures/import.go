package structures

type ImportApp struct {
	Name    string
	Version string
}

func (a *ImportApp) CurrentName() string {
	return a.Name + ":" + a.Version
}