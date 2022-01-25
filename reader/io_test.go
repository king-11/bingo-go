package reader

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGetScanner(t *testing.T) {
	f, err := os.CreateTemp("", "bingo_io_test_*")
	if err != nil {
		t.Errorf("os.CreateTemp() failed: %v", err)
		return
	}
	defer f.Close()
	val := []string{"1", "2", "3", "4", "5"}
	f.WriteString(strings.Join(val, " "))

	scanner := GetScanner(f)
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

func generateRandomNumbers(n int) []int {
	rand.Seed(time.Now().Unix())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(4242)
	}

	return nums
}

func intSliceToString(nums []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(nums), " ", delim, -1), "[]")
}

func TestReadOrder(t *testing.T) {
	f, err := os.CreateTemp("", "bingo_io_order_test_*")
	if err != nil {
		t.Errorf("os.CreateTemp() failed: %v", err)
		return
	}
	defer f.Close()
	val := generateRandomNumbers(6)
	f.WriteString(intSliceToString(val, ","))

	scanner := GetScanner(f)
	scanner.Split(bufio.ScanLines)

	order, err := ReadOrder(scanner)
	if err != nil {
		t.Errorf("ReadOrder() failed: %v", err)
		return
	}
	for i, v := range order {
		if v != val[i] {
			t.Errorf("ReadOrder() failed, expected:%d found:%d", val[i], v)
			return
		}
	}
}

func TestReadBoard(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Errorf("os.CreateTemp() failed: %v", err)
		return
	}
	defer f.Close()
	tests := [][5][5]int{
		{
			{34, 90, 18, 33, 83},
			{27, 7, 25, 61, 15},
			{43, 5, 51, 32, 45},
			{24, 17, 72, 31, 22},
			{77, 46, 78, 16, 9},
		},
		{
			{72, 95, 37, 52, 68},
			{80,  1, 73, 96, 63},
			{16, 49,  9, 42, 97},
			{25, 81, 20, 11, 46},
			{31, 24,  2, 34, 18},
		},
		{
			{88, 29, 95, 98, 57},
			{49, 36,  6, 23, 83},
			{18,  5, 45, 40, 44},
			{62, 81, 74, 99, 87},
			{46, 56, 35, 21, 52},
		},
	}
	scanner := GetScanner(f)
	scanner.Split(bufio.ScanLines)
	for _, test := range tests {
		board, err := ReadBoard(scanner)
		if err != nil {
			t.Errorf("ReadBoard() failed: %v", err)
			return
		}
		for i, v := range board {
			for j, val := range v {
				if test[i][j] != val.Val {
					t.Errorf("ReadBoard() failed at index:(%d,%d) expected:%d found:%d", i, j, test[i][j], val.Val)
					return
				}
			}
		}
		scanner.Scan()
	}
}
