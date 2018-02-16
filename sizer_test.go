package stringsizer

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
		ss, _ := New()
		ss.ShrinkMapToString(a)
	}
}

func ExampleDumps() {
	// make a map
	someMap := make(map[string]interface{})
	someMap["some-long-key"] = "data"

	// Create a new string sizer
	ss, _ := New()

	// create a shortened string JSON version of the map
	shortedJSONOfSomeMap := ss.ShrinkMapToString(someMap)
	fmt.Println(shortedJSONOfSomeMap)
	// "a":"data"

	// Store the string sizer to expand it later
	saved := ss.Save()
	fmt.Println(saved)
	// {"encoding":{"some-long-key":"a"},"current":1}

	// reload the string sizer
	ss2, _ := New(saved)
	// reload the original map from shortened string version
	originalMap, _ := ss2.ExpandMapFromString(shortedJSONOfSomeMap)
	fmt.Println(originalMap)
	// map[some-long-key:data]

	// Output: "a":"data"
	// {"encoding":{"some-long-key":"a"},"current":1}
	// map[some-long-key:data]
}

func TestCompression(t *testing.T) {
	// make a new string sizer
	ss, err := New()
	assert.Nil(t, err)

	// check that its empty
	saved := ss.Save()
	assert.Equal(t, "{\"encoding\":{},\"current\":0}", saved)

	// check that map shrinks correctly
	a := make(map[string]interface{})
	a["zack"] = -42.1
	ss.ShrinkMap(a) // I'm consecutively adding keys here to preserve ordering
	a["88:bb:cc:gg"] = 1.2
	ss.ShrinkMap(a)
	a["bb:dd:ee:ff"] = "test"
	shrinkedA := ss.ShrinkMap(a)
	assert.Equal(t, map[string]interface{}{"a": -42.1, "b": 1.2, "c": "test"}, shrinkedA)

	// test that it outputs strings correctly
	shrinkedAString := ss.ShrinkMapToString(a)
	assert.Equal(t, "\"a\":-42.1,\"b\":1.2,\"c\":\"test\"", shrinkedAString)

	// check that it expands map correctly
	expandedShrinkedA, err := ss.ExpandMap(shrinkedA)
	assert.Nil(t, err)
	assert.Equal(t, a, expandedShrinkedA)

	// check that it throws errors on non-shrinked map
	_, err = ss.ExpandMap(a)
	assert.NotNil(t, err)

	// check that it expands map from string
	expandedShrinkedA, err = ss.ExpandMapFromString(shrinkedAString)
	assert.Nil(t, err)
	for key := range a {
		assert.Equal(t, a[key], expandedShrinkedA[key])
	}

	// check that the single keys work too
	originalString := "some long string"
	newString := ss.ShrinkString(originalString)
	assert.Equal(t, "ba", newString)
	expandedString, err := ss.ExpandString(newString)
	assert.Nil(t, err)
	assert.Equal(t, originalString, expandedString)
	_, err = ss.ExpandString("this string doesn't exist in map")
	assert.NotNil(t, err)
}
