package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	strs, err := util.ReadBytes("1-3", false)
	if err != nil {
		panic(err)
	}

	sum := 0
	for pairs := range util.ChunkSeq(strs, 3) {
		enemyCount := 3
		for _, char := range pairs {
			switch char {
			case 'B':
				sum++
			case 'C':
				sum += 3
			case 'D':
				sum += 5
			case 'x':
				enemyCount--
			}
		}

		switch enemyCount {
		case 3:
			sum += 6
		case 2:
			sum += 2
		}
	}

	fmt.Println(sum)
}
