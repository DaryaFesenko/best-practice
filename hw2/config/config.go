package config

type Config struct {
	path   *string
	remove *bool
	debug  *bool
}

func New(path *string, remove, debug *bool) Config {
	// Линтер gocritic - создание объекта
	c := Config{}

	c.debug = debug
	c.path = path
	c.remove = remove

	return c
}

func (c *Config) GetPath() string {
	return *c.path
}

func (c *Config) GetDebug() bool {
	return *c.debug
}

func (c *Config) GetRemove() bool {
	return *c.remove
}
