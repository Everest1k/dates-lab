# dates-lab — программа для работы с датами

Практическая работа «Модульное и мутационное тестирование программного продукта».

| Пакет | Назначение | Разработчик | Тестировщик |
|-------|-----------|-------------|-------------|
| `datediff` | проверка дат и разница между датами: IsLeapYear, DaysInMonth, IsValid, DayOfYear, DaysBetween | Участник А | Участник Б |
| `weekday` | день недели по дате: DayOfWeek, WeekdayName, IsWeekend, CountWeekdayInMonth, WorkdaysInMonth | Участник Б | Участник А |

Спецификации: `docs/SPEC_datediff.md`, `docs/SPEC_weekday.md`.

## Требования
- Go 1.25 или новее (`brew install go`)
- Gremlins для мутационного тестирования (`brew tap go-gremlins/tap && brew install gremlins`)

## Команды
```bash
make test        # все тесты
make cover       # покрытие по функциям
make html        # HTML-отчёт о покрытии
make mutate-a    # Gremlins для datediff
make mutate-b    # Gremlins для weekday
make manual-a    # ручные мутанты datediff
make manual-b    # ручные мутанты weekday
```
