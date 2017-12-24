package mapslimmer

import (
	"encoding/json"
	"fmt"
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
		mk, _ := Init()
		mk.Dumps(a)
	}
}

func ExampleDumps() {
	// make a map
	a := make(map[string]interface{})
	a["some-long-key"] = "data"

	// Slim the map to string
	ms, _ := Init()
	aString := ms.Dumps(a)
	fmt.Println(aString)

	// Store the slimmer to expand it later
	slimmer := ms.Slimmer()
	fmt.Println(slimmer)

	// Slim another map loading the previous slimmer
	a2 := make(map[string]interface{})
	a2["some-long-key"] = "data2"
	ms2, _ := Init(slimmer)
	a2String := ms2.Dumps(a2)
	fmt.Println(a2String)

	// Output: "a":"data"
	// {"From":{"some-long-key":"a"},"To":{"a":"some-long-key"},"Current":1}
	// "a":"data2"

}

func TestCompression(t *testing.T) {
	a := make(map[string]interface{})
	a["zack"] = -42.0
	a["88:bb:cc"] = "some text"
	a["bb:dd:ee:ff"] = "some other text"
	mk, err := Init()
	assert.Nil(t, err)
	assert.Equal(t, 45, len(mk.Dumps(a)))
	ac := mk.Dumps(a)

	bytesA, _ := json.Marshal(a)
	fmt.Println(string(bytesA))
	fmt.Println(ac)
	fmt.Println(mk.Slimmer())

	b := make(map[string]interface{})
	b["zack"] = -32.0
	b["88:bb:cc"] = "!text"
	b["bb:dd:ee:ff"] = "hi again"
	b["bb:dd:ee:fg"] = "hi again"
	assert.True(t, len(mk.Dumps(b)) <= 50)
	bc := mk.Dumps(b)
	bc = mk.Dumps(b)

	bcd, err := mk.Loads(bc)
	assert.Nil(t, err)
	assert.Equal(t, b, bcd)
	acd, err := mk.Loads(ac)
	assert.Nil(t, err)
	assert.Equal(t, a, acd)

	mkSave := mk.Slimmer()
	mkLoad, err := Init(mkSave)
	assert.Nil(t, err)
	assert.Equal(t, mk.Current, mkLoad.Current)
	bcd2, err := mkLoad.Loads(bc)
	assert.Nil(t, err)
	assert.Equal(t, bcd, bcd2)
}
