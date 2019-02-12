run:
	go run ./cmd/test-direct-integration/main.go --api-key foo --endpoint https://test.com

release: pack dist/direct-endpoint-tester-mac dist/direct-endpoint-tester-linux dist/direct-endpoint-tester-windows.exe

clean:
	rm -rf ./dist
	packr clean

pack:
	packr -i ./cmd/test-direct-integration/

dist/:
	mkdir -p dist

dist/direct-endpoint-tester-mac: | dist/
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o $@ ./cmd/test-direct-integration/

dist/direct-endpoint-tester-linux: | dist/
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o $@ ./cmd/test-direct-integration/

dist/direct-endpoint-tester-windows.exe: | dist/
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o $@ ./cmd/test-direct-integration/

