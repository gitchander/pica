package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// cmpf it is simple analog of cmp (command in Linux/UNIX)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("arguments must have 2 files")
	}

	var (
		name1 = os.Args[1]
		name2 = os.Args[2]
	)

	if name1 == name2 { // files have equal names
		return
	}

	f1, err := os.Open(name1)
	checkError(err)
	defer f1.Close()

	f2, err := os.Open(name2)
	checkError(err)
	defer f2.Close()

	const bufSize = 4096
	var (
		buf1 = make([]byte, bufSize)
		buf2 = make([]byte, bufSize)
	)

	var pos int64
	var eof1, eof2 bool

	for {
		n1, err := io.ReadFull(f1, buf1)
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				eof1 = true
			} else {
				//log.Fatal(err)
				log.Fatalf("%s readFull error: %s", name1, err)
			}
		}

		n2, err := io.ReadFull(f2, buf2)
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				eof2 = true
			} else {
				//log.Fatal(err)
				log.Fatalf("%s readFull error: %s", name2, err)
			}
		}

		n := minInt(n1, n2)
		for i := 0; i < n; i++ {
			var (
				b1 = buf1[i]
				b2 = buf2[i]
			)

			if b1 != b2 {
				shift := pos + int64(i) + 1

				fmt.Println(shift, byteHex(b1), byteHex(b2))
				fmt.Printf("%s %s differ: byte %d\n", name1, name2, shift)

				return
			}
		}

		pos += int64(n)

		if (eof1 || eof2) && (n1 == n2) { // files equal!
			return
		}

		if eof1 {
			printFileEOF(name1, pos)
			return
		}

		if eof2 {
			printFileEOF(name2, pos)
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printFileEOF(name string, pos int64) {
	fmt.Printf("cmpf: EOF in %s after byte %d\n", name, pos)
}

func byteToNibbles(b byte) (lo, hi byte) {
	lo = b & 0xF
	hi = b >> 4
	return
}

var hexUpper = []byte("0123456789ABCDEF")

func byteHex(b byte) string {
	lo, hi := byteToNibbles(b)
	data := []byte{'0', 'x', hexUpper[hi], hexUpper[lo]}
	return string(data)
}
