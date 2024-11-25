package text

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func (t *Text) Multiplication_conditions() {

	//условие на число: не больше 10 при умножении
	//делим строку на две части и записываем их в разные переменные

	razdel_2 := strings.Split(t.Content, "*")
	vtor_2 := razdel_2[1]

	//убираем из второй строки все кроме чисел для корректного перевода в int
	reg_2 := regexp.MustCompile("[^0-9]+")
	vtor_2 = reg_2.ReplaceAllString(vtor_2, "")

	//переводим str в int
	number_2, _ := strconv.Atoi(vtor_2)
	var P int
	if number_2 > 10 {
		panic("oКалькулятор должен принимать на вход целые числа от 1 до 10 включительно, не более!")
	} else if reflect.TypeOf(number_2) != reflect.TypeOf(P) {
		panic("oКалькулятор должен принимать на вход целые числа от 1 до 10 включительно, не более!")
	}

	//условие на определение знака между строками
	//для вычитания
	if strings.Contains(t.Content, "\" * ") == true {
		if strings.Count(t.Content, "\"") == 2 && strings.Count(t.Content, "\" * ") == 1 {

			//делим строку на две части и записываем их в разные переменные
			razdel_0 := strings.Split(t.Content, " * ")
			perv_0 := razdel_0[0]
			delenie := strings.Split(perv_0, "\"")
			X := delenie[1]

			//длина всей строки до разделения
			perv_0 = perv_0[1 : len(perv_0)-1]
			Runes_Y := []rune(perv_0)
			dlina_Y := len(Runes_Y)

			//длина строки в кавычках
			Runes_X := []rune(X)
			dlina_X := len(Runes_X)

			//сравнение длины строки до операциии с длиной строки в кавычках
			if dlina_Y != dlina_X {
				panic("Множитель введет неправильно! Первым аргументом выражения, подаваемым на вход, должна быть строка выделенная двойными кавычками.")
			}

		} else {
			panic("Число записано неправильно!\n Нельзя умножать строку на строку.\n Знак операции должен быть отделен пробелом слева и справа.\n При делении и умножении на число не нужно записывать число в двойных кавычках.\n Число не должно быть больше 10")
		}

		//переводим str в int
		razdel_0 := strings.Split(t.Content, " * ")
		perv_0 := razdel_0[0]
		vtor_0 := razdel_0[1]
		number_0, _ := strconv.Atoi(vtor_0)
		if number_0 > 10 {
			panic("Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более!")
		}

		//условие на строку длиной не более 10 символов в левой части в случае умножения
		perv_0 = perv_0[1 : len(perv_0)-1]
		strRunes_0 := []rune(perv_0)
		dlina_00 := len(strRunes_0)
		if dlina_00 > 10 {
			panic("Калькулятор должен принимать на вход строки длиной не более 10 символов!")
		}

		//возвращаемся к определению знака
		strRunes_0 = []rune(vtor_0)
		dlina_01 := len(strRunes_0)
		if dlina_01 > 3 {
			panic("Число записано неправильно!\n Знак операции должен быть отделен пробелом слева и справа.\n При делении и умножении на число не нужно записывать число в двойных кавычках.\n Число не должно быть больше 10")
		} else if strings.Contains(t.Content, "\" * ") == false {
			panic("Число записано неправильно!\n Знак операции должен быть отделен пробелом слева и справа.\n При делении и умножении на число не нужно записывать число в двойных кавычках.\n Число не должно быть больше 10")
		} else {
			vtor_0 = strings.ReplaceAll(vtor_0, " ", "")
			strRunes_0 = []rune(vtor_0)
			dlina_01 = len(strRunes_0)
			if number_0 < 10 && dlina_01 > 1 == true {
				panic("Знак операции должен быть отделен пробелом слева и справа.\n При делении и умножении на число не нужно записывать число в двойных кавычках.\n Целое число не должно быть больше 10")
			}
		}
		t.multiplication()
	} else {
		panic("Выражение должно иметь одну операцию: сложение, вычитание, деление, умножение.\n Знак операции должен быть отделен пробелом слева и справа.\n При делении и умножении на число не нужно записывать число в двойных кавычках.\n Умножается строка на число, а не наоборот.\n Строка должна быть обозначена двойными кавычками")
	}
}

func (t *Text) multiplication() {

	razdel_0 := strings.Split(t.Content, " * ")
	perv_0 := razdel_0[0]
	vtor_0 := razdel_0[1]

	//выделяем число в строке и записываем его
	reg := regexp.MustCompile("[^0-9]+")
	str := reg.ReplaceAllString(vtor_0, "")

	//конвертируем строку в целочисленный тип для использования Repeat
	nestr, _ := strconv.Atoi(str)

	strRunes_0 := []rune(perv_0)
	perv_0 = perv_0[1 : len(strRunes_0)-1]

	t.Content = strings.Repeat(perv_0, nestr)

	strRunes_0 = []rune(t.Content)
	dlina_01 := len(strRunes_0)

	if dlina_01 > 40 {
		t.Content = t.Content[0:40]

		fmt.Print("\"", t.Content)
		fmt.Print("...\"")
	} else {
		fmt.Printf("%q \n", t.Content)
	}
}
