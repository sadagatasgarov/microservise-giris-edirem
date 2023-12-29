



48 10.26









localhost:3000/aggregate
{
"value":44.4,
"obuID":4,
"unix":454545
}




sudo apt install -y protobuf-compiler

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

PATH="${PATH}:${HOME}/go/bin"

go get google.golang.org/protobuf 