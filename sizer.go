package stringsizer

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

// StringSizer defines the patterns of encoding and the current state of encoding.
type StringSizer struct {
	Encoding map[string]string `json:"encoding"`
	Current  int               `json:"current"`
	sync.RWMutex
}

// New generates a new map key shrinker. You can optionally pass a previously saved StringSizer that was returned from the `.Save()` function.
func New(savedStringSizer ...string) (ss *StringSizer, err error) {
	ss = new(StringSizer)
	ss.Lock()
	defer ss.Unlock()
	if len(savedStringSizer) > 0 {
		err = json.Unmarshal([]byte(savedStringSizer[0]), &ss)
		if err != nil {
			return
		}
	} else {
		ss.Encoding = make(map[string]string)
		ss.Current = 0
	}
	return
}

// StringSizer will return the a JSON string that can be used to reinitialize the previous state.
func (ss *StringSizer) Save() string {
	ss.RLock()
	defer ss.RUnlock()
	s, err := json.Marshal(ss)
	if err != nil {
		panic(err)
	}
	return string(s)
}

// ShrinkMap takes a map with string keys and converts each string key to the smallest possible string, iterating on the current in the compressor. It returns a new Go map with the transformed string keys.
func (ss *StringSizer) ShrinkMap(m map[string]interface{}) (new map[string]interface{}) {
	ss.Lock()
	defer ss.Unlock()
	new = make(map[string]interface{})
	for key := range m {
		compressedKey := Transform(ss.Current)
		if frossey, ok := ss.Encoding[key]; !ok {
			ss.Encoding[key] = compressedKey
			ss.Current++
		} else {
			compressedKey = frossey
		}
		new[compressedKey] = m[key]
	}
	return
}

// ShrinkMapToString takes a map with string keys and converts each string to the smallest possible string, iterating on the current compressor. It returns a shortened JSON string, so that is, the map `{"value_of_pi":3.141}` will be returned as the string `"a":3.141` (the brackets are removed).
func (ss *StringSizer) ShrinkMapToString(m map[string]interface{}) (shortenedJSON string) {
	mapBytes, err := json.Marshal(ss.ShrinkMap(m))
	if err != nil {
		panic(err)
	}
	shortenedJSON = string(mapBytes[1 : len(mapBytes)-1])
	return
}

// ExpandMap will transform back each string key in a map and return the original map subsituted with its original string keys. Returns an error when a key does not exist.
func (ss *StringSizer) ExpandMap(m map[string]interface{}) (decoded map[string]interface{}, err error) {
	ss.Lock()
	defer ss.Unlock()
	encodingTo := make(map[string]string)
	for k, v := range ss.Encoding {
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

// ExpandMapFromString will return a map from the shortened JSON of a shrinked map (i.e. the output of `ShrinkMapToString`). Returns an error when a key does not exist.
func (ss *StringSizer) ExpandMapFromString(s string) (m map[string]interface{}, err error) {
	encoded := make(map[string]interface{})
	err = json.Unmarshal([]byte("{"+s+"}"), &encoded)
	if err != nil {
		return
	}
	m, err = ss.ExpandMap(encoded)
	return
}

// ShrinkString will take a single string and convert it to the smallest possible based on the current iterator.
func (ss *StringSizer) ShrinkString(original string) (transformed string) {
	ss.Lock()
	defer ss.Unlock()
	transformed = Transform(ss.Current)
	ss.Encoding[original] = transformed
	ss.Current++
	return
}

// ExpandString will take a shrunk string and expand it to the original string. Returns an error when a key does not exist.
func (ss *StringSizer) ExpandString(transformed string) (original string, err error) {
	ss.RLock()
	defer ss.RUnlock()
	for k, v := range ss.Encoding {
		if v == transformed {
			original = k
			return
		}
	}
	err = fmt.Errorf("'%s' not found", transformed)
	return
}
