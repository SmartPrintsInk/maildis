package maildis

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	emails := ListFor("Sales")
	fmt.Printf("%+v\n", emails)
}
