build: build-web build-go
build-web:
	make -C web build
build-go: build-go-windows-amd64 build-go-linux-amd64 build-go-linux-arm64
build-go-%:
	$(eval export CGO_ENABLED := 0)
	$(eval OSARCH := $(subst -, ,$*))
	$(eval export GOOS := $(word 1,$(OSARCH)))
	$(eval export GOARCH := $(word 2,$(OSARCH)))
	go build -ldflags "-s -w" -trimpath -o out/$(GOOS)-$(GOARCH)/
	cp LICENSE COPYRIGHT NOTICE README.md out/$(GOOS)-$(GOARCH)/

clean: clean-web clean-go
clean-web:
	make -C web clean
clean-go:
	rm -rf out

dev:
	nodemon --signal SIGKILL --ext go --exec "(go build -o out/dev.exe && out\dev.exe) || exit 1"
