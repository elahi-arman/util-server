package groupings

import (
	"math/rand"
)

func RandomEvenLists(values []string, groups int) [][]string {
	// pre-emptively exit in edge cases
	if groups <= 1 {
		return [][]string{values}
	}

	// initialize lists
	capacity := len(values) / groups
	lists := make([][]string, groups)
	for i := range lists {
		lists[i] = make([]string, 0, capacity)
	}

	// len(values) may not be evenly divisible by groups, so the remaining
	// capacity will go to the first {leftOver} lists
	leftOver := len(values) % groups
	for i := 0; i < leftOver; i++ {
		lists[i] = make([]string, 0, capacity+1)
	}

	// shuffle the elements
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	// fill up each list with a subsection {values} until its reached capacity
	offset := 0
	for i := 0; i < len(lists); i++ {
		for j := 0; j < cap(lists[i]); j++ {
			lists[i] = append(lists[i], values[j+offset])
		}

		// update the offset so we don't keep repeating the
		// same elements
		offset += cap(lists[i])
	}

	return lists
}
