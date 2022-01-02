package resources

import (
	"context"
	"testing"
)

func TestReadMissing(t *testing.T) {
	d := CopyResource().TestResourceData()
	err := d.Set("source_image", "docker://ghcr.io/abergmeier/nonexistingsource")
	if err != nil {
		t.Fatal(err)
	}
	err = d.Set("destination_image", "docker://ghcr.io/abergmeier/nonexistingdestination")
	if err != nil {
		t.Fatal(err)
	}
	d.SetId("dummyid")

	diags := read(context.Background(), d, nil)
	if diags != nil {
		t.Fatalf("Error %#v\n", diags[0])
	}

	id := d.Id()
	if id != "" {
		t.Fatalf("Id not reset: %v", d)
	}

	err = d.Set("destination_image", "docker://ghcr.io/abergmeier/terraform-provider-skopeo/alpine:nonexistant")
	if err != nil {
		t.Fatal(err)
	}

	d.SetId("dummyid")

	diags = read(context.Background(), d, nil)
	if diags != nil {
		t.Fatalf("Error %#v\n", diags[0])
	}

	id = d.Id()
	if id != "" {
		t.Fatalf("Id not reset: %v", d)
	}
}
