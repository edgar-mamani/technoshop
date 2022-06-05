package config

type Configuration interface {
	GetString(name string) (configVal string, found bool)
	GetInt(name string) (configVal int, found bool)
	GetFloat(name string) (configVal float64, found bool)
	GetBool(name string) (configVal bool, found bool)

	GetStringDefault(name string) (configVal string)
	GetIntDefault(name string) (configVal int)
	GetFloatDefault(name string) (configVal float64)
	GetBoolDefault(name string) (configVal bool)

	GetSection(sectionName string) (section Configuration, found bool)
}
