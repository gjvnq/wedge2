.PHONY: webui api

webui:
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

server_code: server/server

api: server/server
	cd server && ./server