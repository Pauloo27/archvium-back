package utils

import (
	"os"
	"strconv"
)

var envMap = map[string]interface{}{}

func Env(name string) interface{} {
	v, ok := envMap[name]
	if !ok {
		v = os.Getenv("ARCHVIUM_" + name)
		envMap[name] = v
	}
	return v
}

func EnvString(name string) string {
	v, ok := envMap[name]
	if !ok {
		v = os.Getenv("ARCHVIUM_" + name)
		envMap[name] = v
	}
	return v.(string)
}

func EnvBool(name string) bool {
	v, ok := envMap[name]
	if !ok {
		v = os.Getenv("ARCHVIUM_" + name)
		parsedValue, err := strconv.ParseBool(v.(string))
		if err != nil {
			// TODO: what?
		}
		envMap[name] = parsedValue
		v = parsedValue
	}
	return v.(bool)
}

func EnvInt64(name string) int64 {
	v, ok := envMap[name]
	if !ok {
		v = os.Getenv("ARCHVIUM_" + name)
		parsedValue, err := strconv.Atoi(v.(string))
		if err != nil {
			// TODO: what?
		}
		envMap[name] = int64(parsedValue)
		v = int64(parsedValue)
	}
	return v.(int64)
}
