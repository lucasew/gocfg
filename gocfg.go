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
	ErrUnfinishedSection         = errors.New("Unfinished section expression")
	ErrInvalidAttributionSection = errors.New("Invalid attribution section")
)

func parseSection(line string, i int, lineno int) (string, error) {
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
			return line[i:j], nil
		}
	}
	return "", fmt.Errorf("%w near line %d", ErrUnfinishedSection, lineno)
}

func parseAttribution(line string, currentSection string, lineno int, c Config) error {
	splitted := strings.Split(line, "=")
	if len(splitted) < 2 {
		return fmt.Errorf("%w near line %d", ErrInvalidAttributionSection, lineno)
	}
	c.RawSet(
		currentSection,
		strings.Trim(splitted[0], " "),
		strings.Trim(strings.Join(splitted[1:], "="), " "),
	)
	return nil
}

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

		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || trimmedLine[0] == ';' || trimmedLine[0] == '#' {
			continue
		}

		if trimmedLine[0] == '[' {
			// Find the start of the section name
			startIndex := strings.Index(line, "[") + 1
			sectionName, err := parseSection(line, startIndex, lineno)
			if err != nil {
				return err
			}
			currentSection = sectionName
			continue
		}

		err := parseAttribution(line, currentSection, lineno, c)
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		reportError(err, map[string]interface{}{"method": "InjestReader"})
		return err
	}
	return nil
}
