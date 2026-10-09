// Тесты модуля datediff. Автор: Участник Б (тестировщик модуля А).
// Тесты «чёрного ящика»: пакет datediff_test видит только экспортируемые имена.
package datediff_test

import (
	"errors"
	"testing"

	"github.com/team/dates-lab/datediff"
)

type D = datediff.Date

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name string
		year int
		want bool
	}{
		{"делится на 4", 2024, true},
		{"не делится на 4", 2023, false},
		{"делится на 100, не на 400", 1900, false},
		{"делится на 400", 2000, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := datediff.IsLeapYear(tc.year); got != tc.want {
				t.Errorf("IsLeapYear(%d) = %v, ожидалось %v", tc.year, got, tc.want)
			}
		})
	}
}

func TestDaysInMonth(t *testing.T) {
	tests := []struct {
		name        string
		year, month int
		want        int
	}{
		{"январь", 2023, 1, 31},
		{"апрель", 2023, 4, 30},
		{"февраль обычного года", 2023, 2, 28},
		{"февраль високосного года", 2024, 2, 29},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := datediff.DaysInMonth(tc.year, tc.month)
			if err != nil || got != tc.want {
				t.Errorf("DaysInMonth(%d, %d) = %d, %v; ожидалось %d, nil", tc.year, tc.month, got, err, tc.want)
			}
		})
	}
}

func TestDaysInMonthInvalid(t *testing.T) {
	if _, err := datediff.DaysInMonth(2024, 13); !errors.Is(err, datediff.ErrInvalidMonth) {
		t.Errorf("DaysInMonth(2024, 13): ожидалась ErrInvalidMonth, получено %v", err)
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		d    D
		want bool
	}{
		{"обычная дата", D{2024, 10, 9}, true},
		{"29 февраля високосного года", D{2024, 2, 29}, true},
		{"29 февраля обычного года", D{2023, 2, 29}, false},
		{"31 апреля", D{2024, 4, 31}, false},
		{"13-й месяц", D{2024, 13, 1}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := datediff.IsValid(tc.d); got != tc.want {
				t.Errorf("IsValid(%v) = %v, ожидалось %v", tc.d, got, tc.want)
			}
		})
	}
}

func TestDayOfYear(t *testing.T) {
	tests := []struct {
		name string
		d    D
		want int
	}{
		{"1 января", D{2023, 1, 1}, 1},
		{"1 марта обычного года", D{2023, 3, 1}, 60},
		{"31 декабря високосного года", D{2024, 12, 31}, 366},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := datediff.DayOfYear(tc.d)
			if err != nil || got != tc.want {
				t.Errorf("DayOfYear(%v) = %d, %v; ожидалось %d, nil", tc.d, got, err, tc.want)
			}
		})
	}
}

func TestDaysBetween(t *testing.T) {
	tests := []struct {
		name string
		a, b D
		want int
	}{
		{"одна и та же дата", D{2024, 5, 1}, D{2024, 5, 1}, 0},
		{"внутри года", D{2024, 1, 1}, D{2024, 12, 31}, 365},
		{"через Новый год", D{2023, 12, 31}, D{2024, 1, 1}, 1},
		{"в обратную сторону", D{2024, 1, 1}, D{2023, 12, 31}, -1},
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

func TestDaysBetweenInvalid(t *testing.T) {
	if _, err := datediff.DaysBetween(D{2023, 2, 30}, D{2024, 1, 1}); !errors.Is(err, datediff.ErrInvalidDate) {
		t.Errorf("DaysBetween с 30 февраля: ожидалась ErrInvalidDate, получено %v", err)
	}
}
