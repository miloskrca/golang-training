run:
	go run app/main.go -q 600

load:
	for i in 1 .. 700000; do curl -d '{ "type": "message_type", "payload": { "message": "hi" } }' -H "Content-Type: application/json" -X POST http://localhost:8000; done
	# ab -l -p post.json -T application/json -n 30000 -c 5 http://localhost:8000/

pprof:
	go tool pprof http://localhost:8000/debug/pprof/heap