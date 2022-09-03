package api

type Config struct {
	BindAddr string `toml:"bind_addr"`
	DatabaseURL string `toml:"database_url"`
	LogFile string `toml:"log_file"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config {
		BindAddr: ":8080",
	}
}