package config

type logger struct {
	Path      string `toml:"path"`
	Format    string `toml:"format"`
	MaxSize   int    `toml:"maxSize"`
	MaxBackup int    `toml:"maxBackup"`
	MaxAge    int    `toml:"maxAge"`
	LocalTime bool   `toml:"localTime"`
	Compress  bool   `toml:"compress"`
	Console   bool   `toml:"console"`
	IsFile    bool   `toml:"IsFile"`
}
