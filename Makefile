APP_NAME=ldapadmin
help:
	@echo "  s|start      start ldapadmin"
	@echo "  b|build      build project"
	@echo "  f|format     format project code style"
s start:build
	./output/bin/ldapadmin
b build:
	go build -o output/bin/${APP_NAME} && \
    cp -r conf output/
f format:
	find . -name '*.go' | grep -Ev 'vendor' | xargs gofmt -s -w
