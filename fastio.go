package fastio

import "os"

const BUF_SIZE = 1 << 20

var buf, obuf [BUF_SIZE]byte
var n, cnt, cnt1 = 0, 0, 0

func GetByte() byte {
	if cnt == n {
		n, _ = os.Stdin.Read(buf[:])
		if n == 0 {
			panic("No more data to read.")
		}
		cnt = 0
	}
	cnt++
	return buf[cnt-1]
}
func GetInt() int {
	var ans = 0
	var c = GetByte()
	for ; c != '-' && (c < '0' || c > '9'); c = GetByte() {
	}
	var flag = false
	if c != '-' {
		ans = int(c - '0')
	} else {
		flag = true
	}
	for c = GetByte(); c >= '0' && c <= '9'; c = GetByte() {
		ans = (ans << 1) + (ans << 3)
		ans += int(c - '0')
	}
	if flag {
		ans = -ans
	}
	return ans
}
func Flush() {
	os.Stdout.Write(obuf[:cnt1])
	cnt1 = 0
}
func PutByte(c byte) {
	if cnt1 == BUF_SIZE {
		Flush()
	}
	obuf[cnt1] = c
	cnt1++
}
func PutInt(x int) {
	if x == 0 {
		PutByte('0')
		return
	}
	var flag = x < 0
	if flag {
		x = -x
	}
	var m = 0
	var tmp [30]byte
	for ; x > 0; x /= 10 {
		tmp[m] = byte(x%10 + '0')
		m++
	}
	if flag {
		PutByte('-')
	}
	for ; m > 0; m-- {
		PutByte(tmp[m-1])
	}
}
