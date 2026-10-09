# Спецификация модуля `weekday`

**Разработчик:** Участник Б · **Тестировщик:** Участник А
**Язык:** Go 1.22+ · **Пакет:** `github.com/team/dates-lab/weekday`
**Назначение:** определение дня недели по дате и подсчёты по дням недели.

Общие правила:
- календарь григорианский, в том числе для дат до 1582 года (пролептический);
- дата — структура `weekday.Date{Year, Month, Day}`; допустимы год >= 1, месяц 1..12, день от 1 до числа дней в месяце;
- дни недели нумеруются по ISO 8601: 1 — понедельник, …, 7 — воскресенье; в пакете есть константы `weekday.Monday` … `weekday.Sunday`.

| № | Функция | Что делает | Особые случаи |
|---|---------|-----------|---------------|
| 1 | `DayOfWeek(d Date) (int, error)` | День недели даты, 1..7 | Несуществующая дата → `0, ErrInvalidDate` |
| 2 | `WeekdayName(n int) (string, error)` | Название дня недели строчными буквами: «понедельник» … «воскресенье» | `n` вне 1..7 → `"", ErrInvalidWeekday` |
| 3 | `IsWeekend(d Date) (bool, error)` | `true` для субботы и воскресенья | Несуществующая дата → `false, ErrInvalidDate` |
| 4 | `CountWeekdayInMonth(year, month, wd int) (int, error)` | Сколько раз день недели `wd` встречается в месяце | `wd` вне 1..7 → `ErrInvalidWeekday`; год < 1 или месяц вне 1..12 → `ErrInvalidDate` |
| 5 | `WorkdaysInMonth(year, month int) (int, error)` | Число дней с понедельника по пятницу в месяце (праздники не учитываются) | Год < 1 или месяц вне 1..12 → `0, ErrInvalidDate` |

Ошибки проверяются через `errors.Is(err, weekday.ErrInvalidDate)` и `errors.Is(err, weekday.ErrInvalidWeekday)`.

Примеры:

```
DayOfWeek(Date{2024, 10, 9})              → 3 (среда), nil
DayOfWeek(Date{2000, 1, 1})               → 6 (суббота), nil
WeekdayName(7)                            → "воскресенье", nil
IsWeekend(Date{2024, 10, 12})             → true, nil
CountWeekdayInMonth(2024, 9, Monday)      → 5, nil
WorkdaysInMonth(2024, 10)                 → 23, nil
```
