packageName = dictionary.tar.gz
dev:
	go run cmd/dictionary.go
build:
	rm -rf dist && mkdir dist
	cp -r config dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o dist/dictionary ./cmd/dictionary.go
package:build
	cd dist && tar -zcvf $(packageName) ./* && mv ${packageName} ../

.PHONY: dev build package
