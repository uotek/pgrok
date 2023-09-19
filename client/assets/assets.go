package assets

import (
	"os"
)

func Asset(arg string) ([]byte, error) {
	return os.ReadFile(arg)
}
