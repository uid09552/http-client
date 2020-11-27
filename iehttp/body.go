package iehttp

import (
	"encoding/json"
	"encoding/xml"
	"strings"
)

const (
	TypeJson = "application/json"
	TypeXml  = "application/xml"
)

type Parser interface {
	Parse(body interface{}) ([]byte, error)
}

func NewParser(contentType string) Parser {
	switch strings.ToLower(contentType) {
	case TypeJson:
		return &jsonParser{}
	case TypeXml:
		return &xmlParser{}
	default:
		return &jsonParser{}
	}
}

type jsonParser struct{}

func (j *jsonParser) Parse(body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	return json.Marshal(body)
}

type xmlParser struct{}

func (j *xmlParser) Parse(body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	return xml.Marshal(body)
}
