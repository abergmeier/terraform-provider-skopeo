package resources

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"testing"

	"github.com/bsquare-corp/terraform-provider-skopeo/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	testAccProviders = map[string]*schema.Provider{
		"skopeo": provider.Provider(),
	}
)

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = "ghcr.io"
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	serveReverseProxy(req.RequestURI, res, req)
}

func TestAccCopy(t *testing.T) {
	t.Parallel()

	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCopyResource(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(fmt.Sprintf("skopeo_copy.alpine_%s", rName), "docker_digest"),
				),
			},
			{
				Config: testAccCopyResource_addTag(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(fmt.Sprintf("skopeo_copy.alpine_%s", rName), "docker_digest"),
				),
			},
		},
	})
}

func testAccCopyResource(name string) string {
	return fmt.Sprintf(`resource "skopeo_copy" "alpine_%s" {
	source_image      = "docker://alpine"
	destination_image = "docker://ghcr.io/bsquare-corp/alpine"
}`, name)
}

func testAccCopyResource_addTag(name string) string {
	return fmt.Sprintf(`resource "skopeo_copy" "alpine_%s" {
	source_image      = "docker://alpine"
	destination_image = "docker://ghcr.io/bsquare-corp/alpine"
	additional_tags   = ["alpine:fine"]
	keep_image        = true
}`, name)
}
