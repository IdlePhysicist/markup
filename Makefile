# Environment Variables
CGO=0
SRC=.
BUILD=bin
version?="0.0.0"
repo=github.com/IdlePhysicist/markup/cmd
ldflags="-X $(repo).Commit=`git rev-list -1 HEAD | head -c 8` -X $(repo).Version=$(version)"

default: darwin

linux: clean
	env CGO_ENABLED=$(CGO) GOOS=$@ go build -ldflags $(ldflags) -o $(BUILD)/ ./$(SRC)

darwin: clean
	env CGO_ENABLED=$(CGO) GOOS=$@ go build -ldflags $(ldflags) -o $(BUILD)/ ./$(SRC)

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

