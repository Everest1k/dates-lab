// Package datediff
// разницы между датами по григорианскому календарю.
package datediff

import "errors"

// Date — календарная дата: год >= 1, месяц 1..12, день 1..число дней в месяце.
type Date struct {
	Year, Month, Day int
}

var (
	// ErrInvalidMonth возвращается, если номер месяца вне диапазона 1..12.
	ErrInvalidMonth = errors.New("datediff: месяц вне диапазона 1..12")
	// ErrInvalidDate возвращается, если дата не существует.
	ErrInvalidDate = errors.New("datediff: некорректная дата")
)

var monthDays = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// IsLeapYear сообщает, високосный ли год: год делится на 4 и не делится
// на 100, либо делится на 400.
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0
}

// DaysInMonth возвращает число дней в месяце month года year.
// Для месяца вне диапазона 1..12 возвращает 0 и ErrInvalidMonth.
func DaysInMonth(year, month int) (int, error) {
	if month < 1 || month > 12 {
		return 0, ErrInvalidMonth
	}
	if month == 2 && IsLeapYear(year) {
		return 29, nil
	}
	return monthDays[month-1], nil
}

// IsValid сообщает, существует ли дата d.
func IsValid(d Date) bool {
	if d.Year < 1 {
		return false
	}
	n, err := DaysInMonth(d.Year, d.Month)
	if err != nil {
		return false
	}
	return d.Day >= 1 && d.Day <= n
}

// DayOfYear возвращает порядковый номер дня в году (1 января — 1).
// Для несуществующей даты возвращает 0 и ErrInvalidDate.
func DayOfYear(d Date) (int, error) {
	if !IsValid(d) {
		return 0, ErrInvalidDate
	}
	n := d.Day
	for m := 1; m < d.Month; m++ {
		k, _ := DaysInMonth(d.Year, m)
		n += k
	}
	return n, nil
}

// DaysBetween возвращает число дней от даты a до даты b: положительное,
// если b позже a, отрицательное, если раньше, и 0 для одинаковых дат.
// Если хотя бы одна дата не существует, возвращает 0 и ErrInvalidDate.
func DaysBetween(a, b Date) (int, error) {
	if !IsValid(a) || !IsValid(b) {
		return 0, ErrInvalidDate
	}
	return ordinal(b) - ordinal(a), nil
}

// ordinal возвращает номер дня, считая 1 января 1 года днём 1.
func ordinal(d Date) int {
	y := d.Year - 1
	doy, _ := DayOfYear(d)
	return y*365 + y/4 - y/100 + y/400 + doy
}
