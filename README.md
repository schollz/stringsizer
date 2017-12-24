# mapslimmer

[![travis](https://travis-ci.org/schollz/mapslimmer.svg?branch=master)](https://travis-ci.org/schollz/mapslimmer) 
[![go report card](https://goreportcard.com/badge/github.com/schollz/mapslimmer)](https://goreportcard.com/report/github.com/schollz/mapslimmer) 
[![coverage](https://img.shields.io/badge/coverage-92%25-brightgreen.svg)](https://gocover.io/github.com/schollz/mapslimmer)
[![godocs](https://godoc.org/github.com/schollz/mapslimmer?status.svg)](https://godoc.org/github.com/schollz/mapslimmer) 

A very simple way to exchange keys in a map for a shorter version of the key (1-2 chars). Basically it converts `{"some-long-key":"data"}` into `"a":"data"`. It keeps track of how the map keys are converted so they can be converted back to the original. _Note_: The resulting data is *not* JSON (its missing `{}`) which makes it a little smaller, and also forces you to transform back to use it.



## Why?

I plan on encoding the same set of 10-100 MAC addresses in JSON payloads that are each inserted to a row of a database. This way will be a fast and efficient way to store encode the JSON names across every row so that it reduces the size of the keys (from 17 bytes to 1).

## Usage

```golang
// make a map
a := make(map[string]interface{})
a["some-long-key"] = "data"

// Slim the map to string
ms, _ := Init()
aString := ms.Dumps(a)
// "a":"data"

// Store the slimmer to expand it later
slimmer := ms.Slimmer()
// {"From":{"some-long-key":"a"},"To":{"a":"some-long-key"},"Current":1}

// Slim another map loading the previous slimmer
a2 := make(map[string]interface{})
a2["some-long-key"] = "data2"
ms2, _ := Init(slimmer)
a2String := ms2.Dumps(a2)
// "a":"data2"
```

# License 

MIT
