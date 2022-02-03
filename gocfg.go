package gocfg

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)


type Config map[string]SectionProvider

func NewConfig() Config {
    return Config{}
}

// RawSet sets a value on a key in a section
func (c Config) RawSet(section string, key string, value string) bool {
    sec, ok := c[section]
    if !ok {
        sec = NewMapSectionProvider()
        c[section] = sec
    }
    return sec.RawSet(key, value)
}

// RawGet returns empty string if undefined
func (c Config) RawGet(section string, key string) string {
    sec, ok := c[section]
    if !ok {
        return ""
    }
    return sec.RawGet(key)
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
    return c[section].RawHasKey(key)
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
        var splitted []string
        if scanner.Err() != nil {
            return scanner.Err()
        }
        lineno++
        line := scanner.Text()
        i := 0
        for ; i < len(line); i++ {
            println(line)
            if line[i] == ' ' {
                continue
            }
            if line[i] == ';' || line[i] == '#' {
                goto lineend // skip line
            }
            if line[i] == '[' {
                i++
                j := i
                stack := 0
                for ; j < len(line); j++ {
                    if line[j] == '[' {
                        stack++
                        continue
                    }
                    if line[j] == ']' {
                        if stack > 0 {
                            stack--
                            continue
                        }
                        currentSection = line[i:j]
                        goto lineend
                    }
                }
                return fmt.Errorf("%w near line %d", ErrUnfinishedSection, lineno)
            }
            goto handle_attribution
        }
        goto lineend
        handle_attribution:
        splitted = strings.Split(line, "=")
        if len(splitted) < 2 {
            return fmt.Errorf("%w near line %d", ErrInvalidAttributionSection, lineno)
        }
        c.RawSet(
            currentSection, 
            strings.Trim(splitted[0], " "),
            strings.Trim(strings.Join(splitted[1:], "="), " "),
        )
        lineend:
    }
    return nil
}
