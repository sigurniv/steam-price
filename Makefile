TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)
export TAG

build:
	make build -C ./service/currency_rate/

deploy:
	envsubst < k8s/currency_rate.yml | kubectl apply -f -
	envsubst < k8s/steamgame.yml | kubectl apply -f -
	envsubst < k8s/steamprice.yml | kubectl apply -f -

stop:
	kubectl delete service currency && kubectl delete deployment currency
	kubectl delete service steamgame && kubectl delete deployment steamgame
	kubectl delete service steamprice && kubectl delete deployment steamprice