package skopeo

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/abergmeier/terraform-provider-skopeo/internal/providerlog"
)

func TestCopy(t *testing.T) {

	t.Parallel()

	reportWriter := providerlog.NewProviderLogWriter(
		log.Default().Writer(),
	)
	defer reportWriter.Close()

	writeDir := t.TempDir()
	result, err := Copy(context.TODO(), "docker://alpine:latest", fmt.Sprintf("dir:%s", writeDir), &CopyOptions{
		ReportWriter: reportWriter,
		SrcImage: &ImageOptions{
			DockerImageOptions: DockerImageOptions{
				Global: &GlobalOptions{
					debug: true,
				},
				Shared: &SharedImageOptions{},
			},
		},
		DestImage: &ImageDestOptions{
			ImageOptions: &ImageOptions{
				DockerImageOptions: DockerImageOptions{
					Global: &GlobalOptions{
						debug: true,
					},
					Shared: &SharedImageOptions{},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = os.Stat(fmt.Sprintf("%s/manifest.json", writeDir))
	if err != nil {
		files := readDir(t, writeDir)
		t.Fatalf("Expected manifest.json. Found %s. Error: %s", files, err)
	}
	if result.Digest == "" {
		t.Fatal("Digest should be empty")
	}
}

func readDir(t *testing.T, dir string) (entries []string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		t.Error(err)
	}

	for _, f := range files {
		entries = append(entries, f.Name())
	}

	return
}
