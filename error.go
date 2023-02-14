package maildis

import (
	"fmt"
)

func checkFor(err error, origin string) {
	if err != nil {
		fmt.Printf("Error on %s\n Details: %+v\n", origin, err.Error())
	}
}
