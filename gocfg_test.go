package gocfg

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestParseLine(t *testing.T) {
    c := NewConfig()
    r := bytes.NewBufferString("; hello world\na=2")
    err := c.InjestReader(r)
    if err != nil {
        t.Error(err)
    }
    json.NewEncoder(os.Stdout).Encode(c)
}
