FROM golang:1.24.3-alpine3.22

# Рабочая директория внутри контейнера
WORKDIR /auth

# # Кэширование зависимостей
COPY go.mod go.sum ./
RUN go mod download

# # Копируем остальное
COPY . .

# Сборка под линукauthс
RUN go build -o main .

CMD ["./main"]