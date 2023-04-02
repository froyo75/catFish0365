# ref: https://vic.demuzere.be/articles/golang-makefile-crosscompile/

BIN_DIR := build
BIN_NAME := catFish0365

default: clean linux darwin windows pack integrity

clean:
	$(RM) $(BIN_DIR)/$(BIN_NAME)*
	go clean -x

install:
	go install

darwin:
	GOOS=darwin GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o '$(BIN_DIR)/$(BIN_NAME)-darwin64'
	GOOS=windows GOARCH=386 go build -ldflags="$(LD_FLAGS)" -o '$(BIN_DIR)/$(BIN_NAME)-darwin32'

linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o '$(BIN_DIR)/$(BIN_NAME)-linux64'
	GOOS=linux GOARCH=386 go build -ldflags="$(LD_FLAGS)" -o '$(BIN_DIR)/$(BIN_NAME)-linux32'

windows:
	GOOS=windows GOARCH=amd64 go build -ldflags="$(LD_FLAGS)" -o '$(BIN_DIR)/$(BIN_NAME)-windows64.exe'
	GOOS=windows GOARCH=386 go build -ldflags="$(LD_FLAGS)" -o '$(BIN_DIR)/$(BIN_NAME)-windows32.exe'

pack:
	cd $(BIN_DIR) && upx $(BIN_NAME)-linux32 $(BIN_NAME)-linux64 $(BIN_NAME)-windows32.exe $(BIN_NAME)-windows64.exe $(BIN_NAME)-darwin32 $(BIN_NAME)-darwin64

integrity:
	cd $(BIN_DIR) && shasum *