#!/usr/bin/env bash
# Прогон ручных мутантов: для каждой строки из файла-списка вносит одну правку
# в копию исходника, запускает go test для пакета и печатает KILLED / LIVED.
# После каждого мутанта исходный файл восстанавливается.
#
# Использование:  ./scripts/manual_mutants.sh mutants/stats.txt
# Работает в bash 3.2 (стандартный bash в macOS) и в Linux.

set -u
LIST="${1:?укажите файл со списком мутантов, например mutants/stats.txt}"
cd "$(dirname "$0")/.." || exit 1

killed=0; lived=0; skipped=0
printf '%-5s %-11s %s\n' "ID" "РЕЗУЛЬТАТ" "ОПИСАНИЕ"

while IFS='|' read -r id file line from to desc; do
  case "$id" in ''|\#*) continue ;; esac

  pkg="./$(dirname "$file")"
  backup="$(mktemp)"
  cp "$file" "$backup"

  # Правка только в указанной строке; \Q...\E — строка «было» без регулярных выражений.
  FROM="$from" TO="$to" LN="$line" perl -pi -e \
    'if ($. == $ENV{LN}) { my ($f,$t)=($ENV{FROM},$ENV{TO}); s/\Q$f\E/$t/ }' "$file"

  if cmp -s "$file" "$backup"; then
    printf '%-5s %-11s %s\n' "$id" "НЕ ПРИМЕНЁН" "в строке $line нет текста «${from}»"
    skipped=$((skipped+1))
  elif go test -count=1 -timeout 20s "$pkg" </dev/null >/dev/null 2>&1; then
    printf '%-5s %-11s %s\n' "$id" "LIVED" "$desc"
    lived=$((lived+1))
  else
    printf '%-5s %-11s %s\n' "$id" "KILLED" "$desc"
    killed=$((killed+1))
  fi

  cp "$backup" "$file"; rm -f "$backup"
done < "$LIST"

total=$((killed+lived))
echo "----"
echo "Убито: $killed, выжило: $lived, не применено: $skipped"
if [ "$total" -gt 0 ]; then
  echo "Mutation score: $((100*killed/total))%"
fi
