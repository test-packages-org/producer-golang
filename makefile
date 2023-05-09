VERSION:=1.0.0

build:
	docker build \
		-f dockerfile.producer \
		-t producer:$(VERSION) \
		.