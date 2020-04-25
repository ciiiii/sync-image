package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/ciiiii/sync-image/config"
)

const errImageNotExist = "Error response from daemon: manifest unknown: manifest unknown"

type Client struct {
	*docker.Client
	C       context.Context
	AuthStr string
}

func NewClient(ctx context.Context) *Client {
	var err error
	dockerClient, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	return &Client{
		Client:  dockerClient,
		C:       ctx,
		AuthStr: auth(config.Parser().MirrorRegistry.Server, config.Parser().MirrorRegistry.Auth.Username, config.Parser().MirrorRegistry.Auth.Password),
	}
}

func auth(server, username, password string) string {
	authConfig := types.AuthConfig{
		ServerAddress: server,
		Username:      username,
		Password:      password,
	}

	encodeJson, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(encodeJson)
}

func (c *Client) Login() error {
	s, err := c.RegistryLogin(c.C, types.AuthConfig{
		ServerAddress: config.Parser().MirrorRegistry.Server,
		Username:      config.Parser().MirrorRegistry.Auth.Username,
		Password:      config.Parser().MirrorRegistry.Auth.Password,
	})
	fmt.Println(s.Status)
	return err
}

func (c *Client) Pull(image string) error {
	o, err := c.ImagePull(c.C, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer o.Close()
	io.Copy(os.Stdout, o)
	return nil
}

func (c *Client) Push(image string) error {
	o, err := c.ImagePush(c.C, image, types.ImagePushOptions{RegistryAuth: c.AuthStr})
	if err != nil {
		return err
	}
	defer o.Close()
	io.Copy(os.Stdout, o)
	return nil
}

func (c *Client) Exist(image string) (bool, error) {
	_, err := c.DistributionInspect(c.C, image, c.AuthStr)
	if err != nil {
		if err.Error() == errImageNotExist {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (c *Client) Tag(src, tgt string) error {
	return c.ImageTag(c.C, src, tgt)
}
