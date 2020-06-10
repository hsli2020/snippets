package config

import "sync"

// Config is a custom type for holding and changing configuration data.
type Config struct {
	sync.RWMutex
	item string
}

// New will create a new instance of configuration.
// Each instance is unique and can hold different information.
func New() *Config {
	// setup the config object
	return &Config{}
}

// Item is used to access Config.item
func (c *Config) Item() string {
	c.RLock()
	defer c.RUnlock()
	return c.item
}

// SetItem is used to change Config.item
func (c *Config) SetItem(s string) {
	c.Lock()
	defer c.Unlock()
	c.item = s
}