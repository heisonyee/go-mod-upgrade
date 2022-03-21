package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestGoModUpgrade_1(t *testing.T) {
	args := []string{
		"list",
		"-u",
		"-mod=mod",
		"-f",
		"'{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}'",
		"-m",
		"all",
	}
	cmdstr := "go " + strings.Join(args, " ")

	fmt.Printf("%s\n", cmdstr)
}
