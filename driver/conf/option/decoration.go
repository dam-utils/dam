package option

import "dam/config"

type Decoration struct {}

func (o *Decoration) GetColorOn() bool {
	return config.COLOR_ON
}

func (o *Decoration) GetMaxDisplayWidth() int {
	return config.DECORATE_MAX_DISPLAY_WIDTH
}

func (o *Decoration) GetRawSeparator() string {
	return config.DECORATE_RAW_SEPARATOR
}

func (o *Decoration) GetBoolFlagSymbol() string {
	return config.DECORATE_BOOL_FLAG_SYMBOL
}

