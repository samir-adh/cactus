package game

type GameRun struct {
	State *GameState
}

func NewGameRun() *GameRun {
	return &GameRun{
		State: NewGameState(5),
	}
}

func (gr *GameRun) Start() {
	print("Starting a new game...\n")
	currentPlayer := gr.State.CurrentPlayer
	print("Starting game with player A's turn.\nPlayer A's hand: " + currentPlayer.String() + "\n")
	gr.PlayerTurn()
}

func (gr *GameRun) PlayerTurn() {
	state := gr.State
	print("Player A pick a card from the deck.\n")
	state.DrawCardFromDeck(state.CurrentPlayer, true)
	print("Player A's hand: " + state.CurrentPlayer.String() + "\n")
	print("Last card in the discard pile: " + Last(state.DiscardPile).String() + "\n")
}
