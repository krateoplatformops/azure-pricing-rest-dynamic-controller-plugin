ARCH?=amd64
REPO?=#your repository here 
VERSION?=0.1

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) go build -o ./bin/azure-pricing-rest-dynamic-controller-plugin main.go

container:
	docker build -t $(REPO)azure-pricing-rest-dynamic-controller-plugin:$(VERSION) .
	docker push $(REPO)azure-pricing-rest-dynamic-controller-plugin:$(VERSION)