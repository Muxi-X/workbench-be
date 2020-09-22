FROM golang:1.13
RUN mkdir /app 
ADD ../../ /app/
WORKDIR /app/service/status
RUN make 
CMD ["./main"]