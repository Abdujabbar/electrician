package randomstring

import (
	"fmt"
	"strings"
	"testing"
)

func TestRandomStringGeneratorEmptyParams(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	rnd := NewGenerator("", 0)
	t.Error("Error on checking params")
	fmt.Println(rnd)
}

func TestRandomStringGeneration(t *testing.T) {
	rnd := NewGenerator("abc", 10)
	rndString := rnd.Generate()
	if len(rndString) != 10 {
		t.Error("Error while generating string with passed length param")
	}
}

func TestRandomStringAvailableChars(t *testing.T) {
	rnd := NewGenerator("abc", 5)
	rndString := rnd.Generate()
	if strings.Contains(rndString, "z") {
		t.Error("Generated char with not available symbol")
	}
}
