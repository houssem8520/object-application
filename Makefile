#####################
# Tools targets     #
#####################
TOOLS_DIR=$(CURDIR)/tools/bin

.PHONY: tools.clean tools.get

#help tools.clean: remove everything in the tools/bin directory
tools.clean:
	rm -fr $(TOOLS_DIR)/*

#help tools.get: retrieve all the tools specified in gex
tools.get:
	cd $(CURDIR)/tools && go generate tools.go


.PHONY:  build.swagger

VERSION=$(shell cat VERSION)
NAME=object-application

#help build.swagger: 
build.swagger:
	@mkdir -p $(CURDIR)/target
	cp swagger.yaml $(CURDIR)/target/swagger.yaml
	sed "s/#VERSION#/$(VERSION)/g" -i $(CURDIR)/target/swagger.yaml
	 $(TOOLS_DIR)/swagger generate server -f $(CURDIR)/target/swagger.yaml -P models.Principal --name=$(NAME) --exclude-main



#help build.vendor: retrieve all the dependencies used for the project
build.vendor:
	go mod vendor

#help build.vendor.full: retrieve all the dependencies after cleaning the go.sum
build.vendor.full:
	@rm -fr $(CURDIR)/vendor
	go mod tidy
	go mod vendor


generate.sqlite:
	@rm -rf $(CURDIR)/resources/sqlite.db
	sqlite3 $(CURDIR)/resources/sqlite.db < $(CURDIR)/resources/init.sql	

#help generate.mocks: Generate/update mocks for testing
generate.mocks:
	$(CURDIR)/tools/bin/mockgen -source=internal/service/interface.go -destination=$(CURDIR)/internal/service/mock/mock.go
	$(CURDIR)/tools/bin/mockgen -source=internal/repo/interface.go -destination=$(CURDIR)/internal/repo/mock/mock.go


build.local: 
	go build -mod=vendor -o $(CURDIR)/target/run $(CURDIR)/cmd/$(NAME)-server/main.go


run.local:
	@$(CURDIR)/target/run  -c resources/.$(NAME).yaml 

run.docker:
	docker run -d --rm -v $(CURDIR)/resources/:/etc/$(NAME)/ --name $(NAME) docker.io/library/$(NAME):latest --config /etc/$(NAME)/.$(NAME).yaml

#help run.infra: start docker-compose 
run.infra:
	docker-compose -f $(CURDIR)/docker-compose.yaml up -d

#help run.infra.stop: stop docker-compose 
run.infra.stop:
	docker-compose -f $(CURDIR)/docker-compose.yaml down

build.docker:
	DOCKER_BUILDKIT=1 docker build  -t $(NAME):latest -f Dockerfile .

fmt:
	go fmt ./...

test:
	go test ./...
