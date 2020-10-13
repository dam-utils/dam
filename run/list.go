package run

import (
	"dam/driver/decorate"
)

type ListSettings struct {
	Raw    bool
	//	Labels string
}

var ListFlags = new(ListSettings)

func List() {
	if ListFlags.Raw {
		decorate.PrintRAWAppsList()
	} else {
		decorate.PrintAppsList()
	}
}