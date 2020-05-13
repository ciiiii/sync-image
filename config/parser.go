package config

import (
	"reflect"
	"sync"

	"os"

	"github.com/caarlos0/env/v6"
	"github.com/fatih/structtag"
)

type Config struct {
	MirrorRegistry struct {
		Server string `env:"MIRROR_REGISTRY_SERVER" default:"registry.cn-shanghai.aliyuncs.com"`
		Auth   struct {
			Username string `env:"MIRROR_REGISTRY_USERNAME"`
			Password string `env:"MIRROR_REGISTRY_PASSWORD"`
		}
		Docker string `env:"DOCKER" default:"docker_mirror_image"`
		Gcr    string `env:"GCR" default:"gcr_mirror_image"`
		Quay   string `env:"QUAY" default:"quay_mirror_image"`
	}
}

var (
	c    Config
	once sync.Once
)

func init() {
	c.Inject()
}

func Parser() *Config {
	once.Do(func() {
		if err := env.Parse(&c); err != nil {
			panic(err)
		}
	})
	return &c
}

func (c Config) Inject() error {
	for _, name := range []string{"Server", "Docker", "Gcr", "Quay"} {
		field, _ := reflect.TypeOf(c.MirrorRegistry).FieldByName(name)
		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			return err
		}
		envName, err := tags.Get("env")
		if err != nil {
			return err
		}
		envDefault, err := tags.Get("default")
		if err != nil {
			return err
		}
		if os.Getenv(envName.Value()) == "" {
			if err := os.Setenv(envName.Value(), envDefault.Value()); err != nil {
				return err
			}
		}
	}
	return nil
}
