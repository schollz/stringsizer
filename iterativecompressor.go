package iterativecompressor

import (
	"encoding/json"
	"fmt"
)

type IterativeCompressor struct {
	From    map[string]string
	To      map[string]string
	Current int
}

// New generates a new compressor
func New(compressorArg ...string) (ic *IterativeCompressor, err error) {
	ic = new(IterativeCompressor)
	if len(compressorArg) == 0 {
		err = json.Unmarshal([]byte(compressorArg[0]), &ic)
		if err != nil {
			return
		}
	} else {
		ic.From = make(map[string]string)
		ic.To = make(map[string]string)
		ic.Current = 1
	}
	return
}

func (ic *IterativeCompressor) Dumps() string {
	s, err := json.Marshal(ic)
	if err != nil {
		panic(err)
	}
	return string(s)
}

func transform(a map[string]interface{}) (new map[string]interface{}) {
	new = make(map[string]interface{})
	fmt.Println(a)
	for key := range a {
		new[key] = a[key]
	}
	return
}
