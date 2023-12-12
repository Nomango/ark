package ark_test

import (
	"fmt"
	"testing"

	"github.com/Nomango/ark"
)

func TestRandom(t *testing.T) {
	s := ark.NewRandomString(32)
	fmt.Println(s)
}
