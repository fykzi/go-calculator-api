package calculator

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

var (
    InvalidExpressionErr = errors.New("invalid expression")
    DivisionByZeroErr = errors.New("division by zero")
)

func Calc(expression string) (float64, error) {
    expression = strings.Trim(expression, " ")
    if  expression == "" {
        return 0, InvalidExpressionErr
    }

    expressionSlice, err := ParseStringExpression(expression)
    if err != nil {
        return 0, err
    }

    rpn, err := parseExpToRpn(expressionSlice)
    if err != nil {
        return 0, err
    }
    return CalculateRpnExpression(rpn)
}

func CalculateRpnExpression(tokens []string) (float64, error) {
    operators := []string{"+", "-", "/", "*", "^"}


    for i := 0; i < len(tokens); i++ {
        if slices.Contains(operators, tokens[i]) {
            if i < 2 {
                return 0, InvalidExpressionErr
            }
            num1, err := strconv.ParseFloat(tokens[i - 2], 64)
            if err != nil {
                return 0, InvalidExpressionErr
            }

            num2, err := strconv.ParseFloat(tokens[i - 1], 64)
            if err != nil {
                return 0, InvalidExpressionErr
            }

            operator := tokens[i]
            
            var actionResult float64

            switch operator {
            case "+":
                actionResult = num1 + num2
            case "-":
                actionResult = num1 - num2
            case "*":
                actionResult = num1 * num2
            case "/":
                if num2 == 0 {
                    return 0, DivisionByZeroErr
                }
                actionResult = num1 / num2
            case "^":
                actionResult = math.Pow(num1, num2)
            }

            tokensEnd := tokens[i + 1:]
            tokens = append(tokens[:i - 2], fmt.Sprint(actionResult))
            tokens = append(tokens, tokensEnd...)
            i -= 2
        }
    } 

    result, _ := strconv.ParseFloat(tokens[0], 64)
    return result, nil
}

func ParseStringExpression(expression string) ([]string, error) {
    var expressionWithSpaces string
   
    for _, chr := range []rune(expression) {
        if slices.Contains([]rune{'-', '+', '*', '/', '^', '(', ')'}, chr) {
            expressionWithSpaces += " " + string(chr) + " "
        } else {
            expressionWithSpaces += string(chr)
        }
    }

    return strings.Fields(expressionWithSpaces), nil
}

func parseExpToRpn(expression []string) ([]string, error) {
    priorities := map[string]int{
        "+": 1,
        "-": 1,
        "*": 2,
        "/": 2,
        "^": 3,
    }

    var stack, resOutput []string

    for _, i := range expression {
        if i == "(" {
            stack = append(stack, i)
            continue
        }

        if curLevel, ok := priorities[i]; ok {
            n := len(stack)

            if n != 0 {
                if lastInStackLevel, _ := priorities[stack[n - 1]]; lastInStackLevel >= curLevel {
                    resOutput = append(resOutput, stack[n - 1])
                    stack = stack[:n - 1]
                }
            }

            stack = append(stack, i)
            continue
        }

        if i == ")" {
            for a := len(stack) - 1; ; a-- {
                if len(stack) == 0 {
                    return make([]string, 0), InvalidExpressionErr
                }
                
                last := stack[a]
                stack = stack[:a]

                if last == "(" {
                    break
                }

                resOutput = append(resOutput, last)
            }
            continue
        }

       resOutput = append(resOutput, i) 
    }
    
    for i := len(stack) - 1; i >= 0; i-- {
        chr := stack[i]
        if chr == "(" {
            return make([]string, 0), InvalidExpressionErr
        }
        resOutput = append(resOutput, chr)
    }
    return resOutput, nil
}
