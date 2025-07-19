package game

import "math/rand"

type GameState struct {
	HandA         Hand
	HandB         Hand
	CurrentPlayer *Hand // 0 for player A, 1 for player B
	DiscardPile   []*Card
	CardDeck      []*Card
}

func NewGameState(initialNumberOfCards int) *GameState {
	state := &GameState{
		HandA:         Hand{Cards: []*Card{}, KnownCards: []*Card{}},
		HandB:         Hand{Cards: []*Card{}, KnownCards: []*Card{}},
		CurrentPlayer: nil, // Start with player A
		DiscardPile:   []*Card{},
		CardDeck:      NewDeck(),
	}
	state.CurrentPlayer = &state.HandA // Start with player A
	// Draw initial cards for both players
	for i := 0; i < initialNumberOfCards; i++ {
		if i < 2 {
			state.DrawCardFromDeck(&state.HandA, true)
			state.DrawCardFromDeck(&state.HandB, true)
		} else {
			state.DrawCardFromDeck(&state.HandA, false)
			state.DrawCardFromDeck(&state.HandB, false)
		}
	}
	var card *Card
	card, state.CardDeck = Pop(state.CardDeck)
	state.DiscardPile = append(state.DiscardPile, card)
	return state
}

func ShuffleCards(cards []*Card) {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

func NewDeck() []*Card {
	ranks := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	var cards []*Card
	for _, rank := range ranks {
		for _, suit := range suits {
			cards = append(cards, &Card{Rank: rank, Suit: suit})
		}
	}
	// Shuffle the cards
	ShuffleCards(cards)
	// Return the shuffled cards
	return cards
}

func ResetCardDeck(gs *GameState) {
	// Put all cards from the discard pile back into the deck except the last one
	gs.CardDeck = append(gs.CardDeck, gs.DiscardPile[:len(gs.DiscardPile)-1]...)
	ShuffleCards(gs.CardDeck)
	// Clear the discard pile except the last card
	gs.DiscardPile = gs.DiscardPile[len(gs.DiscardPile)-1:]
}

func Pop(cards []*Card) (*Card, []*Card) {
	if len(cards) == 0 {
		return nil, cards
	}
	card := cards[len(cards)-1]
	cards = cards[:len(cards)-1]
	return card, cards
}

func Last(cards []*Card) *Card {
	if len(cards) == 0 {
		return nil
	}
	return cards[len(cards)-1]
}

func (gs *GameState) DrawCardFromDeck(hand *Hand, reveal bool) {
	if len(gs.CardDeck) == 0 {
		ResetCardDeck(gs)
	}
	card, updatedDeck := Pop(gs.CardDeck)
	gs.CardDeck = updatedDeck
	hand.AddCard(card)
	if reveal {
		hand.RevealCard(len(hand.Cards) - 1)
	}
}

func (gs *GameState) DiscardCard(hand *Hand, index int) {
	lastDiscarded := Last(gs.DiscardPile)
	toDiscard := hand.Cards[index]
	hand.RemoveCard(toDiscard)
	gs.DiscardPile = append(gs.DiscardPile, toDiscard)
	if lastDiscarded.Rank != toDiscard.Rank {
		gs.DrawCardFromDeck(hand, false)
	}
}
