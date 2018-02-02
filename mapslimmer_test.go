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
	someMap := make(map[string]interface{})
	someMap["some-long-key"] = "data"

	// Slim the map to string
	ms, _ := Init()
	aString := ms.Dumps(someMap)
	fmt.Println(aString)
	// "a":"data"

	// Store the slimmer to expand it later
	slimmerJSON := ms.JSON()
	fmt.Println(slimmerJSON)
	// {"encoding":{"some-long-key":"a"},"current":1}

	// Expand another map loading the previous slimmer
	someOtherMap := make(map[string]interface{})
	someOtherMap["a"] = "data2"
	ms, _ = Init(slimmerJSON)
	someOtherMapDecoded, _ := ms.Expand(someOtherMap)
	fmt.Println(someOtherMapDecoded)
	// map[some-long-key:data2]

	// Output: "a":"data"
	// {"encoding":{"some-long-key":"a"},"current":1}
	// map[some-long-key:data2]

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
	fmt.Println(mk.JSON())

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

	mkSave := mk.JSON()
	mkLoad, err := Init(mkSave)
	assert.Nil(t, err)
	assert.Equal(t, mk.Current, mkLoad.Current)
	bcd2, err := mkLoad.Loads(bc)
	assert.Nil(t, err)
	assert.Equal(t, bcd, bcd2)
}
