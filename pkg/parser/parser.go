package parser

import (
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

func (p *Parser) MarshalTag(tag *id3v2.Tag) ([]byte, error) {
	c := &BufferContents{
		Artist:      tag.Artist(),
		AlbumArtist: tag.GetTextFrame("TPE2").Text,
		Title:       tag.Title(),
		Album:       tag.Album(),
		Year:        tag.Year(),
	}

	return yaml.Marshal(c)
}

func (p *Parser) UnmarshalContents(raw []byte) (*BufferContents, error) {
	contents := &BufferContents{}
	err := yaml.Unmarshal(raw, contents)

	return contents, err
}
