name=feed

build:
	go build -o main
fmt:
	go fmt ./...
vet:
	go vet ./...
proto:
	protoc --go_out=plugins=micro:. ./proto/feed.proto
github:
	git push origin && git push --tags origin
gitea:
	git push --tags muxi
tag:
	git tag release-${name}-${ver}
push: vet tag gitea github