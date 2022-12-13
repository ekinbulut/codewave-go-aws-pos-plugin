.PHONY: build

build:
	sam build
run:
	sam build
	sam local start-api
deploy:
	sam deploy --guided
test:
	sam local invoke -e events/event.json	
