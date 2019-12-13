package cmd

import (
	"fmt"
	"gih/application/cmd"
	"testing"
)

func TestParseRepositoryUrl(t *testing.T) {
	h, r, org, err := cmd.ParseRepositoryUrl("https://github.com/rennnosuke/state")
	fmt.Printf("%v,%v,%v,%v\n", h, r, org, err)
}
