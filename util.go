package gw2api

import (
	"strconv"
	"strings"
)

// concatStrings concates multiple strings using byte buffers for higher performance
func concatStrings(args ...string) string {
	switch len(args) {
	case 0:
		return ""
	case 1:
		return args[0]
	}

	n := 0
	for _, s := range args {
		n += len(s)
	}

	var b strings.Builder
	b.Grow(n)
	for _, s := range args {
		b.WriteString(s)
	}
	return b.String()
}

// iotaSlice converts a slice of ints into a slice of strings
func itoaSlice(sl []int) []string {
	var strs []string
	for _, x := range sl {
		strs = append(strs, strconv.Itoa(x))
	}
	return strs
}
