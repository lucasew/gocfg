package gocfg

import "os"

type SectionProvider interface {
    // RawGet gets a string value from the section provider, "" if not defined
    RawGet(key string) (value string)
    // RawSet sets a string value on a section provider, false if not writeable
    RawSet(key string, value string) (success bool)
    // RawHasKey checks if the key exists in the section
    RawHasKey(key string) (hasKey bool)
}

type MapSectionProvider map[string]string

func NewMapSectionProvider() SectionProvider {
    return MapSectionProvider(map[string]string{})
}

func (sp MapSectionProvider) RawGet(key string) string {
    ret, ok := sp[key]
    if !ok {
        return ""
    }
    return ret
}

func (sp MapSectionProvider) RawSet(key string, value string) bool {
    sp[key] = value
    return true
}

func (sp MapSectionProvider) RawHasKey(key string) bool {
    _, ok := sp[key]
    return ok
}

type envSectionProvider int

var EnvSectionProvider SectionProvider = envSectionProvider(0)

func (envSectionProvider) RawGet(key string) string {
    return os.Getenv(key)
}

func (envSectionProvider) RawSet(key string, value string) bool {
    return os.Setenv(key, value) != nil
}

func (envSectionProvider) RawHasKey(key string) bool {
    _, ok := os.LookupEnv(key)
    return ok
}
