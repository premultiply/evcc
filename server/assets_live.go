// +build !release

package server

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("--- LIVE ---")
	Assets = os.DirFS("dist")
}
