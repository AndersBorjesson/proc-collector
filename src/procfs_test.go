package main

import (
	"fmt"
	"testing"

	"github.com/prometheus/procfs"
)

func TestUptime(t *testing.T) {
	FS, err := procfs.NewDefaultFS()
	if err != nil {
		t.Errorf("Error %v", err)
	}
	a, err := FS.AllProcs()
	if err != nil {
		t.Errorf("Error %v", err)
	}
	for _, v := range a {
		D, err := v.Interrupts()
		if err != nil {
			t.Errorf("Error %v", err)
		} else {

			fmt.Println(" S", D, "SS")
		}
	}
}
