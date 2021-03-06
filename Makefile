.PHONY: webui api dist
IMG_PATH=webui/static/img/

dist: dist/wedge-webui.tgz dist/wedge-api.bin
	
webui: dist/wedge-webui.tgz

dist/wedge-api.bin: server/server
	cp server/server dist/wedge-api.bin

dist/wedge-webui.tgz: webui/config/** webui/src/** webui/static/** webui/test/** $(IMG_PATH)/favicon.ico
	cd webui && npm run build
	cd webui/dist && tar czf ../../dist/wedge-webui.tgz .

webui-dev: webui/config/** webui/src/** webui/static/** webui/test/** $(IMG_PATH)/favicon.ico
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

$(IMG_PATH)/favicon.ico: $(IMG_PATH)/favicon.svg $(IMG_PATH)/apple-icon.png
	cd $(IMG_PATH) && inkscape -z --export-area-page --export-png=favicon-16x16.png --export-width=16 favicon.svg
	cd $(IMG_PATH) && inkscape -z --export-area-page --export-png=favicon-32x32.png --export-width=32 favicon.svg
	cd $(IMG_PATH) && inkscape -z --export-area-page --export-png=favicon-48x48.png --export-width=48 favicon.svg
	cd $(IMG_PATH) && inkscape -z --export-area-page --export-png=favicon-64x64.png --export-width=64 favicon.svg
	cd $(IMG_PATH) && convert favicon-16x16.png favicon-32x32.png favicon-48x48.png favicon-64x64.png favicon.ico
	cd $(IMG_PATH) && rm favicon-16x16.png favicon-32x32.png favicon-48x48.png favicon-64x64.png 

$(IMG_PATH)/apple-icon.png: $(IMG_PATH)/favicon.svg
	cd $(IMG_PATH) && inkscape -z --export-area-page --export-png=apple-icon.png --export-width=76 --export-background=white favicon.svg

api: server/server

api-run: server/server
	cd server && ./server

api-dev: server/server server/main.go
	reflex -s -r '(server|domain)/.*\.(go|json)$$' make api-run
