BINARY_NAME=randomPwd
BUILD_LDFLAGS=-s -w
build:
	#CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="${BUILD_LDFLAGS}" -o "./bin/${BINARY_NAME}_Linux_amd64.bin"  main.go
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags="${BUILD_LDFLAGS}" -installsuffix cgo -o "./bin/${BINARY_NAME}"  main.go
	#CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags="${BUILD_LDFLAGS}" -o "./bin/${BINARY_NAME}_win_amd64.exe"  main.go
	#CGO_ENABLED=1 GOOS=linux GOARCH=arm  go build -ldflags="${BUILD_LDFLAGS}" -o "./bin/${BINARY_NAME}_arm.exe"  main.go
	upx ./bin/${BINARY_NAME}*
run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}
clean:
	go clean
	rm -f "./bin/${BINARY_NAME}*"