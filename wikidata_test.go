package gowiki

import (
	"testing"
)

func TestWikiFetch(t *testing.T) {
	_, err := FetchBatch([]string{"Q963138"})

	if err != nil {
		t.Error("WikiFetch error", err)
	}

	// n := len(resp["response"][0])

	// s, ok := resp["response"].(string)

	// if !ok {
	// 	t.Error("Couldn't convert to string")
	// }
}

func TestWikiFetchEntity(t *testing.T) {
	_, err := Fetch("Q963138")

	if err != nil {
		t.Error("WikiFetchEntity error", err)
	}
}

func TestWikiFetchLocation(t *testing.T) {
	entity, err := Fetch("Q963138")

	if err != nil {
		t.Error("WikiFetchEntity error", err)
	}

	loc := entity.GetLocation()

	t.Log(loc)
}
