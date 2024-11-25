package text

import (
	"fmt"
	"strconv"
	"strings"
)

func (t *Text) Subtraction_conditions() {

	//условие на определение знака между строками
	//для вычитания
	if strings.Contains(t.Content, "\" - \"") == true {
		if strings.Count(t.Content, "\"") == 4 && strings.Count(t.Content, "\" - \"") == 1 {

			//делим строку на две части и записываем их в разные переменные
			razdel_0 := strings.Split(t.Content, " - ")
			perv_0 := razdel_0[0]
			vtor_0 := razdel_0[1]
			delenie := strings.Split(perv_0, "\"")
			delenie_2 := strings.Split(vtor_0, "\"")
			X := delenie[1]
			XZ := delenie_2[1]

			//длина всей строки до разделения
			perv_0 = perv_0[1 : len(perv_0)-1]
			vtor_0 = vtor_0[1 : len(vtor_0)-1]
			Runes_Y := []rune(perv_0)
			Runes_YZ := []rune(vtor_0)
			dlina_Y := len(Runes_Y)
			dlina_YZ := len(Runes_YZ)

			//длина строки в кавычках
			Runes_X := []rune(X)
			Runes_XZ := []rune(XZ)
			dlina_X := len(Runes_X)
			dlina_XZ := len(Runes_XZ)

			//сравнение длины строки до операциии с длиной строки в кавычках
			if dlina_Y != dlina_X {
				panic("Первым аргументом выражения, подаваемым на вход, должна быть строка выделенная двойными кавычками.\n Знак операции должен быть отделен пробелом слева и справа.")
			} else if dlina_YZ != dlina_XZ {
				panic("Вторым аргументом выражения, подаваемым на вход, должна быть строка выделенная двойными кавычками.\n Знак операции должен быть отделен пробелом слева и справа.")
			}

		} else {
			panic("В выражении должно присутствовать два аргумента-строки, которые можно сложить или вычесть. Нельзя сложить или вычесть число со строкой")
		}

		//переводим str в int
		razdel_0 := strings.Split(t.Content, " - ")
		perv_0X := razdel_0[0]
		vtor_0Y := razdel_0[1]

		//условие на строку длиной не более 10 символов
		perv_0X = perv_0X[1 : len(perv_0X)-1] //скопировал выше
		vtor_0Y = vtor_0Y[1 : len(vtor_0Y)-1]
		strRunes_0X := []rune(perv_0X)
		strRunes_0Y := []rune(vtor_0Y)
		dlina_0X := len(strRunes_0X)
		dlina_0Y := len(strRunes_0Y)
		if dlina_0X > 10 {
			panic("Калькулятор должен принимать на вход строки длиной не более 10 символов! Первый аргумент-строка превышает этот лимит.")
		} else if dlina_0Y > 10 {
			panic("Калькулятор должен принимать на вход строки длиной не более 10 символов! Второй аргумент-строка превышает этот лимит.")
		}

		t.subtraction()

	} else {
		panic("Выражение должно иметь одну операцию: сложение, вычитание, деление, умножение.\n Знак операции должен быть отделен пробелом слева и справа.\n Складываются/вычитаются между собой только строки выделенные двойными кавычками.\n Выражение должно иметь два аргумента, не более.")
	}
}

func (t *Text) subtraction() {

	//разделяем строку на уменьшаемое и вычитаемое, получаем срез
	razdelenie := strings.Split(t.Content, " - ")

	//помещаем уменьшаемое и вычитаемое в переменные
	umensh := razdelenie[0]
	vichet := razdelenie[1]

	//удаляем кавычки у уменьшаемой части для удобной работы
	umensh = umensh[1 : len(umensh)-1]

	//разделяем уменьшаемую строку на слова
	words := strings.Split(umensh, " ")

	//за счет цикла разделаем каждое слово двойными кавычками
	//соединяем все в новую переменную newline
	var newline string
	for _, word := range words {
		word = strconv.Quote(word)
		newline = newline + word
	}

	t.Content = strings.ReplaceAll(newline, vichet, "\"")
	if strings.Count(t.Content, "\"\"\"") == 1 {
		fmt.Println(t.Content)
		t.Content = strings.ReplaceAll(t.Content, "\"\"\"", " ")
		fmt.Println(t.Content)
	} else {

		//заменяем двойные кавычки между словами на пробелы
		t.Content = strings.ReplaceAll(t.Content, "\"\"", " ")
		fmt.Println(t.Content)
		t.Content = strings.ReplaceAll(t.Content, "\"", "")
		fmt.Println(t.Content)
		fmt.Printf("%q \n", t.Content)
	}
}
