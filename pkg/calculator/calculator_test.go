package calculator_test

import (
	"fmt"
	"testing"

	"github.com/fykzi/go-calculator-api/pkg/calculator"
)

// Тест на сложные выражения с приоритетами
func TestCalc_ComplexExpressions(t *testing.T) {
	tests := []struct {
		expr     string
		expected float64
	}{
		{"3 + 5 * 2", 13},                // Сложение и умножение
		{"(3 + 5) * 2", 16},              // Скобки и умножение
		{"2 + 3 * (4 - 1)", 11},          // Скобки и приоритет умножения
		{"(2 + 3) * (4 - 1)", 15},        // Два выражения в скобках
		{"10 / 2 + 5", 10},               // Деление и сложение
		{"10 - 3 * 2", 4},                // Вычитание и умножение
		{"2 ^ 3 + 4", 12},                // Возведение в степень и сложение
		{"(3 + 2) ^ 2", 25},              // Скобки и возведение в степень
		{"3 + 2 * 2 ^ 2", 11},            // Смешанные операции (возведение в степень, умножение)
		{"(3 + 2) * 2 ^ 2", 20},          // Скобки, умножение и возведение в степень
		{"(2 + 3) ^ (2 + 1)", 125},       // Возведение в степень с результатом в скобках
		{"(2 + 3) ^ 2 / 5", 5},          // Возведение в степень с делением
		{"9 / 3 + 2 ^ 3", 11},            // Деление и возведение в степень
		{"(5 - 3) * (4 + 2) / 2", 6},    // Выражение с несколькими операциями
		{"(5 * 2 + 3) / 4", 3.25},         // Сложные операции с делением
		{"(10 + 2) * (5 - 1) / 3", 16},   // Выражение с делением и скобками
		{"5 + 3 * (2 + 3 * 2) - 4", 25},  // Сложные выражения с несколькими скобками
	}

	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			result, err := calculator.Calc(tt.expr)
			if err != nil {
				t.Errorf("Calc() error = %v", err)
			}
			if result != tt.expected {
				t.Errorf("Calc() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Тест на некорректные выражения с ошибками синтаксиса
func TestCalc_InvalidExpressions(t *testing.T) {
	tests := []struct {
		expr     string
		expected error
	}{
		{"10 / 0", fmt.Errorf("division by zero")},                   // Деление на ноль
		{"5 +", fmt.Errorf("invalid expression")},                    // Недопустимый оператор
		{"(3 + 5", fmt.Errorf("invalid expression")},        // Отсутствует закрывающая скобка
		{"5 * * 2", fmt.Errorf("invalid expression")},                  // Лишний оператор
		{"3 + 2 *", fmt.Errorf("invalid expression")},                // Отсутствует операнд
		{"2 ^", fmt.Errorf("invalid expression")},                    // Отсутствует операнд после возведения в степень
		{"(3 + 5))", fmt.Errorf("invalid expression")},        // Лишняя закрывающая скобка
		{"(5 + 3 * (2 + 1)", fmt.Errorf("invalid expression")}, // Отсутствует закрывающая скобка
		{"(5 + 3 ^)", fmt.Errorf("invalid expression")},              // Ошибка в выражении с возведением в степень
		{"(2 ^ 3", fmt.Errorf("invalid expression")},        // Отсутствует закрывающая скобка после возведения в степень
	}

	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			_, err := calculator.Calc(tt.expr)
			if err == nil || err.Error() != tt.expected.Error() {
				t.Errorf("Calc() error = %v, want %v", err, tt.expected)
			}
		})
	}
}

