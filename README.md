To run this project go inside the project folder and run the following commands:
```bash
cd project/

# this ups all the backend microservices
make up_build

# this ups the frontend application
make start
```
> For more info go check the files `Makefile` and `docker-compose.yml`.

For gRPC communication you need to install the following packages globally.
> Check mor info about gRPC protocol [here](https://grpc.io/docs/languages/).
> You can also check the protobuf-compiler docs [here](https://protobuf.dev/installation/).
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

sudo apt-get install protobuf-compiler -y

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs.proto
```
Remember to export the path so the executable can be accessed globally.
```bash
export PATH="$HOME/go/bin:$PATH"
```
