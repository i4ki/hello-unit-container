IMAGE=tiago4orion/unit-container-example:latest

all: build push

build:
	go build -o unit-example 

push:
	docker build -t $(IMAGE) .
	docker push $(IMAGE)
