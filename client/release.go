//go:build release
// +build release

package client

var (
	rootCrtPaths = []string{absPath("assets/client/tls/pgrokroot.crt")}
)

func useInsecureSkipVerify() bool {
	return false
}
