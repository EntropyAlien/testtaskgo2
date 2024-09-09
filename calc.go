package main

import (
	"fmt"
	//"io"
	"bufio"
	"os"
	"strings"
)

func bufio_scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода", err)
	}
	return in.Text()
}

func main() {
	fmt.Println("Введите выражение согласло форме (число операция число)")
	var input string
	input = bufio_scan() // использование самодельной функции пееревода введенного с консоли текста с пробелами в строк
	var splitted_input_array []string
	splitted_input_array = strings.Split(input, " ")
	// проверка правильности ввода
	if len(splitted_input_array) != 3 {
		fmt.Println("Ожидается запрос вида: число1 операция число2")
		panic("")
	}
	// получим из массива числа с помощью функции
	word0 := splitted_input_array[0]
	word1 := splitted_input_array[1]
	word2 := splitted_input_array[2]
	var num0, romanornot0 = translate(word0)
	var num2, romanornot2 = translate(word2)

	switch {
	case num0%1 != 0:
		fmt.Println("Калькулятор умеет работать только с целыми числами")
		panic("")
	case num0 < 1:
		fmt.Println("Калькулятор умеет работать только с целыми числами от 1 до 10 включительно")
		panic("")
	case num0 > 10:
		fmt.Println("Калькулятор умеет работать только с целыми числами от 1 до 10 включительно")
		panic("")
	case num2%1 != 0:
		fmt.Println("Калькулятор умеет работать только с целыми числами")
		panic("")
	case num2 < 1:
		fmt.Println("Калькулятор умеет работать только с целыми числами от 1 до 10 включительно")
		panic("")
	case num2 > 10:
		fmt.Println("Калькулятор умеет работать только с целыми числами от 1 до 10 включительно")
		panic("")
	}
	var function func(int, int) int
	switch word1 {
	case "+":
		function = add
	case "-":
		function = sub
	case "*":
		function = mul
	case "/":
		function = div
	default:
		fmt.Println("Неизвестный оператор", word1)
		panic("")
	}

	// ИТОГОВОЕ ВЫЧИСЛЕНИЕ
	if romanornot0 != romanornot2 {
		fmt.Println("Нельзя использовать цифры разных алфавитов")
		panic("")
	} else if romanornot0 == 0 {
		result := function(num0, num2)
		fmt.Println(result)
	} else {
		result := function(num0, num2)
		if result < 1 {
			fmt.Println("Результат вычисления меньше единицы")
			panic("")
		} else {
			fmt.Printf(romanizator(result))
		}
	}
	/*	var eternity string // чтобы исполняемый exe-файл не закрывался
		fmt.Scan(&eternity)
	*/
}

// функция перевода введенных чисел-строк в числа-числа (word0 здесь омоним word0 внутри функции main потому что разные области видимости)
func translate(word0 string) (int, int) {
	switch word0 { //корюзлое решение без использования пакета strconv
	case "1":
		num0 := 1
		return num0, 0
	case "2":
		num0 := 2
		return num0, 0
	case "3":
		num0 := 3
		return num0, 0
	case "4":
		num0 := 4
		return num0, 0
	case "5":
		num0 := 5
		return num0, 0
	case "6":
		num0 := 6
		return num0, 0
	case "7":
		num0 := 7
		return num0, 0
	case "8":
		num0 := 8
		return num0, 0
	case "9":
		num0 := 9
		return num0, 0
	case "10":
		num0 := 10
		return num0, 0
	case "I":
		num0 := 1
		return num0, 1
	case "II":
		num0 := 2
		return num0, 1
	case "III":
		num0 := 3
		return num0, 1
	case "IV":
		num0 := 4
		return num0, 1
	case "V":
		num0 := 5
		return num0, 1
	case "VI":
		num0 := 6
		return num0, 1
	case "VII":
		num0 := 7
		return num0, 1
	case "VIII":
		num0 := 8
		return num0, 1
	case "IX":
		num0 := 9
		return num0, 1
	case "X":
		num0 := 10
		return num0, 1
	default:
		fmt.Println("Введен неизвестный символ", word0)
		panic("")
	}
}

// арифметические функции
func add(num0, num2 int) int {
	return num0 + num2
}
func sub(num0, num2 int) int {
	return num0 - num2
}
func mul(num0, num2 int) int {
	return num0 * num2
}
func div(num0, num2 int) int {
	if num2 == 0 {
		fmt.Println("Нельзя делить на ноль")
		panic("")
	}
	floatnum0 := float32(num0)
	floatnum2 := float32(num2)
	div_result := floatnum0 / floatnum2
	output := int(div_result)
	return output
}

// функция превращения числа из арабских цифр в число из римских цифр
func romanizator(arabic int) string {
	romearray := []string{} // в этот массив без пробелов записываются строки-римские цифры
	for arabic > 0 {
		switch {
		case arabic == 100:
			arabic -= 100
			romearray = append(romearray, "C")
		case arabic >= 90:
			arabic -= 90
			romearray = append(romearray, "XC")
		case arabic >= 50:
			arabic -= 50
			romearray = append(romearray, "L")
		case arabic >= 40:
			arabic -= 40
			romearray = append(romearray, "XL")
		case arabic >= 10:
			arabic -= 10
			romearray = append(romearray, "X")
		case arabic >= 9:
			arabic -= 9
			romearray = append(romearray, "IX")
		case arabic >= 8:
			arabic -= 8
			romearray = append(romearray, "VIII")
		case arabic >= 7:
			arabic -= 7
			romearray = append(romearray, "VII")
		case arabic >= 6:
			arabic -= 6
			romearray = append(romearray, "VI")
		case arabic >= 5:
			arabic -= 5
			romearray = append(romearray, "V")
		case arabic >= 4:
			arabic -= 4
			romearray = append(romearray, "IV")
		case arabic >= 3:
			arabic -= 3
			romearray = append(romearray, "III")
		case arabic >= 2:
			arabic -= 2
			romearray = append(romearray, "II")
		case arabic >= 1:
			arabic -= 1
			romearray = append(romearray, "I")
		}
	}
	/*var reverse []string
	for i := len(romearray) - 1; i >= 0; i-- {
		reverse = append(reverse, romearray[i])
	}*/
	conjoined := strings.Join(romearray, "") // из массива "слипшихся" цифр генерируется строка
	return conjoined
}
