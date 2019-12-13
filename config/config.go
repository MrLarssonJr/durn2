// Package config provides a configuration struct and a default configuration.
package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Default is a configuration containing globally useful values.
var Default = func() *Configuration {
	def := New(nil, map[string]string{
		"WEB_PORT": "8080",
		"DEBUG_MODE": "false",
		"DEFAULT_SITE_TITLE": "DUrn",
	})

	def.values["WEB_TEMPLATE_PATH"] = fmt.Sprintf("%s%s%s", "res", string(filepath.Separator), "template")
	def.values["WEB_STYLE_PATH"] = fmt.Sprintf("%s%s%s", "res", string(filepath.Separator), "style")

	/*
	 * Expose OS environment variables,
	 * these will take precedence over those defined in def.
	 */
	envVals := os.Environ()
	envMap := make(map[string]string)
	for _, envVal := range envVals {
		splits := strings.Split(envVal, "=")
		key := splits[0]
		value := splits[1]
		for _, val := range splits[2:] {
			value += "=" + val
		}
		envMap[key] = value
	}

	env := New(def, envMap)

	return env
}()

// Configuration is a key-value mapping from strings to strings. A
// configuration may be queried for configurations values.
// A configuration may also have a reference to a backup configuration
// which is queried if a value does not exist in this config.
type Configuration struct {
	backup *Configuration
	values map[string]string
}

// New creates a new configuration with the values found in the map values.
// If a value does not exist in the map supplied, the backup configuration
// will be queried if supplied.
func New(backup *Configuration, values map[string]string) *Configuration {
	return &Configuration{backup: backup, values: values}
}

// Get queries the configuration for a value. If a value is present, this
// method will return the corresponding string as first return value and
// true as the second return value. If the value isn't present in the config
// or any backup, an empty string will be returned as the first return value
// and false as the second.
func (c *Configuration) Get(key string) (string, bool) {
	if val, exist := c.values[key]; exist {
		return val, true
	}

	if c.backup != nil {
		return c.backup.Get(key)
	}

	return "", false
}

// GetMust is a helper that wraps Get and calls log.Fatal if the returned err
// from Get is non-nil. In that case the program will halt execution and a
// short message will have been logged explaining why. This is intended to aid
// in mandatory variable initialization.
func (c *Configuration) GetMust(key string) string {
	if val, found := c.Get(key); found {
		return val
	} else {
		log.Fatalf("Unable to retrive %s from config in a must context", key)
		return "" // Unreachable, log.Fatalf will exit program
	}
}
