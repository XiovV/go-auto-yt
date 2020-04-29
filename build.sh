go build -ldflags "-s -w" -o golty
./upx --brute golty
