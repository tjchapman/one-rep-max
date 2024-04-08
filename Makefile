run-manually:
	@go run main.go

test:
	@go test -v ./...

compile:
	echo "Compiling for every OS and Platform"
	@GOOS=linux GOARCH=amd64 go build -o bin/gm-linux -v
	@GOOS=windows GOARCH=amd64 go build -o bin/gm.exe -v
	@GOOS=darwin GOARCH=arm64 go build -o bin/gm -v

move-to-bin:
	@cp bin/1rm /usr/local/bin

delete-bin:
	@rm -rf bin

delete-build:
	@rm -rf 1rm	

it: compile move-to-bin delete-bin

build: 
	@go build -o bin/1rm main.go

run-build:
	@bin/1rm

run: build run-build delete-bin