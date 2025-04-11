.SILENT:

start:
	docker-compose up -d

rebuild:
	docker-compose down
	docker-compose up --build
