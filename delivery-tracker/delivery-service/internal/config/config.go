package config

type Config struct {
	KafkaBroker string 
	KafkaTopic  string
	ServerPort  string 
}

func Load() *Config {
	return &Config{
		KafkaBroker: "localhost:9092",
		KafkaTopic: "delivery-location",
		ServerPort: "8080",
	}
}