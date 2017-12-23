package mapkeys

import (
	"encoding/json"
	"errors"
)

type MapKeys struct {
	From    map[string]string
	To      map[string]string
	Current int
}

// Init generates a new map key compressor
func Init(compressorArg ...string) (mk *MapKeys, err error) {
	mk = new(MapKeys)
	if len(compressorArg) > 0 {
		err = json.Unmarshal([]byte(compressorArg[0]), &mk)
		if err != nil {
			return
		}
	} else {
		mk.From = make(map[string]string)
		mk.To = make(map[string]string)
		mk.Current = 0
	}
	return
}

// Save will generate a JSON of the current map keys structure to be
// initialized for the next time you use it.
func (mk *MapKeys) Save() string {
	s, err := json.Marshal(mk)
	if err != nil {
		panic(err)
	}
	return string(s)
}

// Shrink will convert each key to the smallest possible string, iterating
// on the current in the compressor.
func (mk *MapKeys) Shrink(m map[string]interface{}) (new map[string]interface{}) {
	new = make(map[string]interface{})
	for key := range m {
		compressedKey := Transform(mk.Current)
		if fromKey, ok := mk.From[key]; !ok {
			mk.From[key] = compressedKey
			mk.To[compressedKey] = key
			mk.Current++
		} else {
			compressedKey = fromKey
		}
		new[compressedKey] = m[key]
	}
	return
}

// Expand will convert each key to the original name.
func (mk *MapKeys) Expand(m map[string]interface{}) (decoded map[string]interface{}, err error) {
	decoded = make(map[string]interface{})
	for compressedKey := range m {
		if key, ok := mk.To[compressedKey]; ok {
			decoded[key] = m[compressedKey]
		} else {
			err = errors.New("could not find key '" + compressedKey + "' during decoding")
		}
	}
	return
}

// Dumps will return a string of the JSON encoded shrunk map key structure.
func (mk *MapKeys) Dumps(m map[string]interface{}) (new string) {
	newMap := mk.Shrink(m)
	mapBytes, err := json.Marshal(newMap)
	if err != nil {
		panic(err)
	}
	return string(mapBytes[1 : len(mapBytes)-1])
}

// Loads will return a map from the dumped string.
func (mk *MapKeys) Loads(s string) (m map[string]interface{}, err error) {
	encoded := make(map[string]interface{})
	err = json.Unmarshal([]byte("{"+s+"}"), &encoded)
	if err != nil {
		return
	}
	m, err = mk.Expand(encoded)
	return
}
