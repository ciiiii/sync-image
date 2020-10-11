package config

import "github.com/ciiiii/nodocker/core"

type Config struct {
    MirrorRegistry struct {
        Server string `env:"MIRROR_REGISTRY_SERVER" envDefault:"registry.cn-shanghai.aliyuncs.com"`
        Auth   struct {
            Username string `env:"MIRROR_REGISTRY_USERNAME"`
            Password string `env:"MIRROR_REGISTRY_PASSWORD"`
        }
        Docker string `env:"MIRROR_DOCKER_NAMESPACE" envDefault:"docker_mirror_image"`
        Gcr    string `env:"MIRROR_GCR_NAMESPACE" envDefault:"gcr_mirror_image"`
        Quay   string `env:"MIRROR_QUAY_NAMESPACE" envDefault:"quay_mirror_image"`
    }
    Server struct{
        Port int `env:"PORT"`
        Mode string `env:"MODE" envDefault:"release"`
    }
}

func (c Config) MirrorRegistryAccount() *core.RegistryAccount {
    return &core.RegistryAccount{
        Username: c.MirrorRegistry.Auth.Username,
        Password: c.MirrorRegistry.Auth.Password,
    }
}