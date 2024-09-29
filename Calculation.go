package main

import (
	"fmt"
	"strings"
	"log"
	"github.com/expr-lang/expr"

)

func romanToInt(s string) int {
    romanMap := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    total := 0
    prevValue := 0

    for _, char := range s {
        value := romanMap[char]
        if value > prevValue {
            total += value - 2*prevValue 
            total += value
        }
        prevValue = value
    }

    return total
}

func containsRoman(s string) bool {
    for _, char := range s {
        if _, exists := map[rune]struct{}{
            'I': {}, 'V': {}, 'X': {}, 'L': {}, 'C': {}, 'D': {}, 'M': {},
        }[char]; exists {
            return true
        }
    }
    return false
}

func main() {
	var input string

	fmt.Print("Введите выражение (например, 5+3): ")
    _, err := fmt.Scan(&input)
   	if err != nil {
       log.Fatal("Ошибка ввода:", err)
    }
   	delimiters := "+-*/"
    var data []string

	parts := strings.FieldsFunc(input, func(r rune) bool {
        return strings.ContainsRune(delimiters, r)
    })

	for i, part := range parts {
        trimmedPart := strings.TrimSpace(part)
        if containsRoman(trimmedPart) {
            arabicValue := romanToInt(trimmedPart)
            data = append(data, fmt.Sprint(arabicValue))
        } else {
            data = append(data, trimmedPart)
        }

        if i < len(parts)-1 {
            data = append(data, string(input[len(trimmedPart)+i]))
        }
    }

	result, _ := expr.Eval(strings.Join(data, ""), nil)

	fmt.Println("Результат:", result)
}