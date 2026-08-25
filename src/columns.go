package main

import (
	"math/rand/v2"
	"slices"
)

func get_new_col_prefer_unused(max int) int {
	ret := 0
	var avail []int
	for i := 0; i <= max; i++ {
		if !slices.Contains(config.active_column_indexes, i) {
			avail = append(avail, i)
		}
	}
	if len(avail) == 0 {
		ret = rand.IntN(max + 1)
	} else {
		rand_idx := rand.IntN(len(avail))
		ret = avail[rand_idx]
	}
	return ret
}
