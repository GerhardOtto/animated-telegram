package config

import "flag"

// Config holds all CLI-configured values for the application.
type Config struct {
	Addr       string
	Name       string
	Sort       string // default sort algorithm: merge|quick|insertion
	Order      string // default sort order: asc|desc
	CustomSort bool   // if true, prompt user to choose algorithm and order interactively
}

// Parse reads CLI flags and returns a populated Config.
func Parse() *Config {
	addr       := flag.String("addr", "localhost:50051", "the address to connect to")
	name       := flag.String("name", "arthur", "name to greet")
	sort       := flag.String("sort", "merge", "default sort algorithm (merge|quick|insertion)")
	order      := flag.String("order", "asc", "default sort order (asc|desc)")
	customSort := flag.Bool("custom-sort", false, "prompt to choose sort algorithm and order interactively")
	flag.Parse()

	return &Config{
		Addr:       *addr,
		Name:       *name,
		Sort:       *sort,
		Order:      *order,
		CustomSort: *customSort,
	}
}
