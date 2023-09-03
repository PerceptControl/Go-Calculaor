package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"romannumeral"
	"strconv"
	"strings"
)

const PLUS, MINUS, DIVIDE, MULTIPLY = '+', '-', '/', '*'

func main() {
	defer func() { 
		if err:=recover(); err != nil {
			fmt.Println(err)
			main()
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for true {
		example_as_string, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		operator, operands, is_rome := parseLine(example_as_string)
		
		res := calculate(operator, operands[0], operands[1])
		output := "" 
		if is_rome {
			if res <= 0 {
				panic("В римской системе нет отрицательных чисел и 0")
			}
			output, _ = romannumeral.IntToString(res)
		} else {
			output = strconv.Itoa(res)
		}

		fmt.Println(output)
	}
}

func parseLine(str string) (operator rune, operands [2]int, is_rome bool) {
	str = spaceFieldsJoin(str)

	operator, err := parseOperator(str) 
	if err != nil {
		panic(err)
	}
	operands_as_string := strings.Split(str, string(operator))

	//если во второй части примера еще есть оператор то весь пример гарантированно некорректный
	_, err = parseOperator(operands_as_string[1])
	if err == nil {
		panic(errors.New("Уравнение должно состоять из двух операндов и одного оператора"))
	}
 	
	
	operands, is_rome = parseOperands(operands_as_string)
	return operator, operands, is_rome
}

func spaceFieldsJoin(str string) string {
	return strings.Join(strings.Fields(str), "")
}

func parseOperator(str string) (operator rune, err error) {
	for _, char := range str {
		switch char {
			case '+': 
				return PLUS, nil
			case '-':
				return MINUS, nil
			case '*': 
				return MULTIPLY, nil
			case '/':
				return DIVIDE, nil
		}
	}
	return -1, errors.New("Не указан оператор")
}

func parseOperands(operands_as_string []string) (operands [2]int, is_roman bool) {
	if len(operands_as_string) > 2 {
		panic(errors.New("Уравнение должно состоять из двух операндов и одного оператора"))
	}
	operand1, is_rome_1 := parseNumber(operands_as_string[0])
	operand2, is_rome_2 := parseNumber(operands_as_string[1])

	if (!is_rome_1 && is_rome_2) || (is_rome_1 && !is_rome_2) {
		panic("Оба операнда должы быть либо арабскими либо римскими числами")
	}

	return [2]int{ operand1, operand2 }, is_rome_1
}

func parseNumber(str string) (result int, is_roman bool){
	res, err := strconv.Atoi(str)
	if err == nil {
		if res < 1 || res > 10 {
			panic("Число быть должно от 1 до 10 включительно")
		}
		return res, false
	}


	res, err = romannumeral.StringToInt(str)
	if err == nil {
		if res < 1 || res > 10 {
			panic("Число быть должно от 1 до 10 включительно")
		}
		return res, true
	}

	panic("Операнд должен быть корректным римским или арабским числом")
}

func calculate(operator rune, operand1 int, operand2 int) int {
	switch operator {
	case PLUS:
		return operand1 + operand2
	case DIVIDE:
		return operand1 / operand2
	case MULTIPLY:
		return operand1 * operand2
	case MINUS:
		return operand1 - operand2
	}
	
	return 0
}