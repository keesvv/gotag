package main

import (
	"fmt"
	"os"

	id3 "github.com/bogem/id3v2"
	"github.com/keesvv/gotag/pkg/editor"
)

func main() {
	fname := os.Args[1]
	f, err := os.Open(fname)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	tag, err := id3.ParseReader(f, id3.Options{Parse: true})

	if err != nil {
		panic(err)
	}

	defer tag.Close()

	edt := editor.GetPreferred()

	buf, err := edt.GetTempBuffer()
	if err != nil {
		panic(err)
	}

	buf.Write([]byte("Test\n"))

	contents, err := edt.Edit(buf)

	if err != nil {
		panic(err)
	}

	fmt.Printf("File contents: %v\n", contents)
	fmt.Println("\033[1m✔ Saved\033[0m")
}
