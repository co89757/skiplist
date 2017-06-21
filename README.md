# Skiplist implmentation in Go 
[![Build Status](https://travis-ci.org/co89757/skiplist.svg?branch=master)](https://travis-ci.org/co89757/skiplist)

## How to install 
```
go get github.com/co89757/skiplist
```
## Example
```go
list := New()
list.Add("a", 1)
v, exists := list.Get("a")
```