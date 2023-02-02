package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	buf := bufio.NewScanner(os.Stdin) //считываем ввод
	buf.Scan()
	text := buf.Text() // Преобразуем ввод в строку

	replacer := strings.NewReplacer("\r", "", "\n", "") // Замена символов пробелов в строке
	text = replacer.Replace(text)

	text1 := strings.Fields(text) // Разделение строки на массив строк

	if len(text1) > 3 { // если длина массива больше 3,вывод ошибки
		log.Print("Ошибка,два операнда и один оператор (+, -, /, *)")
		return
	}

	if len(text1) < 3 { // если длина массива меньше 3, вывод ошибки
		log.Print("Не является математической операцией")
		return
	}

	firstNumRoman := romanToInt(text1[0]) //конвертация римских чисел в десятичную
	secondNumRoman := romanToInt(text1[2])

	firstNumArabic, _ := strconv.Atoi(text1[0]) //конвертация string в int
	secondNumArabic, _ := strconv.Atoi(text1[2])

	if text1[1] != "/" && text1[1] != "*" && text1[1] != "+" && text1[1] != "-" { // если 2й эелемент массива не математический знак
		fmt.Println("Ошибка операнда! (+, -, /, *)")
		return
	}

	if firstNumRoman != 0 && secondNumRoman != 0 { //Если конвертация в римские числа прошла успешно
		result := operate(firstNumRoman, secondNumRoman, text1[1])
		if result < 0 { //Если результат отрицательный
			fmt.Println("Ошибка! В римской системе нет отрицательных чисел")
			return
		} else {
			fmt.Println(intToRoman(result))
			return
		}
	}

	if firstNumArabic > 0 && firstNumArabic < 10 && secondNumArabic > 0 && secondNumArabic < 10 { //Если арабские цифры удовлетворяют диапазону по ТЗ
		fmt.Println(operate(firstNumArabic, secondNumArabic, text1[1]))
	} else {
		if firstNumArabic < 0 || firstNumArabic > 10 || secondNumArabic < 0 || secondNumArabic > 10 {
			fmt.Println("Ошибка! На входе должны быть числа от 1 до 10 включительно")
		} else {
			fmt.Println("Ошибка! Используются одновременно разные системы счисления")
		}
	}
}

func romanToInt(s string) int { //Конвертация римского числа в десятичное(найдена в интернете)
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}

func intToRoman(num int) string { //конвертация десятичного чила в строковое представление римского
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}
	return roman
}
func operate(a, b int, s string) int { //Математическая операция на основе арифметического знака
	switch s {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}
