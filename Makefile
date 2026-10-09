.PHONY: test cover html mutate-a mutate-b manual-a manual-b

# Все модульные тесты с подробным выводом
test:
	go test -v ./...

# Покрытие кода по функциям
cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

# HTML-отчёт о покрытии (откроется в браузере)
html: cover
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

# Автоматическое мутационное тестирование (Gremlins)
mutate-a:
	gremlins unleash --timeout-coefficient 50 --workers 1 ./datediff

mutate-b:
	gremlins unleash --timeout-coefficient 50 --workers 1 ./weekday

# Ручные мутанты
manual-a:
	./scripts/manual_mutants.sh mutants/datediff.txt

manual-b:
	./scripts/manual_mutants.sh mutants/weekday.txt
