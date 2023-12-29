obu:
	@go build -o ./bin/obu ./obu/main.go
	@./bin/obu

receiver:
	@go build -o ./bin/receiver ./data_receiver
	@./bin/receiver

calculator:
	@go build -o ./bin/calculator ./distance_calculator
	@./bin/calculator

agg:
	@go build -o ./bin/agg ./aggregator
	@./bin/agg


proto:
	protoc --go_out=. --go_opt=paths=source_relative types/ptypes.proto

protogrpc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative types/ptypes.proto

git:
	@git add .
	@git commit -m"Protobuffers grpc duzeldi ders 50 19.25"
	@git push -u origin main


.PHONY:obu