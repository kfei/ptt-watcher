package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	assert := assert.New(t)

	str := "This is a topic string"

	filters0 := []Filter{}
	assert.True(filteredAny(str, filters0))

	filters1 := []Filter{""}
	assert.True(filteredAny(str, filters1))

	filters2 := []Filter{"This"}
	assert.True(filteredAny(str, filters2))

	filters3 := []Filter{"this is"}
	assert.True(filteredAny(str, filters3))

	filters4 := []Filter{"this is", "topic"}
	assert.True(filteredAny(str, filters4))

	filters5 := []Filter{"this is", "other string"}
	assert.True(filteredAny(str, filters5))

	filters6 := []Filter{"not"}
	assert.False(filteredAny(str, filters6))

	filters7 := []Filter{"not", "appear"}
	assert.False(filteredAny(str, filters7))
}
