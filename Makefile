.PHONY: help build run stop logs clean test docker-build docker-up docker-down docker-logs

help: ## Показать справку
	@echo "Доступные команды:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Собрать приложение локально
	go build -o bin/main .

run: ## Запустить приложение локально
	go run main.go

test: ## Запустить тесты
	go test -v ./...

clean: ## Очистить сборочные артефакты
	rm -rf bin/
	go clean

docker-build: ## Собрать Docker образ
	docker build -t go-hm2-app .

docker-up: ## Запустить все сервисы через Docker Compose
	docker compose up -d

docker-down: ## Остановить все сервисы
	docker compose down

docker-down-v: ## Остановить все сервисы и удалить volumes
	docker compose down -v

docker-logs: ## Показать логи контейнеров
	docker compose logs -f

docker-restart: ## Пересобрать и перезапустить сервисы
	docker compose down
	docker compose up -d --build

docker-ps: ## Показать статус контейнеров
	docker compose ps

# Примеры использования API
create-user: ## Создать тестового пользователя
	curl -X POST http://localhost:8080/api/users \
		-H "Content-Type: application/json" \
		-d '{"name":"John Doe","email":"john@example.com"}'

get-users: ## Получить всех пользователей
	curl http://localhost:8080/api/users

get-metrics: ## Посмотреть метрики
	curl http://localhost:8080/metrics

generate-metrics: ## Сгенерировать тестовые метрики для Grafana
	@echo "Генерация тестовых метрик..."
	@for i in 1 2 3 4 5; do \
		curl -X POST http://localhost:8080/api/users \
			-H "Content-Type: application/json" \
			-d "{\"name\":\"User $$i\",\"email\":\"user$$i@example.com\"}" \
			-s -o /dev/null; \
	done
	@for i in 1 2 3 4 5 6 7 8 9 10; do \
		curl http://localhost:8080/api/users -s -o /dev/null; \
	done
	@echo "Тестовые метрики сгенерированы! Откройте Grafana: http://localhost:3000"

