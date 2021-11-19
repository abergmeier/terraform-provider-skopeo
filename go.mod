module github.com/abergmeier/terraform-provider-skopeo

go 1.14

require (
	github.com/containers/common v0.46.1-0.20211026130826-7abfd453c86f
	github.com/containers/image/v5 v5.16.2-0.20211021181114-25411654075f
	github.com/containers/ocicrypt v1.1.2
	github.com/containers/storage v1.37.0
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.8.0
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.0.2-0.20210819154149-5ad6f50d6283
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635
)
