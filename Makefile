
version = 1.1

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o bin/RedisDBA_linux_amd64
build-mac:
	GOOS=darwin GOARCH=amd64 go build  -o bin/RedisDBA_darwin_amd64
build-windows:
	GOOS=windows GOARCH=amd64 go build  -o bin/RedisDBA_windows_amd64
clean:
	rm -rf ./bin/RedisDBA*




