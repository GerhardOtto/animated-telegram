package config

import "flag"

// Config holds all CLI-configured values for the application.
type Config struct {
	Addr string
}

// Parse reads CLI flags and returns a populated Config.
func Parse() *Config {
	addr := flag.String("addr", "localhost:50051", "the address to connect to")
	flag.Parse()

	return &Config{
		Addr: *addr,
	}
}
