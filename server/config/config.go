package config

type Config struct {
	LocalPath string
}

var config Config = Config{
	LocalPath: "../local",
}

func Get() *Config {
	return &config
}
