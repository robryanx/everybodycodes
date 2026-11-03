package util

import "iter"

func ChunkSeq[T any](seq iter.Seq[T], chunk int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		var curr []T
		for next := range seq {
			curr = append(curr, next)
			if len(curr) == chunk {
				if !yield(curr) {
					return
				}
				curr = []T{}
			}
		}

		if len(curr) > 0 {
			if !yield(curr) {
				return
			}
		}
	}
}
