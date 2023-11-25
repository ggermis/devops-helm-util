PROJECT_NAME  := helm-util

VERSION       := $(shell cat VERSION)
ifdef BUILD_NR
VERSION:="$(VERSION).$(BUILD_NR)"
endif


DIST_DIR := dist
LDFLAGS  := -ldflags "-s -w -X github.com/ggermis/helm-util/pkg/helm_util/version.version=${VERSION}"

all: clean build test


.PHONY: project-name
project-name:
	@echo $(PROJECT_NAME)

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: clean
clean:
	@rm -rf "${DIST_DIR}"

.PHONY: build
build:
	@mkdir -p dist
	@go build ${LDFLAGS} -o "${DIST_DIR}/$(PROJECT_ALIAS)"

.PHONY: test
test:
	#@mkdir -p test-results
	@go install github.com/jstemmer/go-junit-report@latest
	#@go test -v ./... 2>&1 | go-junit-report -set-exit-code > test-results/tests.xml

.PHONY: install
install:
	@go install ${LDFLAGS}

.PHONY: uninstall
uninstall:
	@go clean -i

