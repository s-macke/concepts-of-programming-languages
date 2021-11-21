
all: stats container_simple1 container_simple2 container_simple3 container_simple4 container

stats: stats.go
	go build stats.go

container_simple0: container_simple0.go
	go build container_simple0.go

container_simple1: container_simple1.go
	go build container_simple1.go

container_simple2: container_simple2.go
	go build container_simple2.go

container_simple3: container_simple3.go
	go build container_simple3.go

container_simple4: container_simple4.go
	go build container_simple4.go

container: container.go
	go build container.go

@phony: run

docker:
	docker run -ti alpine:latest /bin/sh

run:
	./container_simple3 run /bin/sh

rootfs:
	wget "https://dl-cdn.alpinelinux.org/alpine/v3.14/releases/x86_64/alpine-minirootfs-3.14.3-x86_64.tar.gz"
	mkdir -p rootfs
	cd rootfs
	tar -xvzf ../alpine-minirootfs-3.14.3-x86_64.tar.gz
