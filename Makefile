obu:
	@go build -o ./bin/obu ./obu/main.go
	@./bin/obu


receiver:
	@go build -o ./bin/receiver ./data_receiver/main.go
	@./bin/receiver


git:
	@git add .
	@git commit -m"qaytardim evvelki veziyyetine muellim ile birlikde tezeden duzeldecem"
	@git push -u origin main

	
	
.PHONY:obu