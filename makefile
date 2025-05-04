run:
	@go build -o /tmp/main main.go && \
		/tmp/main

dev:
	@air
