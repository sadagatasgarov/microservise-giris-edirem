obu:
	@go build -o ./bin/obu ./obu/main.go
	@./bin/obu


receiver:
	@go build -o ./bin/receiver ./data_receiver
	@./bin/receiver


git:
	@git add .
	@git commit -m"bundan sonra loglama baslayiriq"
	@git push -u origin main

	
	
.PHONY:obu