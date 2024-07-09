package main

import (
	"fmt"
	"testing"
)

// IntMin возвращает меньшее из двух целых чисел
func IntMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// TestIntMinBasic выполняет базовый тест функции IntMin
func TestIntMinBasic(t *testing.T) {

	// Тестируем функцию IntMin с входными значениями 2 и -2
	ans := IntMin(2, -2)

	// Проверяем, что результат равен -2
	if ans != -2 {

		// Если результат неверный, сообщаем об ошибке
		t.Errorf("IntMIn(2, -2) = %d; want -2", ans)
	}
}

// TestIntMinTableDriven выполняет тестирование функции IntMin с использованием подхода "table-driven"
func TestIntMinTableDriven(t *testing.T) {

	// Определяем таблицу тестов с различными входными значениями и ожидаемыми результатами
	var tests = []struct {
		a, b int
		want int
	}{
		{1, 0, 0},
		{0, 1, 0},
		{-1, 1, -1},
		{1, -1, -1},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// Проходим по каждому тестовому случаю из таблицы
	for _, tt := range tests {
		// Определяем имя теста для каждого случая
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)

		// Запускаем под-тест для каждого случая
		t.Run(testname, func(t *testing.T) {
			// Вызываем функцию IntMin с текущими входными значениями
			ans := IntMin(tt.a, tt.b)
			// Проверяем, что результат совпадает с ожидаемым
			if ans != tt.want {
				// Если результат неверный, сообщаем об ошибке
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// BenchmarkIntMin выполняет бенчмарк тест для функции IntMin
func BenchmarkIntMin(b *testing.B) {
	// Запускаем тест в цикле b.N раз для измерения производительности
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
