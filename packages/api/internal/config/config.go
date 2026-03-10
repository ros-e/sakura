/* I was gonna write a db but i felt like it would be necessary for this; so we're doing the kuro approach */

package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Profile  *Profile  `yaml:"profile"`
	Settings *Settings `yaml:"settings"`
}

type Profile struct {
	Username  string    `yaml:"username"`
	Password  string    `yaml:"passwordHash"`
	CreatedAt time.Time `yaml:"createdAt"`
}

type Settings struct {
	Wallpaper string `yaml:"wallpaper"`
}

func Default() *Config {
	return &Config{
		Settings: &Settings{
			Wallpaper: "dark_mountains.jpg",
		},
	}
}

func Init(path string) error {
	cfg := &Config{
		Settings: &Settings{
			Wallpaper: "dark_mountains.jpg",
		},
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func Save(cfg *Config, path string) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
