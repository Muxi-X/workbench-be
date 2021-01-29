FROM golang:1.13
RUN mkdir /app 
ADD . /app/
ARG service_name
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY="https://goproxy.cn,direct"
WORKDIR /app/service/$service_name
RUN make 
CMD ["./main"]
