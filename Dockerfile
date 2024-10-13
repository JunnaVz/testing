FROM ubuntu:latest
LABEL authors="pikasoft"

ENTRYPOINT ["top", "-b"]

# Шаг 1: Используем базовый образ с Go
FROM golang:1.21

# Шаг 2: Создаем рабочую директорию внутри контейнера
WORKDIR /app

# Шаг 3: Копируем все файлы проекта в контейнер
COPY . .

## Шаг 4: Устанавливаем зависимости для Go
#RUN go mod t

# Шаг 5: Устанавливаем тестовые зависимости и инструменты
#RUN go get -u github.com/stretchr/testify \
#    && go get -u github.com/golang/mock/gomock

# Шаг 6: Запускаем команду для тестирования
CMD ["go", "test", "./..."]
