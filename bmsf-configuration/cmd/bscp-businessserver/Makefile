SERVER = bk-bscp-businessserver
ARGS   = run

# debug build
PWD         = $(shell pwd)
LOCALBUILD  = $(PWD)/build
OUTPUT_DIR ?= $(LOCALBUILD)

# debug version
LDVersionFLAG ?= "-X bk-bscp/internal/version.VERSION=DEBUG \
	-X bk-bscp/internal/version.BUILDTIME=DEBUG \
	-X bk-bscp/internal/version.GITHASH=DEBUG"

BINDIR = ${OUTPUT_DIR}/$(SERVER)
BIN    = $(BINDIR)/$(SERVER)

GOENVPATH = $(shell go env GOPATH)
PROTOCOL  = $(GOENVPATH)/src/bk-bscp/internal/protocol
SCRIPTS   = $(GOENVPATH)/src/bk-bscp/scripts

export GO111MODULE=on

default:
	@echo -e "\e[34;1mBuilding $(SERVER)...\033[0m"
	go build -ldflags ${LDVersionFLAG} -gcflags "-N -l" -o $(BIN) main.go
	@cp -r etc $(BINDIR)
	@sh $(SCRIPTS)/daemon-control/generator.sh '$(SERVER)' '$(ARGS)' '$(BINDIR)'
	@echo -e "\e[34;1mBuild $(SERVER) success!\n\033[0m"

clean:
	@rm -rf $(BINDIR) $(LOCALBUILD)

image:
	@cd $(BINDIR) && docker build -t bscp/$(SERVER):$(IMAGE_TAG) -f etc/Dockerfile .
