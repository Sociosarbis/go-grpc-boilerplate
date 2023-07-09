## prerequisite
```bash
# install task
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin

# install protoc
apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# install mockery
go install github.com/vektra/mockery/v2@v2.30.1

# clone google proto
git clone https://github.com/protocolbuffers/protobuf.git --branch main --depth 1
```