.PHONY: docker
docker: ## build docker image
	docker build -t hello .

.PHONY: run
run: ## run the 
	docker run -it -v $(PWD):/repo/ --rm hello

.PHONY: build
build: ## Build.
	go build -buildmode=pie -ldflags '-linkmode=external' -o hello

.PHONY: checksec
checksec: build ## run gdb checksec
	gdb -batch -ex "checksec ./hello"
