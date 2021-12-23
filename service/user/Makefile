name=user

build:
	go build -o main
fmt:
	go fmt ./...
vet:
	go vet ./...
protoc:
	protoc --go_out=plugins=micro:. ./proto/user.proto
github:
	git push origin && git push --tags origin
gitea:
	git push --tags muxi
tag:
	git tag release-${name}-${ver}
push: vet tag gitea github