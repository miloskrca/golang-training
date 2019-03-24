Q ?= 1200

run:
	go run app/main.go -q $Q

setup:
	docker-compose up

load:
	while sleep 1; do \
		curl -d '{ "type": "message_type", "payload": { "message": "hi" } }' -H "Content-Type: application/json" -X POST http://localhost:8000; \
	done

consume:
	go run consumer/main.go -q $Q

pprof:
	go build -o bin app/main.go
	./bin -q $Q -cpuprofile profile.cpu

pprof-top:
	pprof -top bin profile.cpu

pprof-web:
	pprof -http=localhost:8888 bin profile.cpu