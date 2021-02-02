module github.com/ciiiii/sync-image

go 1.14

require (
	docker.io/go-docker v1.0.0
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/caarlos0/env/v6 v6.5.0
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker/internal/testutil v0.0.0-00010101000000-000000000000 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/fatih/structtag v1.2.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/hashicorp/go-getter v1.5.0
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/cobra v1.0.0
	golang.org/x/net v0.0.0-20200421231249-e086a090c8fd // indirect
)

replace github.com/docker/docker/internal/testutil => gotest.tools/v3 v3.0.0
