FROM golang

# create a local workspace to copy files
ADD . /go/src/github.com/Lebonesco/
WORKDIR /src/github.com/Lebonesco

RUN git clone https://github.com/Lebonesco/practice_scanner.git
WORKDIR /src/github.com/Lebonesco/practice_scanner

ENTRYPOINT ["./main"]

