package classes

type Config struct {
	DefaultBalance (float64) 
}

func NewConfig() *Config {
	newConfig := Config {
		DefaultBalance: 1000,
	}
	return &newConfig
}