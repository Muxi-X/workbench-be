name=team

fmt:
	go fmt ./...
vet:
	go vet ./...
build:
	go build -o main
proto:
	protoc --go_out=plugins=micro:. ./proto/team.proto
github:
	git push origin && git push --tags origin
gitea:
	git push --tags muxi
tag:
	git tag release-${name}-${ver}
push: vet tag github gitea