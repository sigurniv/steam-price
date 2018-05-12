TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

test:
	go test ./...

build:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-X main.version=$(TAG) -extldflags "-static"' -o steamprice
#	go build -ldflags "-X main.version=$(TAG)" -o nats-chat .

pack:
	docker build --no-cache -t $(DOCKER_ID_USER)/steamprice:latest .

upload:
	docker push $(DOCKER_ID_USER)/steamprice:latest

deploy:
	envsubst < ./../../k8s/steamprice.yml | kubectl apply -f -

stop:
	kubectl delete service steamprice && kubectl delete deployment steamprice