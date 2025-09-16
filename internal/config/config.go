package config

type Config struct {
	UseProxy     bool
	ProxyAddress string
	WorkerAmount int
}

func NewConfig() *Config {
	return &Config{
		WorkerAmount: 5,                                // I recommend set up 5-10 workers
		UseProxy:     true,                             // If you want to use proxy while checking
		ProxyAddress: "http://user:password@host:port", // Residential proxy
	}
}
