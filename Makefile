build:
	go build -v -a -o release/${GOOS}/${GOARCH}/helloworld

docker:
	docker build -t davetestingaccount/test-for-ci:latest .

test:
	go test ./tests/
