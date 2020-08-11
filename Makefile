.PHONY: build

build: outdir
	go build -o convert-hgt$(go env GOEXE) && \
	mv -f convert-hgt* ./build/

outdir:
	mkdir -p build
