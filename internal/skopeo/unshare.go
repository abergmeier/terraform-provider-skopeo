//go:build !linux
// +build !linux

package skopeo

func reexecIfNecessaryForImages(inputImageNames ...string) error {
	return nil
}
