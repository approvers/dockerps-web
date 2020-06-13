FROM golang

RUN mkdir /src
WORKDIR /src

COPY main.go go.mod go.sum template.html ./

RUN cd /src && go build -o main && rm main.go go.mod go.sum

ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ./main
