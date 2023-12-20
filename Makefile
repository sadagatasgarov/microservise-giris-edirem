obu:
	@go build -o ./bin/obu ./obu/main.go
	@./bin/obu


receiver:
	@go build -o ./bin/receiver ./data_receiver/main.go
	@./bin/receiver


git:
	@git add .
	@git commit -m"kanalin limit doldugu ucun datani daaha ardiq qebultmirdi"
	@git push -u origin main

	
.PHONY:obu