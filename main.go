package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Contains функция возвращет true, если элемент содержится в массиве
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// Find функция возвращает индекс элемента в массиве
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// RomanToArabic перевод римских чисел (от 1(I) до 10(X)) в арабские
func RomanToArabic(x string) (int, bool) {
	romanNums := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	if Contains(romanNums, x) {
		return Find(romanNums, x) + 1, true
	} else {
		return -1, false
	}
}

// getNumber возращает число и статус преобразования из строки
// если строка содержала арабское число, то функция возвращет это самое число и статус false (т.е. не было преобразования)
// если строка содержала римское число, то - конвертированное римское в арабское число и статус true
// если функция не смогла провести никакое преобраование, то возвращает -1 и статус false
func getNumber(x string) (int, bool) {
	num, err := strconv.Atoi(x)
	if err == nil {
		return num, false
	} else {
		return RomanToArabic(x)
	}
}

// ArabicToRoman преобразование арабского числа в римское (от 1(I) до 100(С))
func ArabicToRoman(x int) string {
	romanNums := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	hundreds := 0
	tens := 0
	units := 0

	hundreds = x / 100
	tens = (x - hundreds*100) / 10
	units = x - hundreds*100 - tens*10

	if hundreds == 1 {
		return "C"
	}

	if tens == 9 {
		return "XC" + romanNums[units]
	}

	if tens < 9 && tens >= 5 {
		return "L" + strings.Repeat("X", tens%5) + romanNums[units]
	}

	if tens == 4 {
		return "XL" + romanNums[units]
	}

	if tens < 4 && tens > 1 {
		return strings.Repeat("X", tens) + romanNums[units]
	}

	if tens < 1 && hundreds < 1 {
		return romanNums[units]
	}

	return ""
}

func main() {

	var mathRes int

	// Консольный ввод
	in := bufio.NewScanner(os.Stdin)
	fmt.Print("INPUT: ")
	in.Scan()
	line := in.Text()

	x := make([]string, 3)
	symbols := strings.Split(line, " ")

	// Проверка корректности ввода по длине полученного массива
	if len(symbols) != 3 {
		fmt.Println("Incorrect input")
		return

	} else {
		i := 0
		for _, symbol := range symbols {
			if symbol != " " {
				x[i] = symbol
				i += 1
			}
		}

		operand1 := x[0]
		operand2 := x[2]
		operation := x[1]

		// Проверка корректности ввода математическго оператора
		mathOperations := []string{"+", "-", "*", "/"}
		if Contains(mathOperations, operation) == false {
			fmt.Println("Indefinite math operation")
			return

		} else {

			// преобразование введённых чисел string -> int
			// и получение статуса преобразования
			int1, status1 := getNumber(operand1)
			int2, status2 := getNumber(operand2)

			// если оба числа одного представления и если числа корректны,
			// то выполняем математичсекие вычисления
			if status1 == status2 && (int1 > -1 && int2 > -1) {
				switch operation {
				case "+":
					mathRes = int1 + int2
				case "-":
					mathRes = int1 - int2
				case "*":
					mathRes = int1 * int2
				case "/":
					mathRes = int1 / int2
				}
			}

			// ошибка преобразования строки в число
			if int1 == -1 || int2 == -1 {
				fmt.Println("Indefinite operand")
				return
			}

			// числа различного представления (арабское и римское)
			if status1 != status2 {
				fmt.Println("Please use only arabic or only roman numbers")
				return
			}

			// успешно (арабские числа)
			if status1 == false && status2 == false {
				fmt.Printf("OUTPUT: %d", mathRes)
				return
			}

			// успешно (римские числа)
			if (status1 == true && status2 == true) && mathRes > 0 {
				fmt.Printf("OUTPUT: %s", ArabicToRoman(mathRes))
				return
			}

			// ошибка: при введёных римских числах получили неположительный результат вычисления
			if (status1 == true && status2 == true) && mathRes <= 0 {
				fmt.Println("Non-positive calculation result not available for Roman numerals")
				return
			}
		}
	}
}
