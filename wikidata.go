package gowiki

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	apiURL = "http://www.wikidata.org/w/api.php?format=json&"
	client = &http.Client{}
)

type WikiResult struct {
	Entities map[string]interface{} `json:"entities"`
}

func (wr WikiResult) Get(key string) (e *Entity) {
	e = &Entity{id: key, data: wr.Entities[key].(map[string]interface{})}

	return
}

func NewWikiResult(data []byte) (wr *WikiResult, err error) {
	// dec := json.NewDecoder(bytes.NewReader(data))

	// if err := dec.Decode(&WikiResult); err != nil {
	// 	return fmt.Errorf("parse wikidata results: %v", err)
	// }
	var ent map[string]interface{}

	err = json.Unmarshal(data, &ent)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	wr = &WikiResult{Entities: ent["entities"].(map[string]interface{})}

	return
}

func FetchBatch(ids []string) (wr *WikiResult, err error) {
	// query := map[string]interface{}{
	// 	"ids": ids
	// }

	resp, err := http.Get(apiURL + "action=wbgetentities&ids=" + strings.Join(ids, ",") + "&props=claims")

	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	wr, err = NewWikiResult(body)

	return
}

func Fetch(id string) (e *Entity, err error) {
	wr, err := FetchBatch([]string{id})

	if err != nil {
		return
	}

	e = wr.Get(id)
	return
}

type Entity struct {
	id   string
	data map[string]interface{}
}

func NewEntity(data map[string]interface{}) (e *Entity) {
	e = &Entity{data: data}

	return
}

func (e Entity) GetLocation() (location map[string]interface{}) {
	locValue := e.data["claims"].(map[string]interface{})["P625"].([]interface{})[0].(map[string]interface{})["mainsnak"].(map[string]interface{})["datavalue"].(map[string]interface{})["value"].(map[string]interface{})
	location = map[string]interface{}{"latitude": locValue["latitude"], "longitude": locValue["longitude"]}
	return
}
