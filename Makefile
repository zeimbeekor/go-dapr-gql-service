generate:
	# Regenerate our model recursively
	go generate ./...

run:
	# Run the Dapr sidecar
	source .envrc && dapr run --app-id gql --dapr-http-port 3500 --dapr-grpc-port 3501 --components-path ./dapr/bindings/components go run server.go