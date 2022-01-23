package reader

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestGetScanner(t *testing.T) {
	f, err := os.CreateTemp("", "bingo_io_test_*")
	if err != nil {
		t.Errorf("os.CreateTemp() failed: %v", err)
		return
	}
	defer f.Close()
	val := []string{"1", "2", "3", "4" ,"5"}
	f.WriteString(strings.Join(val, " "))

	scanner, err := GetScanner(f)
	if err != nil {
		t.Errorf("GetScanner() failed: %v", err)
		return
	}
	scanner.Split(bufio.ScanWords)

	i := 0
	for scanner.Scan() {
		if scanner.Text() != val[i] {
			t.Errorf("GetScanner() failed: %v", err)
			return
		}
		i++
	}
}
