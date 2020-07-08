FROM golang:1.12.13 
ENV GO111MODULE "on"
WORKDIR $GOPATH/src/muxi-workbench-gateway
COPY . $GOPATH/src/muxi-workbench-gateway
RUN make
EXPOSE 8080
CMD ["./main", "-c", "conf/config.yml"]
