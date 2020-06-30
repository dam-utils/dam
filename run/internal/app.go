package internal

import "strings"

func SplitTag(tag string) (string, string, string) {
	n := strings.Split(tag, "/")
	nameWithVersion := n[len(n)-1]
	server := strings.Join(n[:len(n)-1], "/")

	v := strings.Split(nameWithVersion, ":")
	version := v[len(v)-1]
	name := v[0]

	return server, name, version
}
