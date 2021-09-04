package parser

import (
	"bytes"

	"github.com/bogem/id3v2"
	"gopkg.in/yaml.v2"
)

type Parser struct{}

type BufferContents struct {
	Artist      string `yaml:"Artist"`
	AlbumArtist string `yaml:"AlbumArtist"`
	Title       string `yaml:"Title"`
	Album       string `yaml:"Album"`
	Year        string `yaml:"Year"`
}

func (p *Parser) GetBufferContents(tag *id3v2.Tag) *BufferContents {
	return &BufferContents{
		Artist:      tag.Artist(),
		AlbumArtist: tag.GetTextFrame("TPE2").Text,
		Title:       tag.Title(),
		Album:       tag.Album(),
		Year:        tag.Year(),
	}
}

func (p *Parser) MarshalTag(tag *id3v2.Tag) ([]byte, error) {
	c := p.GetBufferContents(tag)

	res, err := yaml.Marshal(c)
	if err != nil {
		return nil, err
	}

	// Strip unnecessary quotes from empty strings
	return bytes.ReplaceAll(res, []byte("\"\""), []byte("")), nil
}

func (p *Parser) UnmarshalContents(raw []byte) (*BufferContents, error) {
	contents := &BufferContents{}
	err := yaml.Unmarshal(raw, contents)

	return contents, err
}
