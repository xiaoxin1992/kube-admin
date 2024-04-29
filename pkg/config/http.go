package config

type http struct {
	Host   string `toml:"host"`
	Port   int    `toml:"port"`
	Logger logger `toml:"logger"`
}
