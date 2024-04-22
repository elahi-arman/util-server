package groupings_test

import (
	"testing"

	"github.com/elahi-arman/util-server/internal/groupings"
	"github.com/stretchr/testify/assert"
)

func TestRandomEvenLists(t *testing.T) {
	defaultValues := []string{
		"a", "b", "c", "d",
		"e", "f", "g", "h",
		"i", "j", "k", "l",
	}

	type testCase struct {
		name          string
		values        []string
		groups        int
		expectedSizes []int
	}

	cases := []testCase{
		{
			name:          "should return wrapped inputs when groups is less than or equal to 1",
			groups:        1,
			values:        defaultValues,
			expectedSizes: []int{12},
		},
		{
			name:          "should return wrapped inputs when only 1 value in input list",
			groups:        4,
			values:        []string{"lonely"},
			expectedSizes: []int{1, 0, 0, 0},
		},
		{
			name:          "should split lists percetly when even",
			groups:        4,
			values:        defaultValues,
			expectedSizes: []int{3, 3, 3, 3},
		},
		{
			name:          "should split allocate extra capacity to the first few lists when not even",
			groups:        5,
			values:        defaultValues,
			expectedSizes: []int{3, 3, 2, 2, 2},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := groupings.RandomEvenLists(tc.values, tc.groups)
			assert.Len(t, actual, len(tc.expectedSizes))
			for i, l := range actual {
				assert.Len(t, l, tc.expectedSizes[i])
			}
		})
	}
}
