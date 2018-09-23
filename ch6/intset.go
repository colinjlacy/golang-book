package main

type IntSet struct {
	words []uint64
}

func (is *IntSet) Has(x int) bool {
	word, bit := x / 64, uint(x % 64)
	return word < len(is.words) && is.words[word]&(1<<bit) != 0
}

func (is *IntSet) Add(x int) {
	word, bit := x / 64, uint(x % 64)
	for word >= len(is.words) {
		is.words = append(is.words, 0)
	}
	is.words[word] |= 1 << bit
}

func main() {}
