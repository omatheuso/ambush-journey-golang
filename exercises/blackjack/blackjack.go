package blackjack

// 1. Calculate the value of any given card.
// Implement a function to calculate the numerical value of a card:
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
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
	case "ten":
		fallthrough
	case "jack":
		fallthrough
	case "queen":
		fallthrough
	case "king":
		return 10
	default:
		return 0
	}
}

// 2. Implement the decision logic for the first turn.
// Write a function that implements the decision logic as described above:
func FirstTurn(card1, card2, dealerCard string) string {
	switch {
	// acesHand
	case card1 == "ace" && card2 == "ace":
		return "P"
	case SumHand(card1, card2) == 21:
		return DealerHandWhenJacked(dealerCard)
	case 17 <= SumHand(card1, card2) && SumHand(card1, card2) <= 20:
		return "S"
	case 12 <= SumHand(card1, card2) && SumHand(card1, card2) <= 16 && ParseCard(dealerCard) >= 7:
		return "H"
	case 12 <= SumHand(card1, card2) && SumHand(card1, card2) <= 16 && ParseCard(dealerCard) <= 7:
		return "S"
	default:
		return "H"
	}
}

func DealerHandWhenJacked(dealerCard string) string {
	switch {
	case ParseCard(dealerCard) >= 10:
		return "S"
	default:
		return "W"
	}
}

func SumHand(card1, card2 string) int {
	return ParseCard(card1) + ParseCard(card2)
}
