FROM golang:alpine
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN make
CMD ["./main"]