########################################################################################

COVERALLS_TOKEN ?=

########################################################################################

.PHONY = test fmt deps deps-test coveralls

########################################################################################

deps:
	go get -v pkg.re/yaml.v2

deps-test:
	go get -v github.com/axw/gocov/gocov
	go get -v pkg.re/check.v1

test:
	gocov test . | gocov report

fmt:
	find . -name "*.go" -exec gofmt -s -w {} \;
