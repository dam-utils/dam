package option

import "dam/config"

type Save struct {}

func (o *Save) GetOptionalSeparator() string {
	return config.SAVE_OPTIONAL_SEPARATOR
}

func (o *Save) GetFileSeparator() string {
	return config.SAVE_FILE_SEPARATOR
}

func (o *Save) GetTmpFilePostfix() string {
	return config.SAVE_TMP_FILE_POSTFIX
}

func (o *Save) GetFilePostfix() string {
	return config.SAVE_FILE_POSTFIX
}

func (o *Save) GetPolynomialCksum() uint32 {
	return config.SAVE_POLYNOMIAL_CKSUM
}

func (o *Save) GetManifestFile() string {
	return config.SAVE_MANIFEST_FILE
}