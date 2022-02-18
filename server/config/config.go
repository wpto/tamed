package config

type Config struct {
	LocalPath  string
	MediaPath  string
	PostDBPath string
}

var config Config = Config{
	LocalPath:  "../local",
	MediaPath:  "../mediacontent",
	PostDBPath: "../post.json",
}

func Get() *Config {
	return &config
}
