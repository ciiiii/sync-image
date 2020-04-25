package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ciiiii/sync-image/config"
	"github.com/ciiiii/sync-image/docker"
)

type Image map[string][]string

type Registry map[string]Image

func (r Registry) Iter(t string) []ImageInfo {
	var namespace, source string
	switch t {
	case "docker":
		namespace = config.Parser().MirrorRegistry.Docker
		source = "docker.io"
		break
	case "gcr":
		namespace = config.Parser().MirrorRegistry.Gcr
		source = "gcr.io"
		break
	case "quay":
		namespace = config.Parser().MirrorRegistry.Quay
		source = "quay.io"
		break
	}

	var l []ImageInfo
	for registry, imageList := range r {
		for image, tagList := range imageList {
			imageName := strings.Join([]string{source, registry, image}, "/")
			for _, tag := range tagList {
				full := imageName + ":" + tag
				l = append(l, ImageInfo{
					Full:      full,
					Source:    source,
					Namespace: namespace,
					Registry:  registry,
					Image:     image,
					Tag:       tag,
				})
			}
		}
	}
	return l
}

type RegistryMap struct {
	Gcr    Registry `json:"gcr.io"`
	Quay   Registry `json:"quay.io"`
	Docker Registry `json:"docker.io"`
}

func (r RegistryMap) Iter() []ImageInfo {
	var imageList []ImageInfo
	for _, image := range r.Docker.Iter("docker") {
		imageList = append(imageList, image)
	}
	for _, image := range r.Gcr.Iter("gcr") {
		imageList = append(imageList, image)
	}
	for _, image := range r.Quay.Iter("quay") {
		imageList = append(imageList, image)
	}
	return imageList
}

type ImageInfo struct {
	Full      string
	Source    string
	Namespace string
	Registry  string
	Image     string
	Tag       string
}

func (i ImageInfo) Rename() string {
	if i.Registry == "google-containers" {
		return fmt.Sprintf("%s/%s/%s:%s", config.Parser().MirrorRegistry.Server, i.Namespace, i.Image, i.Tag)
	}
	return fmt.Sprintf("%s/%s/%s.%s:%s", config.Parser().MirrorRegistry.Server, i.Namespace, i.Registry, i.Image, i.Tag)
}

func main() {
	ctx := context.Background()
	dockerClient := docker.NewClient(ctx)

	if err := dockerClient.Login(); err != nil {
		panic(err)
	}

	f, _ := ioutil.ReadFile("images.json")
	var registryMap RegistryMap
	if err := json.Unmarshal(f, &registryMap); err != nil {
		panic(err)
	}
	for _, image := range registryMap.Iter() {
		fmt.Printf("[image]%s\n", image.Full)
		mirrorImage := image.Rename()
		fmt.Println(":check")
		exist, err := dockerClient.Exist(mirrorImage)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if exist {
			fmt.Println(":exist")
			continue
		}

		fmt.Println(":pull")
		if err := dockerClient.Pull(image.Full, false); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(":tag")
		if err := dockerClient.Tag(image.Full, mirrorImage); err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(":push")
		if err := dockerClient.Push(mirrorImage, false); err != nil {
			fmt.Println(err)
		}
	}
}
