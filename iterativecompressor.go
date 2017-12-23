package iterativecompressor

import (
	"encoding/json"
	"errors"
)

type IterativeCompressor struct {
	From    map[string]string
	To      map[string]string
	Current int
}

// New generates a new compressor
func New(compressorArg ...string) (ic *IterativeCompressor, err error) {
	ic = new(IterativeCompressor)
	if len(compressorArg) > 0 {
		err = json.Unmarshal([]byte(compressorArg[0]), &ic)
		if err != nil {
			return
		}
	} else {
		ic.From = make(map[string]string)
		ic.To = make(map[string]string)
		ic.Current = 0
	}
	return
}

func (ic *IterativeCompressor) Save() string {
	s, err := json.Marshal(ic)
	if err != nil {
		panic(err)
	}
	return string(s)
}

func (ic *IterativeCompressor) Encode(m map[string]interface{}) (new map[string]interface{}) {
	new = make(map[string]interface{})
	for key := range m {
		compressedKey := Transform(ic.Current)
		if fromKey, ok := ic.From[key]; !ok {
			ic.From[key] = compressedKey
			ic.To[compressedKey] = key
			ic.Current++
		} else {
			compressedKey = fromKey
		}
		new[compressedKey] = m[key]
	}
	return
}

func (ic *IterativeCompressor) Decode(m map[string]interface{}) (decoded map[string]interface{}, err error) {
	decoded = make(map[string]interface{})
	for compressedKey := range m {
		if key, ok := ic.To[compressedKey]; ok {
			decoded[key] = m[compressedKey]
		} else {
			err = errors.New("could not find key '" + compressedKey + "' during decoding")
		}
	}
	return
}

func (ic *IterativeCompressor) Dumps(m map[string]interface{}) (new string) {
	newMap := ic.Encode(m)
	mapBytes, err := json.Marshal(newMap)
	if err != nil {
		panic(err)
	}
	return string(mapBytes[1 : len(mapBytes)-1])
}

func (ic *IterativeCompressor) Loads(s string) (m map[string]interface{}, err error) {
	encoded := make(map[string]interface{})
	err = json.Unmarshal([]byte("{"+s+"}"), &encoded)
	if err != nil {
		return
	}
	m, err = ic.Decode(encoded)
	return
}
