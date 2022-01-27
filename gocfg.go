package gocfg

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Section map[string]string

type Config map[string]Section

func NewConfig() Config {
    return Config{}
}

// RawSet sets a value on a key in a section
func (c Config) RawSet(section string, key string, value string) {
    sec, ok := c[section]
    if !ok {
        sec = Section{}
        c[section] = sec
    }
    sec[key] = value
}

// RawGet returns empty string if undefined
func (c Config) RawGet(section string, key string) string {
    sec, ok := c[section]
    if !ok {
        return ""
    }
    val, ok := sec[key]
    if !ok {
        return ""
    }
    return val
}

func (c Config) RawHasSection(section string) bool {
    _, ok := c[section]
    return ok
}

func (c Config) RawHasKey(section string, key string) bool {
    _, ok := c[section]
    if !ok {
        return false
    }
    _, ok = c[section][key]
    return ok
}

var (
    ErrUnfinishedSection = errors.New("Unfinished section expression")
    ErrInvalidAttributionSection = errors.New("Invalid attribution section")
)

func (c Config) InjestReader(r io.Reader) error {
    scanner := bufio.NewScanner(r)
    currentSection := ""
    lineno := 0
    for scanner.Scan() {
        if scanner.Err() != nil {
            return scanner.Err()
        }
        lineno++
        line := scanner.Text()
        i := 0
        for ; i < len(line); i++ {
            if line[i] == ' ' {
                continue
            }
            if line[i] == ';' || line[i] == '#' {
                break // skip line
            }
            if line[i] == '[' {
                i++
                j := i
                for ; j < len(line); j++ {
                    if line[j] == ']' {
                        currentSection = line[i:j]
                        goto sectionend
                    }
                }
                return fmt.Errorf("%w near line %d", ErrUnfinishedSection, lineno)
                sectionend:
            }
            splitted := strings.Split(line, "=")
            if len(splitted) < 2 {
                return fmt.Errorf("%w near line %d", ErrInvalidAttributionSection, lineno)
            }
            c.RawSet(currentSection, splitted[0], strings.Join(splitted[1:], "="))
        }
    }
    return nil
}