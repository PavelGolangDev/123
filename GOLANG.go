package main

import (
	"fmt" 
)

func main() {
	var a int
	var b int
	var c int
	var q int
			fmt.Println("Введите q") 
			fmt.Scan(&q)               // Ввод переменной с консоли, оператор аналогичен fscan.
			fmt.Println("Введите b")
			fmt.Scan(&b)
			fmt.Println("Введите а") // Вывод на экран.
			fmt.Scan(&a)
	c = a + b // Сумма переменных, введенных с консоли.
	fmt.Println(c)
	if c < q {  // Сравнивание полученной суммы с переменной, введенной с консоли.
		fmt.Println("+")
	} else {
		fmt.Println("-")
	}
}


