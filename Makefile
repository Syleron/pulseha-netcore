.PHONEY: clean get
default: build
build: get
	 if [ ! -d "./bin/" ]; then mkdir ./bin/; fi
	 env GOOS=linux GOARCH=amd64 go build -buildmode=plugin -o ./bin/networking.so ./src/
buildosx: get
	 if [ ! -d "./bin/" ]; then mkdir ./bin/; fi
	 env GOOS=linux GOARCH=amd64 go build -buildmode=plugin -o ./bin/networking.so ./src/
get:
	 go get -d ./src/
install:
	 cp ./bin/networking.so /usr/local/lib/pulseha
clean:
	go clean
