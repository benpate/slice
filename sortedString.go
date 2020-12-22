package slice

import "sort"

// AddUnique expected a sorted string slice.  It returns a new slice that contains
// the unique value.
func AddUnique(s []string, value string) []string {

	index := sort.SearchStrings(s, value)

	if index == len(s) {
		return append(s, value)
	}

	if s[index] == value {
		return s
	}

	result := append(s[:index+1], s[index:]...)
	result[index] = value

	return result
}

// Remove removes a value from the []string
func Remove(s []string, value string) []string {

	index := sort.SearchStrings(s, value)

	if index == len(s) {
		return s
	}

	if s[index] != value {
		return s
	}

	return append(s[:index], s[index+1:]...)
}

// Identical returns TRUE if the two []string contain the same values in the same order.  FALSE otherwise.
func Identical(s1 []string, s2 []string) bool {

	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// Contains returns TRUE if this StringSlice contains the provided value
func Contains(s []string, value string) bool {

	// Find the index where the value SHOULD be
	index := sort.SearchStrings(s, value)

	// Return TRUE if the value is actually there
	return (index < len(s)) && (s[index] == value)
}

// ContainsAll returns TRUE if s1 contains every value in s2
func ContainsAll(s1 []string, s2 []string) bool {

	for _, v := range s2 {
		if Contains(s1, v) == false {
			return false
		}
	}

	return true
}
