package main

import (
	"os"
	"runtime/pprof"
	"testing"
)

func TestFunc(t *testing.T) {
	f, _ := os.OpenFile("mem.profile", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()

	for i := 0; i <= 1000; i++ {
		main()
	}
	pprof.Lookup("heap").WriteTo(f, 0)
}
