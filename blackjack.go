package main

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
