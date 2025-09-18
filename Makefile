IMAGE_NAME := go-api
CONTAINER_NAME := go-api-container

build:
	docker build  --build-arg NUC_DB_VAR=$(NUC_DB) -t  $(IMAGE_NAME) .


run: build
	docker run --name $(CONTAINER_NAME) -p 3000:3000 $(IMAGE_NAME)

stop:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

clean: stop
	docker rmi $(IMAGE_NAME)

logs:
	docker logs -f $(CONTAINER_NAME)

restart: stop run

.PHONY: build run stop clean logs restart
