all:
	cmd/version.sh
	go build -o goproxy cmd/main.go cmd/version.go

clean:
	@rm goproxy

install: all
	rm -rf ${TARGET}/bin
	mkdir -p ${TARGET}/bin
	cp proxy.go ${TARGET}/bin/proxy.go
	chmod a+x ${TARGET}/bin/proxy.go

.PHONY: clean test
