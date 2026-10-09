// Тесты модуля weekday. Автор: Участник А (тестировщик модуля Б).
// Тесты «чёрного ящика»: пакет weekday_test видит только экспортируемые имена.
package weekday

import (
	"errors"
	"testing"
)

type D = Date

func TestDayOfWeek(t *testing.T) {
	tests := []struct {
		name string
		d    D
		want int
	}{
		{"обычная дата (среда)", D{2024, 10, 9}, Wednesday},
		{"полёт Гагарина (среда)", D{1961, 4, 12}, Wednesday},
		{"1 января 2000 (суббота)", D{2000, 1, 1}, Saturday},
		{"29 февраля 2024 (четверг)", D{2024, 2, 29}, Thursday},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := DayOfWeek(tc.d)
			if err != nil || got != tc.want {
				t.Errorf("DayOfWeek(%v) = %d, %v; ожидалось %d, nil", tc.d, got, err, tc.want)
			}
		})
	}
}

func TestDayOfWeekInvalid(t *testing.T) {
	if _, err := DayOfWeek(D{2023, 2, 29}); !errors.Is(err, ErrInvalidDate) {
		t.Errorf("DayOfWeek(2023-02-29): ожидалась ErrInvalidDate, получено %v", err)
	}
}

func TestWeekdayName(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "понедельник"},
		{3, "среда"},
		{7, "воскресенье"},
	}
	for _, tc := range tests {
		got, err := WeekdayName(tc.n)
		if err != nil || got != tc.want {
			t.Errorf("WeekdayName(%d) = %q, %v; ожидалось %q, nil", tc.n, got, err, tc.want)
		}
	}
	if _, err := WeekdayName(0); !errors.Is(err, ErrInvalidWeekday) {
		t.Errorf("WeekdayName(0): ожидалась ErrInvalidWeekday, получено %v", err)
	}
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		name string
		d    D
		want bool
	}{
		{"суббота", D{2024, 10, 12}, true},
		{"среда", D{2024, 10, 9}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := IsWeekend(tc.d)
			if err != nil || got != tc.want {
				t.Errorf("IsWeekend(%v) = %v, %v; ожидалось %v, nil", tc.d, got, err, tc.want)
			}
		})
	}
}

func TestCountWeekdayInMonth(t *testing.T) {
	tests := []struct {
		name string
		wd   int
		want int
	}{
		{"понедельники в сентябре 2024", Monday, 5},
		{"вторники в сентябре 2024", Tuesday, 4},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CountWeekdayInMonth(2024, 9, tc.wd)
			if err != nil || got != tc.want {
				t.Errorf("CountWeekdayInMonth(2024, 9, %d) = %d, %v; ожидалось %d, nil", tc.wd, got, err, tc.want)
			}
		})
	}
}

func TestWorkdaysInMonth(t *testing.T) {
	got, err := WorkdaysInMonth(2024, 10)
	if err != nil || got != 23 {
		t.Errorf("WorkdaysInMonth(2024, 10) = %d, %v; ожидалось 23, nil", got, err)
	}
}
