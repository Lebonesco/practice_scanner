FROM golang

# create a local workspace to copy files
ADD . /go/src/github.com/Lebonesco/
WORKDIR /src/github.com/Lebonesco

RUN go get github.com/Lebonesco/practice_scanner
WORKDIR /src/github.com/Lebonesco/practice_scanner

RUN ls
CMD ["$PWD"]