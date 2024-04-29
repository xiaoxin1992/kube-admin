package config

type MySQL struct {
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	Database    string `toml:"database"`
	User        string `toml:"user"`
	Password    string `toml:"password"`
	MaxLifetime int    `toml:"maxLifeTime"`
	MaxOpen     int    `toml:"maxOpen"`
	MaxIdle     int    `toml:"maxIdle"`
}
