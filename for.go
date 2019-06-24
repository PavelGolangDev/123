package main

import (
	"fmt"
)

func main() {
	var i, min int
	var ms [10]int
	for i = 0; i <= 9; i++ { // В цикле вводим массив
		fmt.Println("Введите элемент массива")
		fmt.Scan(&ms[i])
	}
	i = 0
	min = ms[0]              // Присваиваем минимальному первый элемент массива
	for i = 0; i <= 9; i++ { // Проходимся циклом по массиву и находим минимальный элемент
		if ms[i] < min {
			min = ms[i]
		}
	}
	fmt.Println("Минимальное число в массиве")
	fmt.Println(min)
}
