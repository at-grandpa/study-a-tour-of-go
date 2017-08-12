package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	// bのサイズだけ、bに読み込む
	// 読み込んだ時のerrもそのままerrに格納
	size, err := rot13.r.Read(b)

	// strings.Readerによってbにstringが読み込まれているので、
	// それをforで回し、各byte対してrot13を適用する
	// updatしたいのはbなので、bに格納する
	for i := range b {
		b[i] = rot13byte(b[i])
	}
	return size, err
}

// rot13:
// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
// NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm

func rot13byte(b byte) byte {
	if 'A' <= b && b <= 'M' || 'a' <= b && b <= 'm' {
		return b + 13
	}
	if 'N' <= b && b <= 'Z' || 'n' <= b && b <= 'z' {
		return b - 13
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
