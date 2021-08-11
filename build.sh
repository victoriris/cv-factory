appName=cvfactory-cli
source=cmd/cvfactory/main.go
targetPath=build/$appName

echo "building..."

# Windows 64 bit
GOOS=windows GOARCH=amd64 go build -o $targetPath-amd64.exe $source

# Windows 32 bit
GOOS=windows GOARCH=386 go build -o $targetPath-386.exe $source



# macOS 64-bit
GOOS=darwin GOARCH=amd64 go build -o $targetPath-amd64-darwin $source

# macOS 32-bit
GOOS=darwin GOARCH=arm64 go build -o $targetPath-386-darwin $source


echo "build finished!"