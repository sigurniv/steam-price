TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

test:
	go test ./...

build:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-X main.version=$(TAG) -extldflags "-static"' -o currency_rate
#	go build -ldflags "-X main.version=$(TAG)" -o nats-chat .

pack:
	docker build --no-cache -t $(DOCKER_ID_USER)/currency_rate:$(TAG) .

upload:
	docker push $(DOCKER_ID_USER)/currency_rate:$(TAG)

deploy:
	envsubst < ./../../k8s/currency_rate.yml | kubectl apply -f -

stop:
	kubectl delete service currency && kubectl delete deployment currency