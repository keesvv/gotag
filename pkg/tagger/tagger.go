package tagger

import (
	"io"

	"github.com/bogem/id3v2"
	"github.com/keesvv/gotag/pkg/parser"
)

type Tagger struct {
	reader io.Reader
	parser *parser.Parser
	Tag    *id3v2.Tag
}

func NewTagger(reader io.Reader) *Tagger {
	return &Tagger{
		reader: reader,
		parser: &parser.Parser{},
	}
}

func (t *Tagger) Init() error {
	tag, err := id3v2.ParseReader(t.reader, id3v2.Options{Parse: true})
	t.Tag = tag
	return err
}

func (t *Tagger) isChanged(newBc *parser.BufferContents) bool {
	originalBc := t.parser.GetBufferContents(t.Tag)

	return *originalBc != *newBc
}

func (t *Tagger) SaveEdits(bc *parser.BufferContents) (bool, error) {
	if !t.isChanged(bc) {
		return false, nil
	}

	// This frame seems to conflict when saving
	// frames so I'm stripping it.
	t.Tag.DeleteFrames("TXXX")

	t.Tag.SetArtist(bc.Artist)
	t.Tag.SetTitle(bc.Title)
	t.Tag.SetAlbum(bc.Album)
	t.Tag.SetYear(bc.Year)
	t.Tag.AddTextFrame("TPE2", t.Tag.DefaultEncoding(), bc.AlbumArtist)

	err := t.Tag.Save()
	if err != nil {
		return false, err
	}

	return true, nil
}
