generate:
	# Regenerate our model recursively
	go generate ./...

run: #generate
	# Run the Dapr sidecar
	dapr run --app-id gql --dapr-http-port 3500 --dapr-grpc-port 3501 --components-path ./dapr/bindings/components go run server.go