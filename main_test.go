package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChopperName(t *testing.T) {
	chopper := Chopper{Person{name: "Chris"}}
	assert.Equal(t, "Chris", chopper.name)
}

func TestRiceCookerName(t *testing.T) {
	t.Skip()

	riceCooker := RiceCooker{Person{name: "Rick"}}
	assert.Equal(t, "Rick", riceCooker.name)
}

func TestWrapperName(t *testing.T) {
	t.Skip()

	wrapper := Wrapper{Person{name: "Will"}}
	assert.Equal(t, Ingredient{name: "Will"}, wrapper.name)
}

func TestChopperWorks(t *testing.T) {
	t.Skip()

	chopper := Chopper{Person{name: "Chris"}}
	assert.Equal(t, Ingredient{name: "🥒 cucumber"}, chopper.work())
}

func TestRiceCookerWorks(t *testing.T) {
	t.Skip()

	riceCooker := RiceCooker{Person{name: "Rick"}}
	assert.Equal(t, Ingredient{name: "🍚 rice"}, riceCooker.work())
}

func TestWrapperWorks(t *testing.T) {
	t.Skip()

	wrapper := Wrapper{Person{name: "Will"}}
	assert.Equal(t, Ingredient{name: "🍙 nori"}, wrapper.work())
}
