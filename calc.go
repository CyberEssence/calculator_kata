package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		first_checker string = "Чекер"
		second_checker string = "Чекер"
		first_error    string = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
		second_error    string = "Вывод ошибки, так как строка не является математической операцией."
		third_error    string = "Вывод ошибки, так как используются одновременно разные системы счисления."
		fouth_error    string = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
		first_number  string
		second_number string
		symbol    string
	)

	fmt.Scanln(&first_number, &symbol, &second_number, &first_checker)

	var answer int = count_calculator(first_number, symbol, second_number)

	switch {
		case first_checker != second_checker:
			panic(first_error)

		case (check(symbol) != "symbol") || (check(first_number) != "rome" && check(first_number) != "arabic") || (check(second_number) != "rome" && check(second_number) != "arabic"):
			panic(second_error)

		case (check(first_number) == "rome" && check(second_number) == "arabic") || check(first_number) == "arabic" && check(second_number) == "rome":
			panic(third_error)

		case check(symbol) == "symbol" && check(first_number) == "rome" && check(second_number) == "rome" && answer < 1:
			panic(fouth_error)

		case check(symbol) == "symbol" && check(first_number) == "arabic" && check(second_number) == "arabic":
			fmt.Println(answer)

		case check(symbol) == "symbol" && check(first_number) == "rome" && check(second_number) == "rome" && answer > 0:
			fmt.Println(arabic_to_rome(answer))
	}

}

func check(check_variable string) string {
	arabic_number := [11]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	rome_number := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	check_symbol := [4]string{"+", "-", "*", "/"}

	var checker string

	for _, v := range arabic_number {
		if v == check_variable {
			checker = "arabic"
			break
		}
	}

	for _, v := range rome_number {
		if v == check_variable {
			checker = "rome"
			break
		}
	}

	for _, v := range check_symbol {
		if v == check_variable {
			checker = "symbol"
			break
		}
	}

	return checker
}

func count_calculator(first_number, symbol, second_number string) int {

	if check(first_number) == "rome" && check(second_number) == "rome" {
		first_number = rome_to_arabic(first_number)
		second_number = rome_to_arabic(second_number)
	}

	true_first_number, _ := strconv.Atoi(first_number)
	true_second_number, _ := strconv.Atoi(second_number)

	var result int

	switch {
		case symbol == "+":
			result = true_first_number + true_second_number
		case symbol == "-":
			result = true_first_number - true_second_number
		case symbol == "*":
			result = true_first_number * true_second_number
		case symbol == "/":
			result = true_first_number / true_second_number
	}

	return result
}

func rome_to_arabic(rome_number string) string {

	var result string

	switch rome_number {
		case "I":
			result = "1"
		case "II":
			result = "2"
		case "III":
			result = "3"
		case "IV":
			result = "4"
		case "V":
			result = "5"
		case "VI":
			result = "6"
		case "VII":
			result = "7"
		case "VIII":
			result = "8"
		case "IX":
			result = "9"
		case "X":
			result = "10"
	}

	return result
}

func arabic_to_rome(arabic_number int) string {

	var arabic_tens, arabic_ones int
	var rome_tens, rome_ones string
	var result string

	arabic_tens = arabic_number / 10
	arabic_ones = arabic_number % 10

	switch arabic_tens {
		case 1:
			rome_tens = "X"
		case 2:
			rome_tens = "XX"
		case 3:
			rome_tens = "XXX"
		case 4:
			rome_tens = "XL"
		case 5:
			rome_tens = "L"
		case 6:
			rome_tens = "LX"
		case 7:
			rome_tens = "LXX"
		case 8:
			rome_tens = "LXXX"
		case 9:
			rome_tens = "XC"
		case 10:
			rome_tens = "C"
	}

	switch arabic_ones {
		case 1:
			rome_ones = "I"
		case 2:
			rome_ones = "II"
		case 3:
			rome_ones = "III"
		case 4:
			rome_ones = "IV"
		case 5:
			rome_ones = "V"
		case 6:
			rome_ones = "VI"
		case 7:
			rome_ones = "VII"
		case 8:
			rome_ones = "VIII"
		case 9:
			rome_ones = "IX"
		case 0:
			rome_ones = ""
	}

	result = rome_tens + rome_ones

	return result
}
