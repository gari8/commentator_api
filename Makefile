.PHONY: build

build:
	sam build
clean:
	rm -rf ./commentator/commentator
invoke:
	make build && sam local invoke
run:
	make build && sam local start-api
deploy:
	sam deploy --profile localmaster
deploy-all:
	make build && make deploy