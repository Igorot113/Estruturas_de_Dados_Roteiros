package main

import "fmt"

type Stack struct {
	items []rune
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	lastIndex := len(s.items) - 1
	items := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return items, true
}

func isPalindromo(word string) bool {
	stack := Stack{}

	for _, caractere := range word {
		stack.Push(caractere)
	}

	for _, caractere := range word {
		top, _ := stack.Pop()
		if top != caractere {
			return false
		}
	}
	return true
}
func main() {
	words := []string{"radar", "arara", "golang", "level", "hello"}
	for _, word := range words {
		if isPalindromo(word) {
			fmt.Printf("\"%s\" é um palindromo\n", word)
		} else {
			fmt.Printf("\"%s\" NÃO é um palíndromo\n", word)
		}
	}
}
