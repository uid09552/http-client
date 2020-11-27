package iehttp

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonParser_Parse(t *testing.T) {
	testJsonParser := jsonParser{}

	testObject := struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	}{"me", 55}

	res, err := testJsonParser.Parse(testObject)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	resString := string(res)
	assert.Equal(t, resString, `{"name":"me","id":55}`)
}

func TestXmlParser_Parse(t *testing.T) {
	testXmlParser := xmlParser{}

	testObject := struct {
		Name    string   `xml:"name"`
		XMLName xml.Name `xml:"test"`
		Id      int      `xml:"id"`
	}{Name: "pln", Id: 55}

	res, err := testXmlParser.Parse(testObject)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	resString := string(res)
	assert.Equal(t, resString, `<test><name>pln</name><id>55</id></test>`)
}
