package main

import (
	"errors"
	"fmt"
	"os"
)

// Познакомьтесь с алгоритмом сортировки вставками.
// Напишите приложение, которое принимает на вход набор целых чисел
// и отдаёт на выходе его же в отсортированном виде.
// Сохраните код, он понадобится нам в дальнейших уроках.

func sort(s []int) {
	if len(s) < 2 {
		return
	}
	for i := range s {
		k := i
		for j := i; j < len(s); j++ {
			if s[k] > s[j] {
				k = j
			}
		}
		s[i], s[k] = s[k], s[i]
	}
}

func main() {
	var size int
	fmt.Print("Enter size of sequence:")
	if _, err := fmt.Scanln(&size); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if size < 1 {
		fmt.Println(errors.New("size should be positive and greather than 0"))
		os.Exit(1)
	}
	sequence := make([]int, size)
	fmt.Print("Enter sequence of integers:")
	for i := range sequence {
		fmt.Scan(&sequence[i])
	}
	sort(sequence)
	fmt.Println("sorted sequence =", sequence)
}
