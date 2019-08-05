package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	channel := make(chan FreqMap, len(s))

	for _, r := range s {
		r := r
		go func() {
			channel <- Frequency(r)
		}()
	}

	fmap := make(FreqMap)
	for range s {
		for letter, freq := range <-channel {
			fmap[letter] += freq
		}
	}
	return fmap
}
