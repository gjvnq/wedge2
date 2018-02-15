.PHONY: webui api

webui:
	cd webui && npm run build

webui-dev:
	cd webui && npm start

domain_code: domain/*.go
	cd domain && goimports -w *.go
	cd domain && go fmt

server/server: server/*.go domain/*.go
	cd domain && goimports -w *.go
	cd domain && go fmt
	cd server && goimports -w *.go
	cd server && go fmt
	cd server && go build

api: server/server

api-dev: server/server server/main.go
	reflex -s -r '(server|domain)/.*\.go$$' make api
