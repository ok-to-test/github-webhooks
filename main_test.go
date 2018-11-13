package main


import (
	"fmt"
	"os"
	"testing"
)

func TestRequireEnv(t *testing.T) {
	env := os.Getenv("SOME_ENV_VAR")
	if env == "" {
		fmt.Errorf("Cannot access env variables")
		os.Exit(1)
	}
}
