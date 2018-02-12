

depends:
	go get -u gtihub.com/mjibson/esc
webclient:
	@cd webclient && yarn build
	$(eval JS  := $(shell ls webclient/build/*.js))
	$(eval CSS := $(shell ls webclient/build/*.css))
	rm -f static/index.js && cp $(JS) static/index.js 
	rm -f static/index.css && cp $(CSS) static/index.css
	esc -pkg s2h -o static.go  static

.PHONY: depends webclient