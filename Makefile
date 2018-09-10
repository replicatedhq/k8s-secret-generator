clean:
	rm -rf bin

build:
	mkdir -p bin
	env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o bin/k8s-secret-generator

image: build
	docker build -t replicated/k8s-secret-generator .

push: image
	docker push replicated/k8s-secret-generator
