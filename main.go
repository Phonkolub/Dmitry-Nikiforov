package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/Phonkolub/Dmitry-Nikiforov/text"
)

func main() {
	text := &text.Text{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text.Content = scanner.Text()

		if strings.Contains(text.Content, "\" * ") == true && strings.Count(text.Content, "\" * ") == 1 {
			text.Multiplication_conditions()
		} else if strings.Contains(text.Content, "\" / ") == true && strings.Count(text.Content, "\" / ") == 1 {
			text.Division_conditions()
			// } else if strings.Contains(text.Content, "\" + \"") == true && strings.Count(text.Content, "\" + \"") == 1 {
			// 	text.Addition_conditions()
			// } else if strings.Contains(text.Content, "\" - \"") == true && strings.Count(text.Content, "\" - \"") == 1 {
			// 	text.Subtraction_conditions()
		} else {
			panic("Выражение должно иметь одну операцию: сложение, вычитание, деление, умножение.\n Знак операции должен быть отделен пробелом слева и справа.\n При делении и умножении на число не нужно записывать число в двойных кавычках.\n Умножается строка на число, а не наоборот.\n Складываются/вычитаются между собой только строки выделенные двойными кавычками.\n Выражение должно иметь два аргумента, не более.")
		}
	}

}
