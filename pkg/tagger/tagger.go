package tagger

import (
	"io"

	"github.com/bogem/id3v2"
)

type Tagger struct {
	reader io.Reader
	Tag    *id3v2.Tag
}

func NewTagger(reader io.Reader) *Tagger {
	return &Tagger{
		reader: reader,
	}
}

func (t *Tagger) Init() error {
	tag, err := id3v2.ParseReader(t.reader, id3v2.Options{Parse: true})
	t.Tag = tag
	return err
}
