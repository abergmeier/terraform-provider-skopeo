package resources

import (
	"context"
	"log"
	"os"
	"regexp"

	"github.com/abergmeier/terraform-provider-skopeo/internal/providerlog"
	"github.com/abergmeier/terraform-provider-skopeo/internal/skopeo"
	skopeoPkg "github.com/abergmeier/terraform-provider-skopeo/pkg/skopeo"
	"github.com/containers/common/pkg/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var ghcr = regexp.MustCompile(`(?::\/\/)?ghcr\.io\/`)

func create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	source := d.Get("source_image").(string)
	destination := d.Get("destination_image").(string)

	reportWriter := providerlog.NewProviderLogWriter(
		log.Default().Writer(),
	)
	defer reportWriter.Close()

	result, err := skopeo.Copy(ctx, source, destination, newCopyOptions(d, reportWriter))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(destination)
	return diag.FromErr(d.Set("docker_digest", result.Digest))
}

func read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	destination := d.Get("destination_image").(string)

	result, err := skopeo.Inspect(ctx, destination, newInspectOptions(d))
	if err != nil {
		return diag.FromErr(err)
	}

	return diag.FromErr(d.Set("docker_digest", result.Digest))
}

func update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if !d.HasChanges("additional_tags", "source_image") {
		return nil
	}

	source := d.Get("source_image").(string)
	destination := d.Get("destination_image").(string)

	reportWriter := providerlog.NewProviderLogWriter(
		log.Default().Writer(),
	)
	defer reportWriter.Close()

	result, err := skopeo.Copy(ctx, source, destination, newCopyOptions(d, reportWriter))
	if err != nil {
		return diag.FromErr(err)
	}
	return diag.FromErr(d.Set("docker_digest", result.Digest))
}

func delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	keep := d.Get("keep_image").(bool)
	if keep {
		return nil
	}

	// We need to delete
	destination := d.Get("destination_image").(string)

	if ghcr.Match([]byte(destination)) {
		return diag.Errorf("GitHub does not support deleting specific container images. Set keep_image to true.")
	}

	return diag.FromErr(skopeoPkg.Delete(ctx, destination, newDeleteOptions(d)))
}

func exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	destination := d.Get("destination_image").(string)

	_, err := skopeo.Inspect(context.TODO(), destination, newInspectOptions(d))
	if err != nil {
		return false, err
	}

	return true, nil
}

func getStringList(d *schema.ResourceData, key string, def []string) []string {
	at := d.Get("additional_tags")
	if at == nil {
		return def
	}
	atl := at.([]interface{})
	additionalTags := make([]string, 0, len(atl))
	for _, t := range atl {
		additionalTags = append(additionalTags, t.(string))
	}
	return additionalTags
}

func newCopyOptions(d *schema.ResourceData, reportWriter *providerlog.ProviderLogWriter) *skopeo.CopyOptions {
	additionalTags := getStringList(d, "additional_tags", nil)

	opts := &skopeo.CopyOptions{
		ReportWriter:   reportWriter,
		SrcImage:       newImageOptions(d),
		DestImage:      newImageDestOptions(d),
		RetryOpts:      newRetyOptions(),
		AdditionalTags: additionalTags,
	}
	return opts
}

func newDeleteOptions(d *schema.ResourceData) *skopeoPkg.DeleteOptions {
	opts := &skopeoPkg.DeleteOptions{
		Image: newImageDestOptions(d).ImageOptions,
	}
	return opts
}

func newGlobalOptions() *skopeoPkg.GlobalOptions {
	opts := &skopeoPkg.GlobalOptions{}
	return opts
}

func newImageDestOptions(d *schema.ResourceData) *skopeoPkg.ImageDestOptions {
	opts := &skopeoPkg.ImageDestOptions{
		ImageOptions: &skopeoPkg.ImageOptions{
			DockerImageOptions: skopeoPkg.DockerImageOptions{
				Global:       newGlobalOptions(),
				Shared:       newSharedImageOptions(),
				AuthFilePath: os.Getenv("REGISTRY_AUTH_FILE"),
			},
		},
	}
	return opts
}

func newImageOptions(d *schema.ResourceData) *skopeoPkg.ImageOptions {
	opts := &skopeoPkg.ImageOptions{
		DockerImageOptions: skopeoPkg.DockerImageOptions{
			Global:       newGlobalOptions(),
			Shared:       newSharedImageOptions(),
			AuthFilePath: os.Getenv("REGISTRY_AUTH_FILE"),
		},
	}
	return opts
}

func newInspectOptions(d *schema.ResourceData) *skopeo.InspectOptions {
	opts := &skopeo.InspectOptions{
		Image: newImageOptions(d),
	}
	return opts
}

func newRetyOptions() *retry.RetryOptions {
	opts := &retry.RetryOptions{}
	return opts
}

func newSharedImageOptions() *skopeoPkg.SharedImageOptions {
	opts := &skopeoPkg.SharedImageOptions{}
	return opts
}
