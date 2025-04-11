.SILENT:

start:
	docker-compose -f docker/docker-compose.yaml up -d

rebuild:
	docker-compose -f docker/docker-compose.yaml down
	docker-compose -f docker/docker-compose.yaml up --build
