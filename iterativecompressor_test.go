package iterativecompressor

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkNormalMarshal(b *testing.B) {
	a := make(map[string]interface{})
	a["zack"] = -42
	a["88:bb:cc"] = "some text"
	a["bb:dd:ee:ff"] = "some other text"
	for n := 0; n < b.N; n++ {
		json.Marshal(a)
	}
}

func BenchmarkCompressing(b *testing.B) {
	a := make(map[string]interface{})
	a["zack"] = -42
	a["88:bb:cc"] = "some text"
	a["bb:dd:ee:ff"] = "some other text"
	for n := 0; n < b.N; n++ {
		ic, _ := New()
		ic.Dumps(a)
	}
}

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
