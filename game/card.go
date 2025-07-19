package game

import "strconv"

type Card struct {
	Rank string
	Suit string
}



func (c Card) String() string {
	return c.Rank + " of " + c.Suit
}

func (c Card) Value() int {
	switch c.Rank {
	case "Ace":
		return 1
	case "King":
		if c.Suit == "Hearts" {
			return 0
		}
		return 11
	case "Queen", "Jack":
		return 10
	default:
		value, error := strconv.Atoi(c.Rank)
		if error != nil {
			panic("failed to get value of an unsupported card") // Return error if conversion fails
		}
		return value
	}
}
