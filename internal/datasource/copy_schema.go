package datasource

/*
import (
	"fmt"
	"strings"

	"github.com/containers/image/v5/transports"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	imageDescriptionTemplate = fmt.Sprintf(`Container "%s" uses a "transport":"details" format.

Supported transports:
%s
`)
)

func CopyResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"source_image": {
				Description: fmt.Sprintf(imageDescriptionTemplate, strings.Join(transports.ListNames(), ", ")),
				Type:        schema.TypeString,
				Required:    true,
			},
			"destination_image": {
				Description: fmt.Sprintf(imageDescriptionTemplate, strings.Join(transports.ListNames(), ", ")),
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"additional_tags": {
				Type:        schema.TypeList,
				Elem:        schema.TypeString,
				Optional:    true,
				Default:     []string{},
				Description: "additional tags (supports docker-archive)",
			},
			"digest": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Description: `Copy an image (manifest, filesystem layers, signatures) from one location to another.

Uses the system's trust policy to validate images, rejects images not trusted by the policy.

source-image and destination-image are interpreted completely independently; e.g. the destination name does not automatically inherit any parts of the source name.`,
		CreateContext: create,
		ReadContext:   read,
		UpdateContext: update,
		DeleteContext: delete,
		Exists:        exists,
	}
}
*/
