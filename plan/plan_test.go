package plan

import "testing"
import "github.com/stretchr/testify/assert"

func TestNew(t *testing.T) {
	p := New()
	assert.NotNil(t, p)
}

func TestGenerate(t *testing.T) {
	p := New()
	assert.Nil(t, p.ReadDir("../examples/"))
	assert.Nil(t, p.Generate())
}
