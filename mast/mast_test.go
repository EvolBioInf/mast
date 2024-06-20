package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestMas(t *testing.T) {
	tests := make([]*exec.Cmd, 0)
	test := exec.Command("./mast", "q1.fasta", "r1.fasta")
	tests = append(tests, test)
	test = exec.Command("./mast", "q2.fasta", "r2.fasta")
	tests = append(tests, test)
	for i, test := range tests {
		get, err := test.Output()
		if err != nil {
			t.Error(err)
		}
		f := "r" + strconv.Itoa(i+1) + ".txt"
		want, err := os.ReadFile(f)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(get, want) {
			t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
		}
	}
}
