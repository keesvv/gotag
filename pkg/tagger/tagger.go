package tagger

import (
	"io"
	"io/ioutil"

	"github.com/bogem/id3v2"
	"github.com/gabriel-vasile/mimetype"
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

func (t *Tagger) AddFrontCover(fname string) error {
	pic, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	t.Tag.AddAttachedPicture(id3v2.PictureFrame{
		Encoding:    t.Tag.DefaultEncoding(),
		MimeType:    mimetype.Detect(pic).String(),
		PictureType: id3v2.PTFrontCover,
		Description: "Front cover",
		Picture:     pic,
	})

	return nil
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

	if bc.Front != "" {
		err := t.AddFrontCover(bc.Front)

		if err != nil {
			return false, err
		}
	}

	err := t.Tag.Save()
	if err != nil {
		return false, err
	}

	return true, nil
}
