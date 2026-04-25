package main

func nextGreatestLetter(letters []byte, target byte) byte {
	lo, hi := 0, len(letters)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if letters[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if lo == len(letters) {
		return letters[0]
	}
	return letters[lo]
}

func main() {}
