# Используем образ Golang
FROM golang:1.22.5

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем и загружаем зависимости для каждого микросервиса
COPY ./account-microservice/go.mod ./account-microservice/go.sum ./account-microservice/
RUN cd account-microservice && go mod download

COPY ./document-microservice/go.mod ./document-microservice/go.sum ./document-microservice/
RUN cd document-microservice && go mod download

COPY ./hospital-microservice/go.mod ./hospital-microservice/go.sum ./hospital-microservice/
RUN cd hospital-microservice && go mod download

COPY ./timetable-microservice/go.mod ./timetable-microservice/go.sum ./timetable-microservice/
RUN cd timetable-microservice && go mod download

# Копируем исходный код всех микросервисов и файлы конфигурации
COPY ./account-microservice ./account-microservice/
COPY ./document-microservice ./document-microservice/
COPY ./hospital-microservice ./hospital-microservice/
COPY ./timetable-microservice ./timetable-microservice/


# Сборка всех микросервисов
RUN cd account-microservice && go build -o /app/account-microservice/account-microservice ./cmd/main.go
RUN cd document-microservice && go build -o /app/document-microservice/document-microservice ./cmd/main.go
RUN cd hospital-microservice && go build -o /app/hospital-microservice/hospital-microservice ./cmd/main.go
RUN cd timetable-microservice && go build -o /app/timetable-microservice/timetable-microservice ./cmd/main.go

#set config config все равно одинаковый 
#RUN cp /app/account-microservice/.env /app/bin/.env
#RUN cp /app/account-microservice/config.yml /app/bin/config.yml
RUN ls -la /app/account-microservice/
# Запуск всех микросервисов одновременно 
CMD ["sh", "-c", "/app/account-microservice/account-microservice & /app/hospital-microservice/hospital-microservice & /app/timetable-microservice/timetable-microservice & /app/document-microservice/document-microservice"]