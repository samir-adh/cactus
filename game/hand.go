package game

type Hand struct {
	Cards      []*Card
	KnownCards []*Card // Cards that have been revealed
}

/*AddCard adds a card to the deck.*/
func (d *Hand) AddCard(card *Card) {
	d.Cards = append(d.Cards, card)
}

/*
RemoveCard removes a card from the deck.
It searches for the card in the Cards slice and removes it if found.
If the card is not found, it does nothing.
It also removes the card from KnownCards if it exists there.
*/
func (d *Hand) RemoveCard(card *Card) {

	// Remove the card from Cards slice
	for i, c := range d.Cards {
		if c == card {
			d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
			return
		}
	}
	// If the card is in KnownCards, remove it from there as well
	for i, c := range d.KnownCards {
		if c == card {
			d.KnownCards = append(d.KnownCards[:i], d.KnownCards[i+1:]...)
			return
		}
	}
}

/*
RevealCard reveals a card in the deck.
It searches for the card in the Cards slice and adds it to KnownCards if found.
*/
func (d *Hand) RevealCard(index int) *Card {
	card := d.Cards[index]
	d.KnownCards = append(d.KnownCards, card)
	return card
}

func (d *Hand) CardIsKnown(index int) bool {
	card := d.Cards[index]
	for _, knownCard := range d.KnownCards {
		if knownCard == card {
			return true
		}
	}
	return false
}

func (d *Hand) String() string {
	result := "Hand:\n"
	for i, card := range d.Cards {
		if d.CardIsKnown(i) {
			result += "[Known] "
		} else {
			result += "[Unknown] "
		}
		result += card.String() + "\n"
	}
	return result
}

func (d *Hand) TotalValue() int {
	total := 0
	for _, card := range d.Cards {
		total += card.Value()
	}
	return total
}