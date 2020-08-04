package gw2api

import (
	"strconv"
	"strings"
)

func concatStrings(args ...string) string {
	var str strings.Builder
	for _, x := range args {
		str.WriteString(x)
	}
	return str.String()
}

func itoaSlice(sl []int) (strs []string) {
	for _, x := range sl {
		strs = append(strs, strconv.Itoa(x))
	}
	return
}
