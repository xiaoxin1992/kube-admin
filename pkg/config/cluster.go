package config

type Cluster struct {
	Name  string `toml:"name"`
	Host  string `toml:"host"`
	Token string `toml:"token"`
}
