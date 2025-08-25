package ehex

import (
	"fmt"
)

// Ehex represents a value in extended hexadecimal system (0-Z excluding O and I)
// It encapsulates an integer value between 0 and 33 inclusive
type Ehex struct {
	value int
}

// toChar maps integer values to their corresponding extended hex characters
var toChar = map[int]string{
	0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9",
	10: "A", 11: "B", 12: "C", 13: "D", 14: "E", 15: "F", 16: "G", 17: "H", 18: "J",
	19: "K", 20: "L", 21: "M", 22: "N", 23: "P", 24: "Q", 25: "R", 26: "S", 27: "T",
	28: "U", 29: "V", 30: "W", 31: "X", 32: "Y", 33: "Z",
}

// toNum maps extended hex characters to their corresponding integer values
var toNum = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
	"A": 10, "B": 11, "C": 12, "D": 13, "E": 14, "F": 15, "G": 16, "H": 17,
	"J": 18, "K": 19, "L": 20, "M": 21, "N": 22,
	"P": 23, "Q": 24, "R": 25, "S": 26, "T": 27,
	"U": 28, "V": 29, "W": 30, "X": 31, "Y": 32, "Z": 33,
}

// String converts an integer to its extended hex representation
// Panics if the integer is outside the valid range [0, 33]
func String(i int) string {
	if _, ok := toChar[i]; !ok {
		panic(fmt.Sprintf("%v is out of bound for ehex code", i))
	}
	return toChar[i]
}

// Int converts an extended hex character to its integer value
// Panics if the string is not a valid extended hex character
func Int(s string) int {
	if _, ok := toNum[s]; !ok {
		panic(fmt.Sprintf("'%v' is out of bound for ehex index", s))
	}
	return toNum[s]
}

// FromInt creates an Ehex instance from an integer value
// Panics if the integer is outside the valid range [0, 33]
func FromInt(n int) Ehex {
	if n < 0 || n > 33 {
		panic(fmt.Sprintf("ExtendedHex value out of range [0, 33]: %d", n))
	}
	return Ehex{value: n}
}

// FromString creates an Ehex instance from a string representation
// Panics if the string is not a valid extended hex character
func FromString(s string) Ehex {
	if num, exists := toNum[s]; exists {
		return Ehex{value: num}
	}
	panic(fmt.Sprintf("Invalid ExtendedHex character: %s", s))
}

// Int returns the integer value of the Ehex instance
func (eh Ehex) Int() int {
	return eh.value
}

// String returns the string representation of the Ehex instance
func (eh Ehex) String() string {
	return toChar[eh.value]
}

// Difference returns the absolute difference between two Ehex values
// The result is always a positive integer between 0 and 33
func Difference(e1, e2 Ehex) int {
	diff := e1.value - e2.value
	if diff < 0 {
		return -diff
	}
	return diff
}
