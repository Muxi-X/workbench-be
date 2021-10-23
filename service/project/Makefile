name=project

build:
	go build -o main
proto:
	protoc --go_out=plugins=micro:. ./proto/project.proto
github:
	git push origin && git push --tags origin
gitea:
	git tag release-${name}-${ver} && git push --tags muxi
push: gitea github
