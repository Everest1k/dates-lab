// Дополнительные тесты (Часть 2, по итогам мутационного тестирования) модуля datediff.
// Автор: Участник Б (тестировщик модуля А).
package datediff_test

import (
	"errors"
	"testing"

	"github.com/team/dates-lab/datediff"
)

// Убивает MA2, MA3 и мутанта Gremlins `d.Year < 1` → `d.Year <= 1`:
// граничные год, месяц и день.
func TestIsValidEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		d    D
		want bool
	}{
		{"первый день 1 года", D{1, 1, 1}, true},
		{"год 0", D{0, 1, 1}, false},
		{"отрицательный год", D{-5, 1, 1}, false},
		{"день 0", D{2024, 1, 0}, false},
		{"месяц 0", D{2024, 0, 1}, false},
		{"31 декабря", D{2024, 12, 31}, true},
		{"32 января", D{2024, 1, 32}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := datediff.IsValid(tc.d); got != tc.want {
				t.Errorf("IsValid(%v) = %v, ожидалось %v", tc.d, got, tc.want)
			}
		})
	}
}

// Правило високосного года на других вековых годах.
func TestIsLeapYearCenturies(t *testing.T) {
	tests := []struct {
		year int
		want bool
	}{
		{1600, true},
		{2100, false},
		{2400, true},
		{4, true},
		{1, false},
	}
	for _, tc := range tests {
		if got := datediff.IsLeapYear(tc.year); got != tc.want {
			t.Errorf("IsLeapYear(%d) = %v, ожидалось %v", tc.year, got, tc.want)
		}
	}
}

// Месяц 0 и декабрь.
func TestDaysInMonthEdgeCases(t *testing.T) {
	if _, err := datediff.DaysInMonth(2024, 0); !errors.Is(err, datediff.ErrInvalidMonth) {
		t.Errorf("DaysInMonth(2024, 0): ожидалась ErrInvalidMonth, получено %v", err)
	}
	if got, err := datediff.DaysInMonth(2024, 12); err != nil || got != 31 {
		t.Errorf("DaysInMonth(2024, 12) = %d, %v; ожидалось 31, nil", got, err)
	}
}

// Ветка ошибки DayOfYear.
func TestDayOfYearInvalid(t *testing.T) {
	if _, err := datediff.DayOfYear(D{2024, 2, 30}); !errors.Is(err, datediff.ErrInvalidDate) {
		t.Errorf("DayOfYear(2024-02-30): ожидалась ErrInvalidDate, получено %v", err)
	}
}

// Убивает MA4 и мутантов Gremlins в ordinal: длинные интервалы через века
// и через год, кратный 400.
func TestDaysBetweenLongSpans(t *testing.T) {
	tests := []struct {
		name string
		a, b D
		want int
	}{
		{"век без 29 февраля 1900", D{1900, 1, 1}, D{2000, 1, 1}, 36524},
		{"високосный 2000 год", D{2000, 1, 1}, D{2001, 1, 1}, 366},
		{"век с 29 февраля 2000", D{2000, 1, 1}, D{2100, 1, 1}, 36525},
		{"год после 29 февраля", D{2024, 3, 1}, D{2025, 3, 1}, 365},
		{"первый год эры", D{1, 1, 1}, D{2, 1, 1}, 365},
		{"от полёта Гагарина", D{1961, 4, 12}, D{2024, 10, 9}, 23191},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := datediff.DaysBetween(tc.a, tc.b)
			if err != nil || got != tc.want {
				t.Errorf("DaysBetween(%v, %v) = %d, %v; ожидалось %d, nil", tc.a, tc.b, got, err, tc.want)
			}
		})
	}
}

// Убивает MA5: некорректной может быть и вторая дата.
func TestDaysBetweenInvalidSecond(t *testing.T) {
	if _, err := datediff.DaysBetween(D{2024, 1, 1}, D{2023, 2, 29}); !errors.Is(err, datediff.ErrInvalidDate) {
		t.Errorf("DaysBetween с 29 февраля 2023 во второй дате: ожидалась ErrInvalidDate, получено %v", err)
	}
}
