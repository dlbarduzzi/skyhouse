package skyhouse

import "fmt"

type Config struct {
	Port int
}

func (c *Config) parse() (*Config, error) {
	if c.Port < 3000 || c.Port > 9999 {
		return nil, fmt.Errorf("invalid skyhouse port number '%d'", c.Port)
	}
	return c, nil
}
