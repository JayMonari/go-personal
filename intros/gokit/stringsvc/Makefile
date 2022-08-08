proxy:
	go run . --port=:8997 & 
	go run . --port=:8998 & 
	go run . --port=:8999 & 
	sleep 1
	go run . --proxy=localhost:8997,localhost:8998,localhost:8999

shutdown:
	killall kit

.PHONY: proxy shutdown
