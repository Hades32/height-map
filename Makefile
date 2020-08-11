.PHONY: build

build: outdir
	cd cmd/convert-hgt &&\
	go build && \
	mv convert-hgt* ../../build/

outdir:
	mkdir -p build