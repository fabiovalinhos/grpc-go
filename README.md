# grpc-go
Tutorial for creating a Go app using grpc    
Teacher: Clement Jean


When you install protocol buffer, putting the path in  .zshrc file
and attach that line
`export PATH="$PATH:/Users/fabiovilalba/go/bin"`   

For create a protos
`protoc -Igreet/proto --go_out=. --go_opt=module=github.com/fabiovalinhos/grpc-go --go-grpc_out=. --go-grpc_opt=module=github.com/fabiovalinhos/grpc-go greet/proto/dummy.proto`