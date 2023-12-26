obu:
	@go build -o ./bin/obu ./obu/main.go
	@./bin/obu

receiver:
	@go build -o ./bin/receiver ./data_receiver
	@./bin/receiver

calculator:
	@go build -o ./bin/calculator ./distance_calculator
	@./bin/calculator

git:
	@git add .
	@git commit -m"47 invoice agr der basladoq servisdeki sliceni refactoru"
	@git push -u origin main

.PHONY:obu