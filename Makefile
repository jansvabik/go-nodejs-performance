.PHONY: go nodejs nodejs_install

go:
	cd go && make build && ./main && cd ..

nodejs_install:
	cd nodejs && npm i && cd ..

nodejs: nodejs_install
	cd nodejs && node app.js && cd ..
