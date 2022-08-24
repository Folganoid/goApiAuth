package api

type Config struct {
	BindAddr string `toml:"bind_addr"`
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config {
		BindAddr: ":8080",
	}
}