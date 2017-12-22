package iterativecompressor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompression(t *testing.T) {
	a := make(map[string]interface{})
	a["zack"] = -42
	a["88:bb:cc"] = "some text"
	a["bb:dd:ee:ff"] = "some other text"
	ic, err := New()
	assert.Nil(t, err)
	assert.Equal(t, 45, len(ic.Dumps(a)))

	a = make(map[string]interface{})
	a["zack"] = -32
	a["88:bb:cc"] = "!text"
	a["bb:dd:ee:ff"] = "hi again"
	a["bb:dd:ee:fg"] = "hi again"
	assert.Equal(t, 50, len(ic.Dumps(a)))
}
