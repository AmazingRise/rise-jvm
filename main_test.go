package main

import (
	"bytes"
	"fmt"
	core "github.com/AmazingRise/rise-jvm-core"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	arr := []byte{202, 254, 186, 190, 0, 0, 0, 55, 0, 41, 10, 0, 10, 0, 21, 9, 0, 22, 0, 23, 8, 0, 24, 10, 0, 25, 0, 26, 10, 0, 9, 0, 27, 8, 0, 28, 10, 0, 25,
		0, 29, 10, 0, 25, 0, 30, 7, 0, 31, 7, 0, 32, 1, 0, 6, 60, 105, 110, 105, 116, 62, 1, 0, 3, 40, 41, 86, 1, 0, 4, 67, 111, 100, 101, 1, 0, 15, 76, 105,
		110, 101, 78, 117, 109, 98, 101, 114, 84, 97, 98, 108, 101, 1, 0, 4, 109, 97, 105, 110, 1, 0, 22, 40, 91, 76, 106, 97, 118, 97, 47, 108, 97,
		110, 103, 47, 83, 116, 114, 105, 110, 103, 59, 41, 86, 1, 0, 4, 67, 97, 108, 99, 1, 0, 6, 40, 73, 73, 73, 41, 73, 1, 0, 10, 83, 111, 117, 114,
		99, 101, 70, 105, 108, 101, 1, 0, 8, 65, 100, 100, 46, 106, 97, 118, 97, 12, 0, 11, 0, 12, 7, 0, 33, 12, 0, 34, 0, 35, 1, 0, 12, 72, 101, 108, 108,
		111, 32, 119, 111, 114, 108, 100, 33, 7, 0, 36, 12, 0, 37, 0, 38, 12, 0, 17, 0, 18, 1, 0, 14, 40, 49, 48, 48, 43, 50, 48, 48, 41, 42, 51, 48, 48, 61,
		12, 0, 39, 0, 38, 12, 0, 37, 0, 40, 1, 0, 3, 65, 100, 100, 1, 0, 16, 106, 97, 118, 97, 47, 108, 97, 110, 103, 47, 79, 98, 106, 101, 99, 116, 1, 0,
		16, 106, 97, 118, 97, 47, 108, 97, 110, 103, 47, 83, 121, 115, 116, 101, 109, 1, 0, 3, 111, 117, 116, 1, 0, 21, 76, 106, 97, 118, 97, 47, 105,
		111, 47, 80, 114, 105, 110, 116, 83, 116, 114, 101, 97, 109, 59, 1, 0, 19, 106, 97, 118, 97, 47, 105, 111, 47, 80, 114, 105, 110, 116, 83, 116,
		114, 101, 97, 109, 1, 0, 7, 112, 114, 105, 110, 116, 108, 110, 1, 0, 21, 40, 76, 106, 97, 118, 97, 47, 108, 97, 110, 103, 47, 83, 116, 114, 105,
		110, 103, 59, 41, 86, 1, 0, 5, 112, 114, 105, 110, 116, 1, 0, 4, 40, 73, 41, 86, 0, 33, 0, 9, 0, 10, 0, 0, 0, 0, 0, 3, 0, 1, 0, 11, 0, 12, 0, 1, 0, 13, 0,
		0, 0, 29, 0, 1, 0, 1, 0, 0, 0, 5, 42, 183, 0, 1, 177, 0, 0, 0, 1, 0, 14, 0, 0, 0, 6, 0, 1, 0, 0, 0, 1, 0, 9, 0, 15, 0, 16, 0, 1, 0, 13, 0, 0, 0, 96, 0, 3, 0, 5,
		0, 0, 0, 44, 178, 0, 2, 18, 3, 182, 0, 4, 16, 100, 60, 17, 0, 200, 61, 17, 1, 44, 62, 27, 28, 29, 184, 0, 5, 54, 4, 178, 0, 2, 18, 6, 182, 0, 7, 178,
		0, 2, 21, 4, 182, 0, 8, 177, 0, 0, 0, 1, 0, 14, 0, 0, 0, 34, 0, 8, 0, 0, 0, 4, 0, 8, 0, 5, 0, 11, 0, 6, 0, 15, 0, 7, 0, 19, 0, 8, 0, 27, 0, 9, 0, 35, 0, 10, 0,
		43, 0, 11, 0, 9, 0, 17, 0, 18, 0, 1, 0, 13, 0, 0, 0, 30, 0, 2, 0, 3, 0, 0, 0, 6, 26, 27, 96, 28, 104, 172, 0, 0, 0, 1, 0, 14, 0, 0, 0, 6, 0, 1, 0, 0, 0, 17,
		0, 1, 0, 19, 0, 0, 0, 2, 0, 20}
	buf := bytes.NewBufferString("")
	core.Run(bytes.NewReader(arr), buf, os.Stdin, true)
	fmt.Println(buf.String())
}
