FROM golang:1.21 as builder

WORKDIR /app
# Copy the source from the current directory to the Working Directory inside the container
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download -x

# Compile daemon for hot reloading
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

RUN go install golang.org/x/tools/gopls@latest

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

EXPOSE 8080

ENTRYPOINT CompileDaemon --build="go build main.go"  --directory="." --color=true --graceful-kill=true 

