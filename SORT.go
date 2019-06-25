package main

import "fmt"

func main() {
	var i, j, k, g int
	var m string
	var ms [3][3]int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Println("Введите элемент массива")
			fmt.Print(i + 1)
			fmt.Print(" Строки ")
			fmt.Scan(&ms[i][j])
		}
	}
	fmt.Println(ms)
	fmt.Println("Введите какую сортировку вы хотите применить к двумерному массиву(Возрастание/Убывание)")
	fmt.Scan(&m)
	g = ms[1][1]
	switch m {
	case "Возрастание":
		for i = 0; i < 3; i++ {
			for j = 0; j < 3; j++ {
				for k = 0; k < 3; k++ {
					if ms[i][j] < ms[i][k] {
						g = ms[i][j]
						ms[i][j] = ms[i][k]
						ms[i][k] = g
					}
				}
			}
		}
	case "Убывание":
		for i = 0; i < 3; i++ {
			for j = 0; j < 3; j++ {
				for k = 0; k < 3; k++ {
					if ms[i][j] > ms[i][k] {
						g = ms[i][j]
						ms[i][j] = ms[i][k]
						ms[i][k] = g
					}
				}
			}
		}
	}
	fmt.Println("Отсортированный массив:")
	// i = 0; i < 3; i++ {
	//	for j = 0; j < 3; j++ {
	fmt.Println(ms)
	//	}
	//}
}

// C:\Go\bin\go.exe run C:\Users\Работа\Desktop\go\GOLANG.go
