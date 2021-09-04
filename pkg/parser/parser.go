package parser

import (
	"github.com/bogem/id3v2"
	"gopkg.in/yaml.v2"
)

type Parser struct{}

type BufferContents struct {
	Artist string `yaml:"Artist"`
	Title  string `yaml:"Title"`
	Album  string `yaml:"Album"`
}

func (p *Parser) MarshalTag(tag *id3v2.Tag) ([]byte, error) {
	c := &BufferContents{
		Artist: tag.Artist(),
		Title:  tag.Title(),
		Album:  tag.Album(),
	}

	return yaml.Marshal(c)
}

func (p *Parser) UnmarshalContents(raw []byte) (*BufferContents, error) {
	contents := &BufferContents{}
	err := yaml.Unmarshal(raw, contents)

	return contents, err
}
