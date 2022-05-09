generate:
	# Regenerate our model recursively
	go generate ./...

run:
	go run server.go