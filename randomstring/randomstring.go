package randomstring

import "math/rand"

//Generator struct
type Generator struct {
	availableChars string
	length         int
}

//NewGenerator generates random string generator instance
func NewGenerator(availableChars string, length int) *Generator {
	if len(availableChars) == 0 || length == 0 {
		panic("availableChars cannot be empty, length must be greater than 0")
	}
	return &Generator{
		availableChars: availableChars,
		length:         length,
	}
}

//Generate generates random string
func (g *Generator) Generate() string {
	r := []byte{}
	for i := 0; i < g.length; i++ {
		r = append(r, g.availableChars[rand.Intn(len(g.availableChars))])
	}
	return string(r)
}
