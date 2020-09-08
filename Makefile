# Environment Variables
CGO=0
SRC=cmd/.
BUILD=bin
version?="0.0.0"
repo=github.com/IdlePhysicist/markup/cmd/cmds
ldflags="-X $(repo).Commit=`git rev-list -1 HEAD | head -c 8` -X $(repo).Version=$(version)"

default: build

build: clean
	env CGO_ENABLED=$(CGO) go build -ldflags $(ldflags) -o $(BUILD)/markup ./$(SRC)

clean:
	if [[ ! -d $(BUILD) ]]; then \
		mkdir ./$(BUILD); \
	fi
	rm -f $(BUILD)/*
	touch $(BUILD)/.keep

install:
	if [[ ! -d $(HOME)/.config/markup ]]; then \
		mkdir -p $(HOME)/.config/markup; \
	fi
	mv $(BUILD)/markup $(GOPATH)/bin/.	

