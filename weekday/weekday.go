// и подсчёты по дням недели (григорианский календарь).
package weekday

import "errors"

// Date — календарная дата: год >= 1, месяц 1..12, день 1..число дней в месяце.
type Date struct {
	Year, Month, Day int
}

// Номера дней недели по ISO 8601.
const (
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
	Sunday    = 7
)

var (
	// ErrInvalidDate возвращается для несуществующей даты или месяца.
	ErrInvalidDate = errors.New("weekday: некорректная дата")
	// ErrInvalidWeekday возвращается, если номер дня недели вне диапазона 1..7.
	ErrInvalidWeekday = errors.New("weekday: номер дня недели вне диапазона 1..7")
)

var names = [7]string{
	"понедельник", "вторник", "среда", "четверг",
	"пятница", "суббота", "воскресенье",
}

// DayOfWeek возвращает день недели даты d: 1 — понедельник, …, 7 — воскресенье.
// Вычисляется по формуле Целлера. Для несуществующей даты — 0 и ErrInvalidDate.
func DayOfWeek(d Date) (int, error) {
	if !isValid(d) {
		return 0, ErrInvalidDate
	}
	y, m := d.Year, d.Month
	if m < 2 {
		m += 12
		y--
	}
	k, j := y%100, y/100
	h := (d.Day + 13*(m+1)/5 + k + k/4 + j/4 + 5*j) % 7 // 0 — суббота, 1 — воскресенье, …
	return (h+5)%7 + 1, nil
}

// WeekdayName возвращает название дня недели по номеру 1..7.
// Для номера вне диапазона — пустая строка и ErrInvalidWeekday.
func WeekdayName(n int) (string, error) {
	if n < Monday || n > Sunday {
		return "", ErrInvalidWeekday
	}
	return names[n-1], nil
}

// IsWeekend сообщает, приходится ли дата на субботу или воскресенье.
func IsWeekend(d Date) (bool, error) {
	wd, err := DayOfWeek(d)
	if err != nil {
		return false, err
	}
	return wd >= Saturday, nil
}

// CountWeekdayInMonth возвращает, сколько раз день недели wd встречается
// в месяце month года year (например, сколько понедельников в сентябре).
func CountWeekdayInMonth(year, month, wd int) (int, error) {
	if wd < Monday || wd > Sunday {
		return 0, ErrInvalidWeekday
	}
	n, err := daysInMonth(year, month)
	if err != nil {
		return 0, err
	}
	count := 0
	for day := 1; day <= n; day++ {
		w, _ := DayOfWeek(Date{year, month, day})
		if w == wd {
			count++
		}
	}
	return count, nil
}

// WorkdaysInMonth возвращает число рабочих дней (понедельник–пятница)
// в месяце month года year. Праздники не учитываются.
func WorkdaysInMonth(year, month int) (int, error) {
	n, err := daysInMonth(year, month)
	if err != nil {
		return 0, err
	}
	count := 0
	for day := 1; day <= n; day++ {
		w, _ := DayOfWeek(Date{year, month, day})
		if w <= Friday {
			count++
		}
	}
	return count, nil
}

// daysInMonth возвращает число дней в месяце или ErrInvalidDate.
func daysInMonth(year, month int) (int, error) {
	if year < 1 || month < 1 || month > 12 {
		return 0, ErrInvalidDate
	}
	switch month {
	case 2:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return 29, nil
		}
		return 28, nil
	case 4, 6, 9, 11:
		return 30, nil
	default:
		return 31, nil
	}
}

// isValid сообщает, существует ли дата d.
func isValid(d Date) bool {
	n, err := daysInMonth(d.Year, d.Month)
	return err == nil && d.Day >= 1 && d.Day <= n
}
