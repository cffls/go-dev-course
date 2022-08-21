lint:
	golangci-lint run --config ./.golangci.yml

run:
	find exercises -name "*.go" | xargs -I % sh -c 'echo ==== Running exercise: % ==== && go run %'