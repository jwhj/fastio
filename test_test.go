package fastio

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	var f, _ = os.Open("test.in")
	defer f.Close()
	os.Stdin = f
	var n = GetInt()
	for i := 0; i < n; i++ {
		var m = GetInt()
		PutInt(m)
		PutByte('\n')
	}
	Flush()
}
