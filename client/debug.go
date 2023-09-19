//go:build !release
// +build !release

package client

import "path/filepath"

var (
	rootCrtPaths = []string{absPath("assets/client/tls/pgrokroot.crt"), absPath("assets/client/tls/snakeoilca.crt")}
)

func useInsecureSkipVerify() bool {
	return true
}
func absPath(path string) string {
	res, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return res
}
