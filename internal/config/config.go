package config

type Config struct {
	ShuffleWallets bool
	UseProxy       bool
	ProxyAddress   string
}

func NewConfig() *Config {
	return &Config{
		UseProxy:     true,                             // If you want to use proxy while checking
		ProxyAddress: "http://user:password@host:port", // Residential proxy
	}
}
