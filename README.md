# iterativecompressor

[![travis](https://travis-ci.org/schollz/iterativecompressor.svg?branch=master)](https://travis-ci.org/schollz/iterativecompressor) 
[![go report card](https://goreportcard.com/badge/github.com/schollz/iterativecompressor)](https://goreportcard.com/report/github.com/schollz/iterativecompressor) 
[![coverage](https://img.shields.io/badge/coverage-92%25-brightgreen.svg)](https://gocover.io/github.com/schollz/iterativecompressor)
[![godocs](https://godoc.org/github.com/schollz/iterativecompressor?status.svg)](https://godoc.org/github.com/schollz/iterativecompressor) 

A very simple way to encode short strings. Basically the converted keeps a running tally and converts any string to a base representation of the next number in the tally.

Basically it converts something like:

```json
{"ab:cd:ef:gh:ij":-42,"zack":"something"}
```

into

```
"a":-42,"b":"something"
```

while generating a compressor that includes the values that were exchanged:

```json
{
  "To":{"a":"ab:cd:ef:gh:ij","b":"zack"}
}
```
