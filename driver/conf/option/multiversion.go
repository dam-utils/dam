package option

import "dam/config"

type Multiversion struct {}

func (o *Multiversion) GetTrueFlag() string {
	return config.MULTIVERSION_TRUE_FLAG
}

func (o *Multiversion) GetFalseFlag() string {
	return config.MULTIVERSION_FALSE_FLAG
}
