package internal

import "testing"

func TestSplitTag(t *testing.T) {
	var testTags = []string{
		"server/name:version",
		"docker.io/library/ubuntu:14.04",
		"ubuntu:16.04",
		"localhost:5000/ubuntu:18.04",
	}

	for _, tag := range testTags {
		s, n, v := SplitTag(tag)
		if s != "" {
			s = s+"/"
		}
		if tag != s+n+":"+v {
			t.Fatalf("Not equal split tag '%v': '%v', '%v', '%v'",tag, s, n, v)
		}
	}
}
