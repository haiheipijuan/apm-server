package error

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pr "github.com/elastic/apm-server/processor"
)

func TestImplementProcessorInterface(t *testing.T) {
	p := NewProcessor()
	assert.NotNil(t, p)
	_, ok := p.(pr.Processor)
	assert.True(t, ok)
	assert.IsType(t, &processor{}, p)
}

func TestAddProcessorToRegistryOnInit(t *testing.T) {
	p := pr.Registry.GetProcessor("/v1/errors")
	assert.NotNil(t, p)
}
