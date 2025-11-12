.PHONY: dev
dev:
	go tool air \
		-build.cmd="make build" \
		-build.full_bin="./tmp/scaffold" \
		-build.include_ext="go,html,css,js" \
		-build.exclude_regex="tmp/" \
		-proxy.enabled=true \
		-proxy.proxy_port=3000 \
		-proxy.app_port=4000


.PHONY: build
build:
	go build -o ./tmp/scaffold ./cmd/scaffold/main.go

.PHONY: test
test:
	go tool gotestsum --format testname -- ./test