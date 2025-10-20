package fetch_test

import (
	"github.com/njhoffman/i3-tree/pkg/fetch"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromFake(t *testing.T) {
	f := fetch.FromFake{}
	got, gotErr := f.Fetch()

	assert.Nil(t, gotErr)
	assert.NotNil(t, got)
}
