package exitem

import (
	"testing"
)

func TestRoot_ReadConfig(t *testing.T) {
	root := ReadConfig("config.yaml")

	if root != nil {
		t.Log(root)
	} else {
		t.Error("")
	}
}
