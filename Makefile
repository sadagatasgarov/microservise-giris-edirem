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

git:
	@git add .
	@git commit -m"Protobuffers ders 50 typesi duzeldik readmi faylin"
	@git push -u origin main


.PHONY:obu