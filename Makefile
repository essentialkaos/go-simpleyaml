########################################################################################

.PHONY = test fmt deps deps-test

########################################################################################

deps:
	go get -v pkg.re/yaml.v2

deps-test:
	go get -v github.com/axw/gocov/gocov
	go get -v pkg.re/check.v1

test: deps-test
	go test -covermode=count ./...

fmt:
	find . -name "*.go" -exec gofmt -s -w {} \;
