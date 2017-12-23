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
	a["zack"] = -42.0
	a["88:bb:cc"] = "some text"
	a["bb:dd:ee:ff"] = "some other text"
	ic, err := New()
	assert.Nil(t, err)
	assert.Equal(t, 45, len(ic.Dumps(a)))
	ac := ic.Dumps(a)

	b := make(map[string]interface{})
	b["zack"] = -32.0
	b["88:bb:cc"] = "!text"
	b["bb:dd:ee:ff"] = "hi again"
	b["bb:dd:ee:fg"] = "hi again"
	assert.True(t, len(ic.Dumps(b)) <= 50)
	bc := ic.Dumps(b)
	bc = ic.Dumps(b)

	bcd, err := ic.Loads(bc)
	assert.Nil(t, err)
	assert.Equal(t, b, bcd)
	acd, err := ic.Loads(ac)
	assert.Nil(t, err)
	assert.Equal(t, a, acd)

	icSave := ic.Save()
	icLoad, err := New(icSave)
	assert.Nil(t, err)
	assert.Equal(t, ic.Current, icLoad.Current)
	bcd2, err := icLoad.Loads(bc)
	assert.Nil(t, err)
	assert.Equal(t, bcd, bcd2)
}
