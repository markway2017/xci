.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/xcareteam/xci/cmd/swarm
	go build -o bin/geth     github.com/xcareteam/xci/cmd/geth
	go build -o bin/bootnode github.com/xcareteam/xci/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/xcareteam/xci/swarm/...
