FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# RUN go get github.com/githubnemo/CompileDaemon
# ENTRYPOINT CompileDaemon --build="-v -o /usr/local/bin/app ./cmd/headcount" --command="sh -c "app""
RUN go build -v -o /usr/local/bin/app ./cmd
