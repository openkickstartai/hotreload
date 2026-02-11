package main

import ("os";"gopkg.in/yaml.v3")

type Config struct {
	Watch []string `yaml:"watch"`
	Ignore []string `yaml:"ignore"`
	Command string `yaml:"command"`
	Restart bool `yaml:"restart"`
	Debounce int `yaml:"debounce"`
}

func LoadConfig(p string) (Config, error) {
	var c Config
	d, err := os.ReadFile(p); if err != nil { return c, err }
	err = yaml.Unmarshal(d, &c)
	if c.Debounce == 0 { c.Debounce = 200 }
	return c, err
}
