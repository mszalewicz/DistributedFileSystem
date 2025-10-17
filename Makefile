build:
	@go build -o bin/dfs

run: build
	@./bin/dfs

test:
	@go test ./... -v

clean:
	-@rm log
	-@rm bin/dfs
