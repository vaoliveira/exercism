package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var cardSum int = ParseCard(card1) + ParseCard(card2)
	var dealerValue int = ParseCard(dealerCard)
	switch {
	case cardSum == 22:
		return "P"
	case cardSum == 21 && dealerValue < 10:
		return "W"
	case (cardSum == 21 && dealerValue >= 10) || (cardSum >= 17 && cardSum <= 20) || (cardSum >= 12 && cardSum <= 16 && dealerValue < 7):
		return "S"
	default:
		return "H"
	}
}
