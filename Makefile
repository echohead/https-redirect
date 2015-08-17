
all: build

build:
	go build --tags netgo --ldflags '-extldflags "-static"' -installsuffix cgo -o https-redirect .

release: all
	$(eval VERSION := $(shell ./https-redirect --version))
	git tag $(VERSION)
	git push origin --tags
	hub release create -a https-redirect $(VERSION)

