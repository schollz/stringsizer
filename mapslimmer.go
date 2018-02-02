package mapslimmer

import (
	"encoding/json"
	"errors"
)

// MapSlimmer defines the patterns of encoding and the current state of encoding.
type MapSlimmer struct {
	Encoding map[string]string `json:"encoding"`
	Current  int               `json:"current"`
}

// Init generates a new map key shrinker
func Init(slimmerArg ...string) (mk *MapSlimmer, err error) {
	mk = new(MapSlimmer)
	if len(slimmerArg) > 0 {
		err = json.Unmarshal([]byte(slimmerArg[0]), &mk)
		if err != nil {
			return
		}
	} else {
		mk.Encoding = make(map[string]string)
		mk.Current = 0
	}
	return
}

// Slimmer will return the MapSlimmer JSON that can be used to
// reinitialize the previous state.
func (mk *MapSlimmer) JSON() string {
	s, err := json.Marshal(mk)
	if err != nil {
		panic(err)
	}
	return string(s)
}

// Slim will convert each key to the smallest possible string, iterating
// on the current in the compressor.
func (mk *MapSlimmer) Slim(m map[string]interface{}) (new map[string]interface{}) {
	new = make(map[string]interface{})
	for key := range m {
		compressedKey := transform(mk.Current)
		if fromKey, ok := mk.Encoding[key]; !ok {
			mk.Encoding[key] = compressedKey
			mk.Current++
		} else {
			compressedKey = fromKey
		}
		new[compressedKey] = m[key]
	}
	return
}

// Expand will convert each key to the original name.
func (mk *MapSlimmer) Expand(m map[string]interface{}) (decoded map[string]interface{}, err error) {
	encodingTo := make(map[string]string)
	for k, v := range mk.Encoding {
		encodingTo[v] = k
	}
	decoded = make(map[string]interface{})
	for compressedKey := range m {
		if key, ok := encodingTo[compressedKey]; ok {
			decoded[key] = m[compressedKey]
		} else {
			err = errors.New("could not find key '" + compressedKey + "' during decoding")
		}
	}
	return
}

// Dumps will return a string of the JSON encoded shrunk map key structure.
func (mk *MapSlimmer) Dumps(m map[string]interface{}) (new string) {
	newMap := mk.Slim(m)
	mapBytes, err := json.Marshal(newMap)
	if err != nil {
		panic(err)
	}
	return string(mapBytes[1 : len(mapBytes)-1])
}

// Loads will return a map from the dumped string.
func (mk *MapSlimmer) Loads(s string) (m map[string]interface{}, err error) {
	encoded := make(map[string]interface{})
	err = json.Unmarshal([]byte("{"+s+"}"), &encoded)
	if err != nil {
		return
	}
	m, err = mk.Expand(encoded)
	return
}
