StackName:=AossSample
AossEndpoint:=https://xxxxxxxxxxxxxxxxxxxx.ap-northeast-1.aoss.amazonaws.com
LambdaRole:=arn:aws:iam::123456789012:role/sample-aoss-role

.PHONY: deploy-stack
deploy-stack:
	aws cloudformation deploy --stack-name $(StackName) --template-file template.yaml --capabilities CAPABILITY_NAMED_IAM --output table

.PHONY: delete-stack
delete-stack:
	aws cloudformation delete-stack --stack-name $(StackName)

.PHONY: build
build:
	go mod tidy
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
	zip main.zip main

.PHONY: deploy-lambda
deploy-lambda: build
	aws lambda create-function \
	  --function-name sample-aoss-function \
		--runtime go1.x \
		--role $(LambdaRole) \
		--timeout 30 \
		--handler main \
		--zip-file fileb://main.zip \
		--environment "Variables={AOSS_ENDPOINT=$(AossEndpoint)}"

.PHONY: delete-lambda
delete-lambda:
	aws lambda delete-function \
		--function-name sample-aoss-function

.PHONY: invoke
invoke:
	aws lambda invoke --function-name sample-aoss-function response.json

.PHONY: delete
delete: delete-lambda delete-stack
	rm -f main main.zip response.json
