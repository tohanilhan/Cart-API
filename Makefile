help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

## deploy-all: Deploy all services
deploy-all:
	@echo "Deploying all services"
	make deploy-postgres
	sleep 5
	make deploy-api

## deploy-postgres: Deploy postgres service
deploy-postgres:
	@echo "Deploying postgres"
	make build-postgres
	sleep 5
	docker-compose up -d postgres-svc

## deploy-api: Deploy cart-api service
deploy-api:
	@@echo "Deploying cart-api"
	make build-api
	docker-compose up -d cart-api-svc

## build-all: Build all images
build-all:
	@echo "Building all images"
	make build-postgres
	make build-api

## build-postgres: Build postgres image
build-postgres:
	@echo "Building postgres image"
	cd postgres && docker build -t postgres-img:1.0.0 .

## build-api: Build cart-api image
build-api:
	@echo "Building cart-api image"
	cd cart-api && docker build -t cart-api:1.0.0 .

## clean-all: Clean all images and containers
clean-all:
	@echo "Cleaning all images and containers"
	make clean-postgres

## clean-images: Clean all images
clean-images:
	@echo "Cleaning all images"
	docker-compose down --rmi all --remove-orphans

## clean-containers: Clean all containers
clean-containers:
	@echo "Cleaning all containers"
	docker-compose down --volumes --remove-orphans

## clean-postgres: Clean postgres image and container
clean-postgres:
	@echo "Cleaning postgres image and container"
	docker-compose down --rmi all --volumes --remove-orphans postgres-svc
	rm -rf postgres-data

## clean-api: Clean cart-api image and container
clean-api:
	@echo "Cleaning cart-api image and container"
	docker-compose down --rmi all --volumes --remove-orphans cart-api-svc

## clean-volumes: Clean all volumes
clean-volumes:
	@echo "Cleaning all volumes"
	docker-compose down --volumes

## clean-networks: Clean all networks
clean-networks:
	@echo "Cleaning all networks"
	docker network prune -f

## down-all: Down all services
down-all:
	@echo "Downing all services"
	docker-compose down

## down-postgres: Down postgres service
down-postgres:
	@echo "Downing postgres service"
	docker-compose down postgres-svc

## down-api: Down cart-api service
down-api:
	@echo "Downing cart-api service"
	docker-compose down cart-api-svc

## up-all: Up all services
up-all:
	@echo "Uping all services"
	docker-compose up -d
