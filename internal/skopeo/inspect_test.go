package skopeo

import (
	"context"
	"testing"
)

func TestInspect(t *testing.T) {
	t.Parallel()

	out, err := Inspect(context.TODO(), "docker://alpine:latest", &InspectOptions{
		Image: &ImageOptions{
			DockerImageOptions: DockerImageOptions{
				Global: &GlobalOptions{
					debug: true,
				},
				Shared: &SharedImageOptions{},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if out.Digest == "" {
		t.Fatal("Digest not expected")
	}
}
