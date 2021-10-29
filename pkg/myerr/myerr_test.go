package myerr

import (
	"fmt"
	"testing"
)

func TestMyerr(t *testing.T) {
	err := Newf("error occur")
	fmt.Printf("%s", err)
	err = Newf("error occur: %s", "error")
	fmt.Printf("%s", err)
}
