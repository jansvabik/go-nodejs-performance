.PHONY: go nodejs nodejs_install tests

go:
	cd go && make build && ./main && cd ..

nodejs_install:
	cd nodejs && npm i && cd ..

nodejs: nodejs_install
	cd nodejs && node app.js && cd ..

tests: nodejs_install
	bash tests/run.sh
