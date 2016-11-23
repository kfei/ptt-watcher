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

func TestContains(t *testing.T) {
	assert := assert.New(t)

	str1 := []string{"slack", "line"}
	assert.True(contains(str1, "slack"))
	assert.True(contains(str1, "line"))
	assert.False(contains(str1, "telegram"))

	str2 := []string{"slack"}
	assert.True(contains(str2, "slack"))
	assert.False(contains(str2, "line"))

	str3 := []string{}
	assert.False(contains(str3, "slack"))
}
