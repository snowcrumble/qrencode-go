package qrencode

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestTerminalOutput(t *testing.T) {
	//gen random string
	var b = make([]byte, 100)
	_, err := rand.Read(b)
	if err != nil {
		t.Fatalf("rand.Read err: %v", err)
	}
	s := fmt.Sprintf("%x", b)

	grid, err := Encode(s, ECLevelL)
	if err != nil {
		t.Fatalf("Encode err: %v", err)
	}

	//compare two qrcode size in terminal
	grid.TerminalOutput(os.Stdout)
	grid.TerminalOutputCompress(os.Stdout)
}
