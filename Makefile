.PHONY: build

build:
	sam build
clean:
	rm -rf ./commentator/commentator
invoke:
	sam local invoke
run:
	make build && sam local start-api
package:
	sam package --template-file template.yaml --output-template-file output-template.yaml --s3-bucket commentator-store --profile localmaster
deploy:
	sam deploy --template-file output-template.yaml --stack-name commentator-store --capabilities CAPABILITY_IAM --profile localmaster
deploy-all:
	make build && make package && make deploy