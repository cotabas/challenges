package main

import (
	"fmt"
)

func applySkipHopJumpPattern(input string, skip, hop, jump int) string {
	var result string

	for i := 0; i < len(input); {
		// Apply Skip
		i += skip
		if i >= len(input) {
			break
		}
		result += string(input[i])

		// Apply Hop
		i += hop
		if i >= len(input) {
			break
		}
		result += string(input[i])

		// Apply Jump
		i += jump
	}

	return result
}

func main() {

  input := "Fasten your seatbelts. It's going to be a bumpy night."

  fmt.Println("Original:", input)
  for skip := 1; skip <= 10; skip++ {
    for hop := 1; hop <= 10; hop++ {
      for jump := 1; jump <= 10; jump++ {
        result := applySkipHopJumpPattern(input, skip, hop, jump)
        fmt.Println("Skip Hop Jump:", result)
      }
    }
  }
}

