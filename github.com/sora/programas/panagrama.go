package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var numFrases int
	fmt.Scanln(&numFrases)

	reader := bufio.NewReader(os.Stdin)

	//for _, tc := range []string{
	for i := 0; i < numFrases; i++ {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)
		/*	"abcdefGHIJKLMNopqrstuvwxyz",
			"abcdefghijklwnopqrstuvwxyz",
			"Este es un texto de ejemplo",
			"The quick brown fox jumps over the lazy dog",
			"El veloz murcielago hindu comia feliz cardillo y kiwi. La ciguena tocaba el saxofon detras del palenque de paja",
			"El viejo Señor Gómez pedía queso, kiwi y habas, pero le ha tocado un saxo",
			"Quiere la boca exhausta vid, kiwi, pina y fugaz jamon",*/
		//} {
		if isPangram(text) {
			fmt.Println("SI")
		} else {
			fmt.Println("NO")
		}
	}
}

func isPangram(s string) bool {
	mm := mm()
	for _, r := range []rune(s) {
		if unicode.IsLetter(r) {
			t := unicode.ToLower(r)
			if _, ok := mm[t]; ok {
				mm[t] = true
			} else {
				return false
			}
		}
	}
	for _, v := range mm {
		if !v {
			return false
		}
	}
	return true
}

func mm() map[rune]bool {
	m := make(map[rune]bool, 28)
	for i := 'a'; i <= 'z'; i++ {
		m[i] = false
	}
	return m
}
