package main

import (
	"fmt"
	"math"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
	fmt.Println(part3())
}

func part1() int {
	numSeq, err := util.ReadInts("1-1", false, ',')
	if err != nil {
		panic(err)
	}

	count := 0
	for num := range numSeq {
		for i := 1; ; i++ {
			candidate := i * i
			if candidate == num {
				break
			} else if candidate > num {
				count += candidate - num
				break
			}
		}
	}

	return count
}

func part2() int {
	numSeq, err := util.ReadInts("1-2", false, ',')
	if err != nil {
		panic(err)
	}

	count := 0
	for num := range numSeq {
		for i := 3; ; i++ {
			candidate := i*4 - 4
			if candidate == num {
				break
			} else if candidate > num {
				count += candidate - num
				break
			}
		}
	}

	return count
}

func part3() int {
	numSeq, err := util.ReadInts("1-3", false, ',')
	if err != nil {
		panic(err)
	}

	count := 0
	for num := range numSeq {
		val := int(math.Ceil(math.Sqrt(float64(num))))
		if val%2 == 0 {
			val++
		}
		if val*val == num {
			val += 2
		}

		candidate := val*val - 1
		if candidate > num {
			count += candidate - num
		}
	}

	return count
}
