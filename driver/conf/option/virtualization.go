package option

import "dam/config"

type Virtualization struct {}

func (o *Virtualization) GetType() string {
	return config.VIRTUALIZATION_TYPE
}
