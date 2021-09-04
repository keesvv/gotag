package tagger

import (
	"io"

	"github.com/bogem/id3v2"
	"github.com/keesvv/gotag/pkg/parser"
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

func (t *Tagger) SaveEdits(bc *parser.BufferContents) error {
	// This frame seems to conflict with IDv3
	// so I'm stripping it.
	t.Tag.DeleteFrames("TXXX")

	t.Tag.SetArtist(bc.Artist)
	t.Tag.SetTitle(bc.Title)
	t.Tag.SetAlbum(bc.Album)

	return t.Tag.Save()
}
