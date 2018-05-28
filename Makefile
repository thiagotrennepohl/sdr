.DEFAULT_GOAL := help

.PHONY: all tools clean goget env env-ip test do-test env-stop test do-cover cover build image help

NAME    = sdr
VERSION = 1.0.0
GOTOOLS = \
	github.com/golang/dep/cmd/dep

all: tools build docker docker-push

tools: ## Install tools for test cover and dep mgmt
	go get -u -v $(GOTOOLS)

clean: ## Remove old binary
	-@rm -f $(NAME)-app; \
	find vendor/* -maxdepth 0 -type d -exec rm -rf '{}' \;

goget: tools ## [tools] Download dependencies
	dep ensure

build: clean  ## [clean test] Build binary file
	CGO_ENABLED=0 go build -v -a -installsuffix cgo -o $(NAME)-app .

docker: ## Build Docker image
	docker build -t=olikoloko/$(NAME)-app:$(VERSION) .
	rm -rf $(NAME)-app

docker-login: #docker login
	echo $(DOCKER_PASSWORD) | docker login -u $(DOCKER_USERNAME) --password-stdin 

docker-push: docker-login # [docker-login] push docker image to docker hub
	docker push olikoloko/$(NAME)-app:$(VERSION)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
