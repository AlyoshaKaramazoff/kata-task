package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func romanToArabic(x string) (int, bool) {
	romanNums := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	if Contains(romanNums, x) {
		return Find(romanNums, x) + 1, true
	} else {
		return -1, false
	}
}

func getNumber(x string) (int, bool) {
	num, err := strconv.Atoi(x)
	if err == nil {
		return num, false
	} else {
		return romanToArabic(x)
	}
}

func arabicToRoman(x int) string {
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

	in := bufio.NewScanner(os.Stdin)
	fmt.Print("INPUT: ")
	in.Scan()
	line := in.Text()

	x := make([]string, 3)
	symbols := strings.Split(line, " ")

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

		mathOperations := []string{"+", "-", "*", "/"}
		if Contains(mathOperations, operation) == false {
			fmt.Println("Indefinite math operation")
			return

		} else {

			int1, status1 := getNumber(operand1)
			int2, status2 := getNumber(operand2)

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

			if int1 == -1 || int2 == -1 {
				fmt.Println("Indefinite operand")
				return
			}
			if status1 != status2 {
				fmt.Println("Please use only arabic or only roman numbers")
				return
			}

			if status1 == false && status2 == false {
				fmt.Printf("OUTPUT: %d", mathRes)
				return
			}

			if (status1 == true && status2 == true) && mathRes > 0 {
				fmt.Printf("OUTPUT: %s", arabicToRoman(mathRes))
				return
			}

			if (status1 == true && status2 == true) && mathRes <= 0 {
				fmt.Println("Non-positive calculation result not available for Roman numerals")
				return
			}
		}
	}
}
