package sync

import (
	"fmt"
	"log"
	"os"

	"github.com/ciiiii/nodocker/core"

	"github.com/ciiiii/sync-image/config"
)

var MirrorRegistryNamespaceMap = map[string]string{
	"registry-1.docker.io": config.Parser().MirrorRegistry.Docker,
	"quay.io": config.Parser().MirrorRegistry.Quay,
	"gcr.io": config.Parser().MirrorRegistry.Gcr,
}

func GetMirrorImageName(i *core.Image) (string, error) {
	namespace, exist := MirrorRegistryNamespaceMap[i.Registry]
	if !exist {
		return "", fmt.Errorf("unknown registry: %s", i.Registry)
	}
	return fmt.Sprintf("%s/%s/%s.%s:%s", config.Parser().MirrorRegistry.Server, namespace, i.Registry, i.Name, i.Tag), nil
}

func Sync(imageStr string) error {
	rawImage, err := core.NewImage(imageStr, false, nil)
	if err != nil {
		return err
	}

	tmpDir := os.TempDir()
	log.Println(tmpDir)
	if err := rawImage.Pull(tmpDir); err != nil {
		return err
	}

	mirrorImageName, err := GetMirrorImageName(rawImage)
	if err != nil {
		return err
	}
	mirrorImage, err := core.NewImage(mirrorImageName, false, config.Parser().MirrorRegistryAccount())
	if err != nil {
		return err
	}
	if err := mirrorImage.Push(rawImage.TargetPath(tmpDir)); err != nil {
		return err
	}
	return nil
}