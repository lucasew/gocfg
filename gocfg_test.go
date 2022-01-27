package gocfg

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

const demoCFG = `
# some comment
; Another comment
 ; An indented comment

a = 2
b = 3

[section]
other_a = 2
other_b = 9

[https://google.com]
title = Google
jsonValue = {"os": "Linux"}

[[helloworld]]
# should be the [helloworld] section
isthisright = true
`

func TestParseLine(t *testing.T) {
    c := NewConfig()
    r := bytes.NewBufferString(demoCFG)
    err := c.InjestReader(r)
    if err != nil {
        t.Error(err)
    }
    json.NewEncoder(os.Stdout).Encode(c)
}
