package skopeo

import (
	"context"
	"fmt"

	"github.com/containers/common/pkg/retry"
	"github.com/containers/image/v5/transports/alltransports"
)

type DeleteOptions struct {
	Image     *ImageOptions
	retryOpts *retry.RetryOptions
}

func Delete(ctx context.Context, imageName string, opts *DeleteOptions) error {
	if err := reexecIfNecessaryForImages(imageName); err != nil {
		return err
	}

	ref, err := alltransports.ParseImageName(imageName)
	if err != nil {
		return fmt.Errorf("Invalid source name %s: %v", imageName, err)
	}

	sys, err := opts.Image.newSystemContext()
	if err != nil {
		return err
	}

	return retry.RetryIfNecessary(ctx, func() error {
		return ref.DeleteImage(ctx, sys)
	}, opts.retryOpts)
}
