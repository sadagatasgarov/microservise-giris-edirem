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
	@git commit -m"ders 47 impemende logmidele=ware anlasilmazligim var 12.07de bitdi yazmaq"
	@git push -u origin main

.PHONY:obu, invoicer