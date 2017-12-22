package iterativecompressor

import (
	"encoding/json"
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

func (ic *IterativeCompressor) Dumps(m map[string]interface{}) (new string) {
	newMap := ic.Encode(m)
	mapBytes, err := json.Marshal(newMap)
	if err != nil {
		panic(err)
	}
	return string(mapBytes[1 : len(mapBytes)-1])
}
