package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func main() {
	// from stdin
	/*
		r := os.Stdin
		w := os.Stdout
	*/
	// from buffer 任意の文字列をinputとできる
	var r, w bytes.Buffer
	fmt.Fprintf(&r, "main ")
	fmt.Fprintf(&r, "run")

	print(&r, &w)
}

// readerが読み取って、writerへ書き出す
func print(r io.Reader, w io.Writer) {
	var s string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s = scanner.Text()
		fmt.Fprintf(w, "%s\n", s)
	}
	fmt.Print(w)
}
