package iterativecompressor

import (
	"fmt"
	"testing"
)

func TestCompression(t *testing.T) {
	a := make(map[string]interface{})
	a["zack"] = -42
	a["88:bb:cc"] = "some text"
	a["bb:dd:ee:ff"] = "some other text"
	fmt.Println(a)
}
