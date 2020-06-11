mkdir build
GOOS=windows GOARCH=amd64 go build -o build/fileserver.exe main.go
GOOS=linux GOARCH=amd64 go build -o build/fileserver main.go
GOOS=linux GOARCH=arm go build -o build/fileserver-arm main.go