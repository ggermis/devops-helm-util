PROJECT_NAME  := helm-util

VERSION       := $(shell cat VERSION)
ifdef BUILD_NR
VERSION:="$(VERSION).$(BUILD_NR)"
endif

IMAGE_NAME := germis/${PROJECT_NAME}


all: build


.PHONY: project-name
project-name:
	@echo $(PROJECT_NAME)

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: clean
clean:
	-@docker rmi -f ${IMAGE_NAME}:${VERSION}
	-@docker rmi -f ${IMAGE_NAME}:latest
	-@docker system prune -f

.PHONY: build
build:
	@docker build --rm -t ${IMAGE_NAME}:${VERSION} -t ${IMAGE_NAME}:latest --build-arg VERSION=${VERSION} .

.PHONY: test
test:
	@go install github.com/jstemmer/go-junit-report@latest
	#@mkdir -p test-results
	#@go test -v ./... 2>&1 | go-junit-report -set-exit-code > test-results/tests.xml

.PHONY: publish
publish:
	-@docker push ${IMAGE_NAME}:${VERSION}
	-@docker push ${IMAGE_NAME}:latest

.PHONY: install
install:
	@go install ${LDFLAGS}

.PHONY: uninstall
uninstall:
	@go clean -i

