FROM golang:1.21

WORKDIR /app
# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all the dependencies
RUN go mod download -x

# Compile daemon for hot reloading
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

RUN go install -v golang.org/x/tools/gopls@latest

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go get -u gorm.io/gorm

RUN go get -u gorm.io/driver/postgres

ENTRYPOINT CompileDaemon --build="go build main.go"  --directory="." --color=true --graceful-kill=true 
EXPOSE 8080
