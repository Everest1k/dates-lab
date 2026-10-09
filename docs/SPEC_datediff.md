# Спецификация модуля `datediff`

**Разработчик:** Участник А · **Тестировщик:** Участник Б
**Язык:** Go 1.22+ · **Пакет:** `github.com/team/dates-lab/datediff`
**Назначение:** проверка дат и вычисление разницы между датами.

Общие правила:
- календарь григорианский, в том числе для дат до 1582 года (пролептический);
- дата — структура `datediff.Date{Year, Month, Day}`; допустимы год >= 1, месяц 1..12, день от 1 до числа дней в месяце;
- високосный год делится на 4 и не делится на 100, либо делится на 400 (2000 — високосный, 1900 — нет).

| № | Функция | Что делает | Особые случаи |
|---|---------|-----------|---------------|
| 1 | `IsLeapYear(year int) bool` | Сообщает, високосный ли год | — |
| 2 | `DaysInMonth(year, month int) (int, error)` | Число дней в месяце | Месяц вне 1..12 → `0, ErrInvalidMonth` |
| 3 | `IsValid(d Date) bool` | Сообщает, существует ли дата | Год < 1, месяц вне 1..12, день вне диапазона → `false` |
| 4 | `DayOfYear(d Date) (int, error)` | Номер дня в году, 1 января — 1 | Несуществующая дата → `0, ErrInvalidDate` |
| 5 | `DaysBetween(a, b Date) (int, error)` | Число дней от `a` до `b` | `b` раньше `a` → отрицательное число; одинаковые даты → 0; любая из дат не существует → `0, ErrInvalidDate` |

Ошибки проверяются через `errors.Is(err, datediff.ErrInvalidDate)` и `errors.Is(err, datediff.ErrInvalidMonth)`.

Примеры:

```
IsLeapYear(2024)                            → true
IsLeapYear(1900)                            → false
DaysInMonth(2024, 2)                        → 29, nil
IsValid(Date{2023, 2, 29})                  → false
DayOfYear(Date{2023, 3, 1})                 → 60, nil
DaysBetween(Date{2024, 1, 1}, Date{2024, 12, 31}) → 365, nil
DaysBetween(Date{2024, 1, 1}, Date{2023, 12, 31}) → -1, nil
```
