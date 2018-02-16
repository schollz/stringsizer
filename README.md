# stringsizer

[![travis](https://travis-ci.org/schollz/stringsizer.svg?branch=master)](https://travis-ci.org/schollz/stringsizer) 
[![go report card](https://goreportcard.com/badge/github.com/schollz/stringsizer)](https://goreportcard.com/report/github.com/schollz/stringsizer) 
[![coverage](https://img.shields.io/badge/coverage-95%25-brightgreen.svg)](https://gocover.io/github.com/schollz/stringsizer)
[![godocs](https://godoc.org/github.com/schollz/stringsizer?status.svg)](https://godoc.org/github.com/schollz/stringsizer) 

A very simple way to exchange keys in a map for a shorter version of the key (1-2 chars). Basically it converts `{"some-long-key":"data"}` into `"a":"data"`. It keeps track of how the map keys are converted so they can be converted back to the original. _Note_: The resulting data is *not* JSON (its missing `{}`) which makes it a little smaller, and also forces you to transform back to use it.

## Why?

I plan on encoding the same set of 10-100 MAC addresses in JSON payloads that are each inserted to a row of a database. This way will be a fast and efficient way to store encode the JSON names across every row so that it reduces the size of the keys (from 17 bytes to 1).

## Usage

```golang
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
```

# License 

MIT
