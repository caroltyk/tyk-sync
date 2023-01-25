package examplesrepo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHasExamples(t *testing.T) {
	t.Run("should return false if index is nil", func(t *testing.T) {
		assert.False(t, IndexHasExamples(nil))
	})

	t.Run("should return false if no examples are present", func(t *testing.T) {
		index := RepositoryIndex{}
		assert.False(t, IndexHasExamples(&index))
	})

	t.Run("should return true if index has at least one example", func(t *testing.T) {
		index := RepositoryIndex{
			Examples: ExamplesCategories{
				UDG: []ExampleMetadata{
					{
						Location: "location",
					},
				},
			},
		}

		assert.True(t, IndexHasExamples(&index))
	})
}

func TestMergeExamples(t *testing.T) {
	t.Run("should return empty slice when index is nil", func(t *testing.T) {
		examples := MergeExamples(nil)
		assert.Len(t, examples, 0)
	})

	t.Run("should merge examples successfully", func(t *testing.T) {
		udgExample := ExampleMetadata{
			Location: "udg",
		}

		index := RepositoryIndex{
			Examples: ExamplesCategories{
				UDG: []ExampleMetadata{
					udgExample,
				},
			},
		}

		expectedExamples := []ExampleMetadata{
			udgExample,
		}

		examples := MergeExamples(&index)
		assert.Equal(t, expectedExamples, examples)
	})
}
