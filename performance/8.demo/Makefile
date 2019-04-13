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

pprof-heap:
	pprof -http :3000 http://localhost:8000/debug/pprof/heap

pprof-cpu:
	pprof -http :3000 http://localhost:8000/debug/pprof/profile?seconds=30
