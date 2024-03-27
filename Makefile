.PHONY:
.SILENT:
include .env

IMAGE_NAME=postgres
CONTAINER_NAME=container

create_docker:
	docker run --name $(CONTAINER_NAME) -e POSTGRES_USER=$(DB_USERNAME) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -p $(DB_PORT):$(DB_PORT) -d $(IMAGE_NAME)

reuse_docker:
	docker start $(CONTAINER_NAME)

clean:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

clean-image:
	docker rmi $(IMAGE_NAME)

