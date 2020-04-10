cd proto
protoc --micro_out=. --go_out=. test.proto

protoc --go_out=plugins=grpc:../grpcGWproto test.proto
protoc --grpc-gateway_out=logtostderr=true:../grpcGWproto test.proto

cd .. && cd ..