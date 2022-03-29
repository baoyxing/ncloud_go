package core

import "time"

type Config struct {
	Scheme   string
	Endpoint string
	Timeout  time.Duration
}

func NewConfig() *Config {
	return &Config{SchemeHttps, "www.ncloud-api.com", 10 * time.Second}
}

func (c *Config) SetScheme(scheme string) {
	c.Scheme = scheme
}

func (c *Config) SetEndpoint(endpoint string) {
	c.Endpoint = endpoint
}

func (c *Config) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}
