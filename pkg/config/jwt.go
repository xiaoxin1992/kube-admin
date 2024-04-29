package config

type jwt struct {
	Token      string `toml:"token"`
	ExpireTime int64  `toml:"expire_time"`
}
