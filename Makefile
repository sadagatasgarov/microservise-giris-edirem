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

git:
	@git add .
	@git commit -m"storeda duzelis"
	@git push -u origin main


.PHONY:obu