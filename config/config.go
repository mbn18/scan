package config

import "fmt"

type Config struct {
	Address string
	DB      struct {
		name, user, pass, address string
	}
}

func (c Config) GetConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.DB.user, c.DB.pass, c.DB.address, c.DB.name)
}

// Get @TODO data should be loaded from environment. perhaps flags.
func Get() Config {
	return Config{
		Address: ":8080",
		DB: struct{ name, user, pass, address string }{
			name:    "cypago",
			user:    "cypago_user",
			pass:    "pass",
			address: "10.0.0.16",
		},
	}
}
