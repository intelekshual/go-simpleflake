package simpleflake

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	id, err := New()
	assert.Nil(err, "Error should be nil")
	assert.NotNil(id, "Id should not be nil")
}

func TestParse(t *testing.T) {
	// TODO: add more test cases
	testCases := []struct {
		id  uint64
		ts  uint64
		seq uint64
	}{
		{3884649614048016403, 1409771120644, 3192851},
		{3884626385633042485, 1409768351601, 4481077},
	}

	for _, tc := range testCases {
		id := Parse(tc.id)
		assert.Equal(t, id[0], tc.ts, "Timestamps should match")
		assert.Equal(t, id[1], tc.seq, "Random sequences should match")
	}
}
func TestSetCustomPrecision(t *testing.T) {
	// TODO: write test
}
func TestSetCustomEpoch(t *testing.T) {
	// TODO: write test
}
func TestBuildId(t *testing.T) {
	// TODO: write test
}
func TestCustomTimestamp(t *testing.T) {
	// TODO: write test
}
