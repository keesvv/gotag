package main

import (
	"fmt"
	"os"

	"github.com/keesvv/gotag/pkg/editor"
	"github.com/keesvv/gotag/pkg/parser"
	"github.com/keesvv/gotag/pkg/tagger"
)

func main() {
	fname := os.Args[1]
	f, err := os.Open(fname)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	tagger := tagger.NewTagger(f)

	if err := tagger.Init(); err != nil {
		panic(err)
	}

	edt := editor.GetPreferred()

	buf, err := edt.GetTempBuffer()
	if err != nil {
		panic(err)
	}

	edt.WriteDefaults(buf, tagger)

	raw, err := edt.Edit(buf)

	if err != nil {
		panic(err)
	}

	p := parser.Parser{}
	contents, err := p.UnmarshalContents(raw)

	if err != nil {
		panic(err)
	}

	if err := tagger.SaveEdits(contents); err != nil {
		panic(err)
	}

	fmt.Println("\033[1m✔ Saved\033[0m")
}
