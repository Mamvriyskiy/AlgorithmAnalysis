package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s1, s2 string
	_, err := fmt.Fscan(reader, &s1)
	if err != nil {
		os.Exit(-1)
	}

	_, err = fmt.Fscan(reader, &s2)
	if err != nil {
		os.Exit(-1)
	}

	r1 := []rune(s1)
	r2 := []rune(s2)

	//матричный алгоритм Левенштейна
	matrixLevenshtein(r1, r2)
}

