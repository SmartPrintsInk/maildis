package maildis

import (
	"os"
)

var Host string

func setup() {
	Host = os.Getenv("HostIP")
}
