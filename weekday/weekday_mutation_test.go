// Дополнительные тесты (Часть 2, по итогам мутационного тестирования) модуля weekday.
// Автор: Участник А (тестировщик модуля Б).
package weekday

import (
	"errors"
	"testing"
)

// Убивает MB1: проверяются все семь названий и номер за верхней границей.
func TestWeekdayNameAll(t *testing.T) {
	want := []string{"понедельник", "вторник", "среда", "четверг", "пятница", "суббота", "воскресенье"}
	for i, w := range want {
		if got, err := WeekdayName(i + 1); err != nil || got != w {
			t.Errorf("WeekdayName(%d) = %q, %v; ожидалось %q, nil", i+1, got, err, w)
		}
	}
	if _, err := WeekdayName(8); !errors.Is(err, ErrInvalidWeekday) {
		t.Errorf("WeekdayName(8): ожидалась ErrInvalidWeekday, получено %v", err)
	}
}

// Убивает MB2: воскресенье тоже выходной; ветка ошибки IsWeekend.
func TestIsWeekendSundayAndInvalid(t *testing.T) {
	if got, err := IsWeekend(D{2024, 10, 13}); err != nil || !got {
		t.Errorf("IsWeekend(2024-10-13, воскресенье) = %v, %v; ожидалось true, nil", got, err)
	}
	if _, err := IsWeekend(D{2024, 2, 30}); !errors.Is(err, ErrInvalidDate) {
		t.Errorf("IsWeekend(2024-02-30): ожидалась ErrInvalidDate, получено %v", err)
	}
}

// Убивает мутантов Gremlins `m < 3` → `m <= 3` и `month > 12` → `month >= 12`:
// март, декабрь, первый год эры, вековые февральские даты.
func TestDayOfWeekMoreDates(t *testing.T) {
	tests := []struct {
		name string
		d    D
		want int
	}{
		{"март", D{2024, 3, 8}, Friday},
		{"31 декабря", D{2024, 12, 31}, Tuesday},
		{"1 января 1 года", D{1, 1, 1}, Monday},
		{"1 января, воскресенье", D{2023, 1, 1}, Sunday},
		{"29 февраля 2000", D{2000, 2, 29}, Tuesday},
		{"28 февраля 2100", D{2100, 2, 28}, Sunday},
		{"1 марта 1900", D{1900, 3, 1}, Thursday},
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

// Убивает MB5, MB6 и мутантов Gremlins в правиле високосного года:
// несуществующие даты должны давать ошибку.
func TestDayOfWeekInvalidDates(t *testing.T) {
	invalid := []D{
		{2024, 11, 31}, // в ноябре 30 дней
		{2024, 1, 0},   // день 0
		{0, 1, 1},      // год 0
		{2024, 13, 1},  // месяц 13
		{2024, 0, 1},   // месяц 0
		{1900, 2, 29},  // 1900 — не високосный
		{2100, 2, 29},  // 2100 — не високосный
	}
	for _, d := range invalid {
		if _, err := DayOfWeek(d); !errors.Is(err, ErrInvalidDate) {
			t.Errorf("DayOfWeek(%v): ожидалась ErrInvalidDate, получено %v", d, err)
		}
	}
}

// Убивает мутанта `wd > Sunday` → `wd >= Sunday`; граничные номера и месяцы.
func TestCountWeekdayInMonthEdgeCases(t *testing.T) {
	tests := []struct {
		name                  string
		year, month, wd, want int
	}{
		{"воскресенья в сентябре 2024", 2024, 9, Sunday, 5},
		{"четверги в феврале 2024", 2024, 2, Thursday, 5},
		{"понедельники в феврале 2021", 2021, 2, Monday, 4},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CountWeekdayInMonth(tc.year, tc.month, tc.wd)
			if err != nil || got != tc.want {
				t.Errorf("CountWeekdayInMonth(%d, %d, %d) = %d, %v; ожидалось %d, nil", tc.year, tc.month, tc.wd, got, err, tc.want)
			}
		})
	}
	for _, wd := range []int{0, 8} {
		if _, err := CountWeekdayInMonth(2024, 9, wd); !errors.Is(err, ErrInvalidWeekday) {
			t.Errorf("CountWeekdayInMonth(2024, 9, %d): ожидалась ErrInvalidWeekday, получено %v", wd, err)
		}
	}
	if _, err := CountWeekdayInMonth(2024, 13, Monday); !errors.Is(err, ErrInvalidDate) {
		t.Errorf("CountWeekdayInMonth(2024, 13, 1): ожидалась ErrInvalidDate, получено %v", err)
	}
}

// Рабочие дни в феврале, ноябре, декабре и в вековые годы.
func TestWorkdaysInMonthEdgeCases(t *testing.T) {
	tests := []struct {
		year, month, want int
	}{
		{2024, 2, 21},
		{2021, 2, 20},
		{2024, 11, 21},
		{2024, 12, 22},
		{1900, 2, 20},
		{2000, 2, 21},
		{1, 1, 23},
	}
	for _, tc := range tests {
		if got, err := WorkdaysInMonth(tc.year, tc.month); err != nil || got != tc.want {
			t.Errorf("WorkdaysInMonth(%d, %d) = %d, %v; ожидалось %d, nil", tc.year, tc.month, got, err, tc.want)
		}
	}
	for _, bad := range [][2]int{{2024, 0}, {0, 5}} {
		if _, err := WorkdaysInMonth(bad[0], bad[1]); !errors.Is(err, ErrInvalidDate) {
			t.Errorf("WorkdaysInMonth(%d, %d): ожидалась ErrInvalidDate, получено %v", bad[0], bad[1], err)
		}
	}
}
