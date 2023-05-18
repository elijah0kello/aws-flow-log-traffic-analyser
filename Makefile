all:
	@ mkdir build
	@ GOOS=linux GOARCH=amd64 go build -o ./build/main main.go
	@ chmod 775 ./build/main
	@ zip main.zip ./build/main