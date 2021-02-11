package servers

import (
	"testing"

	"dam/config"
)

func TestLabel_String(t *testing.T) {
	config.LABEL_REPOS_SEPARATOR = ","

	test := []struct {
		name string
		init string
		want string
	}{
		{
			name: "empty test",
			init: "",
			want: "",
		},
		{
			name: "simple test",
			init: "test,test2,test3",
			want: "test,test2,test3",
		},
		{
			name: "test official repo",
			init: "test,,test3",
			want: "test,test3,",
		},
		{
			name: "test duplicates",
			init: "test,,test3,test3",
			want: "test,test3,",
		},
	}

	for _, tt := range test {
		t.Log("Starting test: " + tt.name)
		s := NewLabel(tt.init)
		if tt.want != s.String() {
			t.Fatalf("String '%v' not equal '%v'", s.String(), tt.want)
		}
	}
}

func TestLabel_AddRepo(t *testing.T) {
	config.LABEL_REPOS_SEPARATOR = ","

	test := []struct {
		name string
		init string
		repo string
		want string
	}{
		{
			name: "add official to empty repos",
			init: "",
			repo: "",
			want: ",",
		},
		{
			name: "simple adding repo test",
			init: "test,test2,test3",
			repo: "test1",
			want: "test,test1,test2,test3",
		},
		{
			name: "add official repo test",
			init: "test,test3",
			repo: "",
			want: "test,test3,",
		},
		{
			name: "repeated add official repo test",
			init: "test,,test3",
			repo: "",
			want: "test,test3,",
		},
		{
			name: "adding duplicates test",
			init: "test,,test3,test3",
			repo: "test3",
			want: "test,test3,",
		},
	}

	for _, tt := range test {
		t.Log("Starting test: " + tt.name)
		s := NewLabel(tt.init)
		t.Logf("Initial struct: %v\n", *s)
		s.AddRepo(tt.repo)
		t.Logf("After adding struct: %v\n", *s)
		if tt.want != s.String() {
			t.Fatalf("Result string '%v' not equal '%v'", s.String(), tt.want)
		}
	}
}

func TestLabel_ValidateRepos(t *testing.T) {
	config.LABEL_REPOS_SEPARATOR = ","

	test := []struct {
		name string
		init string
		error error
		want string
	}{
		{
			name: "simple adding repo test",
			init: "test1",
			error: nil,
			want: "test1",
		},
		//TODO
	}

	for _, tt := range test {
		t.Log("Starting test: " + tt.name)
		s := NewLabel(tt.init)
		t.Logf("Initial struct: %v\n", *s)
		if tt.want != s.String() || tt.error != s.ValidateRepos() {
			t.Errorf("Result string '%v' not equal '%v' ?", s.String(), tt.want)
			t.Fatalf("or error '%v' not equal '%v'", s.ValidateRepos(), tt.error)
		}
	}
}