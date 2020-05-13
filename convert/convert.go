package convert

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"fmt"

	"github.com/ciiiii/sync-image/sync"
	"github.com/hashicorp/go-getter"
)

type Converter struct {
	Source      string
	Destination string
	RegistryMap *sync.RegistryMap
	Target      string
}

func (c *Converter) Parse() error {
	if err := c.Get(); err != nil {
		return err
	}
	content, err := ioutil.ReadFile(c.Destination)
	if err != nil {
		return err
	}
	var registryMap sync.RegistryMap
	if err := json.Unmarshal(content, &registryMap); err != nil {
		return err
	}
	c.RegistryMap = &registryMap
	return nil
}

func (c *Converter) StringMapper() []string {
	var s []string
	for _, i := range c.RegistryMap.Iter() {
		s = append(s, i.Full, i.Rename())
		if i.Source == "docker.io" {
			if i.Registry == "library" {
				s = append(s, i.Image, i.Rename())
			} else {
				s = append(s, fmt.Sprintf("%s/%s", i.Registry, i.Image), i.Rename())
			}
		}
	}
	return s
}

// Deprecated
func (c *Converter) Mapper() map[string]string {
	m := make(map[string]string, c.RegistryMap.Len())
	for _, i := range c.RegistryMap.Iter() {
		m[i.Full] = i.Rename()
	}
	return m
}

func (c *Converter) Replace(s []string) error {
	replacer := ReplacerGenerator(c.Target, s)
	if err := filepath.Walk(c.Target, replacer); err != nil {
		return err
	}
	return nil
}

func (c *Converter) Get() error {
	ctx, _ := context.WithCancel(context.Background())
	if dstFile, err := NewTempFile(); err != nil {
		return err
	} else {
		c.Destination = dstFile
	}
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	getterClient := getter.Client{
		Ctx:  ctx,
		Src:  c.Source,
		Dst:  c.Destination,
		Pwd:  pwd,
		Mode: getter.ClientModeFile,
	}
	return getterClient.Get()
}
