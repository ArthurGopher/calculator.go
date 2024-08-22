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

	fmt.Print("Напишите простое математическое выражение: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)

	fmt.Println(input)

	switch {
	case strings.Contains(input, "1"):
		expressionArab(input)
	case strings.Contains(input, "2"):
		expressionArab(input)
	case strings.Contains(input, "3"):
		expressionArab(input)
	case strings.Contains(input, "4"):
		expressionArab(input)
	case strings.Contains(input, "5"):
		expressionArab(input)
	case strings.Contains(input, "6"):
		expressionArab(input)
	case strings.Contains(input, "7"):
		expressionArab(input)
	case strings.Contains(input, "8"):
		expressionArab(input)
	case strings.Contains(input, "9"):
		expressionArab(input)
	case strings.Contains(input, "10"):
		expressionArab(input)
	case strings.Contains(input, "0"):
		expressionArab(input)
	default:
		fmt.Println(expressionRome(input))

	}

}

func expressionArab(input string) (int, error) {
	if strings.Contains(input, "+") {
		numbers := strings.Split(input, "+")
		if len(numbers) != 2 {
			panic("Выражение должно содержать только два числа!")
		}
		numbOne, errOne := strconv.Atoi(strings.TrimSpace(numbers[0]))
		numbTwo, errTwo := strconv.Atoi(strings.TrimSpace(numbers[1]))
		if errOne != nil || errTwo != nil {
			panic("Что-то пошло не так...")
		}
		if numbOne <= 0 || numbOne > 10 {
			panic("Значениями могут быть только цифры от 1 до 10!")
		}
		if numbTwo <= 0 || numbTwo > 10 {
			panic("Значениями могут быть только цифры от 1 до 10!")
		}
		return fmt.Println(numbOne + numbTwo)
	} else if strings.Contains(input, "-") {
		numbers := strings.Split(input, "-")
		if len(numbers) != 2 {
			panic("Выражение должно содержать только два числа!")
		}
		numbOne, errOne := strconv.Atoi(strings.TrimSpace(numbers[0]))
		numbTwo, errTwo := strconv.Atoi(strings.TrimSpace(numbers[1]))
		if errOne != nil || errTwo != nil {
			panic("Что-то пошло не так...")
		}
		if numbOne <= 0 || numbOne > 10 {
			panic("Значениями могут быть только цифры от 1 до 10!")
		}
		if numbTwo <= 0 || numbTwo > 10 {
			panic("Значениями могут быть только цифры от 1 до 10!")
		}
		return fmt.Println(numbOne - numbTwo)

	} else if strings.Contains(input, "*") {
		numbers := strings.Split(input, "*")
		if len(numbers) != 2 {
			panic("Выражение должно содержать только два числа!")
		}
		numbOne, errOne := strconv.Atoi(strings.TrimSpace(numbers[0]))
		numbTwo, errTwo := strconv.Atoi(strings.TrimSpace(numbers[1]))
		if errOne != nil || errTwo != nil {
			panic("Что-то пошло не так...")
		}
		if numbOne <= 0 || numbOne > 10 {
			panic("Значениями могут быть только цифры от 1 до 10!")
		}
		if numbTwo <= 0 || numbTwo > 10 {
			panic("Значениями могут быть только цифры от 1 до 10!")
		}
		return fmt.Println(numbOne * numbTwo)
	} else if strings.Contains(input, "/") {
		numbers := strings.Split(input, "/")
		if len(numbers) != 2 {
			panic("Выражение должно содержать только два числа!")
		}
		numbOne, errOne := strconv.Atoi(strings.TrimSpace(numbers[0]))
		numbTwo, errTwo := strconv.Atoi(strings.TrimSpace(numbers[1]))
		if errOne != nil || errTwo != nil {
			panic("Что-то пошло не так...")
		}
		if numbOne <= 0 || numbOne > 10 {
			panic("Значениями могут быть только цифры от 1 до 10!")
		}
		if numbTwo <= 0 || numbTwo > 10 {
			panic("Значениями могут быть только цифры от 1 до 10!")
		}
		return fmt.Println(numbOne / numbTwo)
	} else {
		panic("Неверное выражение")
	}

}

func expressionRome(input string) (string, error) {
	romeToArab := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	arabToRome := []string{
		"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	}

	if strings.Contains(input, "+") {
		numbers := strings.Split(input, "+")
		if len(numbers) != 2 {
			panic("Выражение должно содержать только два числа!")
		}
		numb1 := strings.TrimSpace(numbers[0])
		numb2 := strings.TrimSpace(numbers[1])
		// конвертируем римские цифры в арабские
		digitOne, ok1 := romeToArab[numb1]
		digitTwo, ok2 := romeToArab[numb2]
		if !ok1 || !ok2 {
			panic("Неверные римские числа")
		}
		result := digitOne + digitTwo
		if result < 1 || result > 10 {
			panic("Результат не может быть меньше нуля или больше десяти!")
		}
		return (arabToRome[result]), nil // обратно в римские
	} else if strings.Contains(input, "-") {
		numbers := strings.Split(input, "-")
		if len(numbers) != 2 {
			panic("Выражение должно содержать только два числа!")
		}
		numb1 := strings.TrimSpace(numbers[0])
		numb2 := strings.TrimSpace(numbers[1])

		digitOne, ok1 := romeToArab[numb1]
		digitTwo, ok2 := romeToArab[numb2]
		if !ok1 || !ok2 {
			panic("Неверные римские числа")
		}
		result := digitOne - digitTwo
		if result < 1 || result > 10 {
			panic("Результат не может быть меньше нуля или больше десяти!")
		}
		return (arabToRome[result]), nil
	} else if strings.Contains(input, "*") {
		numbers := strings.Split(input, "*")
		if len(numbers) != 2 {
			panic("Выражение должно содержать только два числа!")
		}
		numb1 := strings.TrimSpace(numbers[0])
		numb2 := strings.TrimSpace(numbers[1])

		digitOne, ok1 := romeToArab[numb1]
		digitTwo, ok2 := romeToArab[numb2]
		if !ok1 || !ok2 {
			panic("Неверные римские числа")
		}
		result := digitOne * digitTwo
		if result < 1 || result > 10 {
			panic("Результат не может быть меньше нуля или больше десяти!")
		}
		return (arabToRome[result]), nil
	} else if strings.Contains(input, "/") {
		numbers := strings.Split(input, "/")
		if len(numbers) != 2 {
			panic("Выражение должно содержать только два числа!")
		}
		numb1 := strings.TrimSpace(numbers[0])
		numb2 := strings.TrimSpace(numbers[1])

		digitOne, ok1 := romeToArab[numb1]
		digitTwo, ok2 := romeToArab[numb2]
		if !ok1 || !ok2 {
			panic("Неверные римские числа")
		}
		result := digitOne / digitTwo
		if result < 1 || result > 10 {
			panic("Результат не может быть меньше нуля или больше десяти!")
		}
		return (arabToRome[result]), nil
	} else {
		panic("Неверное выражение")
	}

}
