.PHONY: build docker

build:
	go build -o bin/controller ./cmd/controller

docker:
	docker build -t k8s-self-healer:latest .

clean:
	rm -rf bin
