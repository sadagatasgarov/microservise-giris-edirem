obu:
	@go build -o ./bin/obu ./obu/main.go
	@./bin/obu

receiver:
	@go build -o ./bin/receiver ./data_receiver
	@./bin/receiver

calculator:
	@go build -o ./bin/calculator ./distance_calculator
	@./bin/calculator

invoicer:
	@go build -o ./bin/invoicer ./invoicer/main.go
	@./bin/invoicer

git:
	@git add .
	@git commit -m"ders47 aggregator elave dirik"
	@git push -u origin main

.PHONY:obu, invoicer