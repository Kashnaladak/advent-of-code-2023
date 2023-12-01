package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Opening input file...")
	dat, err := os.ReadFile("input2.txt")

	if err != nil {
		panic(fmt.Sprintf("Couldn't read the input file: %s", err))
	}

	fmt.Println("Computing result...")

	res := 0
	isFwdSeeking := true
	isBwdSeeking := true
	var fwdByte byte = 0
	var bwdByte byte = 0

	for fwd := 0; fwd < len(dat); fwd++ {
		fwdByte = dat[fwd]

		if isFwdSeeking {
			// 48 = 0 ~ 57 = 9
			if fwdByte >= 48 && fwdByte <= 57 {
				// found a number
				res = res + ((int(fwdByte) - 48) * 10)
				isFwdSeeking = false
			}
		} else {
			if fwdByte == 10 {
				isFwdSeeking = true
			}
		}
	}

	for bwd := len(dat) - 1; bwd > 0; bwd-- {
		bwdByte = dat[bwd]

		if isBwdSeeking {
			// 48 = 0 ~ 57 = 9
			if bwdByte >= 48 && bwdByte <= 57 {
				// found a number
				res = res + (int(bwdByte) - 48)
				isBwdSeeking = false
			}
		} else {
			if bwdByte == 10 {
				isBwdSeeking = true
			}
		}
	}

	fmt.Printf("res is: %d\n", res)
}
