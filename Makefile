# Environment Variables
CGO=0
SRC=.
BUILD=bin
version?="0.0.0"
ldflags="-X main.commit=`git rev-list -1 HEAD | head -c 8` -X main.version=$(version)"

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

