VERSION=1.0.0
DOCKER_IMAGE=anthedung/chatbots

binary:
	docker build -t holayoga-dialogflow-service:build -f Dockerfile.build .
	docker create --name holayoga-dialogflow-service holayoga-dialogflow-service:build  /bin/bash
	docker cp holayoga-dialogflow-service:/build/holayoga-dialogflow-service .
	docker rm holayoga-dialogflow-service

docker:
	docker build --no-cache -t $(DOCKER_IMAGE):$(VERSION) .
	docker tag $(DOCKER_IMAGE):$(VERSION) $(DOCKER_IMAGE):latest

clean:
	rm -f holayoga-dialogflow-service
	docker rm holayoga-dialogflow-service || true

all: clean binary docker
