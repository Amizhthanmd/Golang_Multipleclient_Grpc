# Golang_Multipleclient_Grpc

Step 1:- Install the protocol compiler plugins for Go using the following commands:

→  go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.28   

→  go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

→  go get google.golang.org/grpc

Step 2:- Update your PATH so that the protoc compiler can find the plugins: 

→  export PATH="$PATH:$(go env GOPATH)/bin"

Step 3:- To create the pb file, run the following command:- 

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *your_proto_file_directory/your_proto_file.proto

*(It will automatically create two files with extension of  _grpc.pb.go & .pb.go ) 

Step 4:- To install all the required modules use the following command:-

→  go mod tidy

*notes: 
In gRPC protocol buffers, the supported floating-point types are ‘float’ and ‘double’
→ ‘float’ equivalent to float32
→ ‘double’ equivalent to float64
