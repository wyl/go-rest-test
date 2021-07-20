
releaser:
	goreleaser --snapshot --skip-publish --rm-dist;

build: releaser
	docker build -t go-rest-test:latest .

run: build
	docker run -d -p 801:80 --restart=always --name go-rest-test go-rest-test
