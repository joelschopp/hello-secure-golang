.PHONY: docker
docker: ## build docker image
	docker build -t hello .

.PHONY: run
run: ## run the 
	docker run -it -v $(PWD):/root/go/src/github.com/joelschopp/hello-secure-golang/ --rm hello

.PHONY: build
build: ## Build.
	go get
	go build -buildmode=pie -ldflags '-linkmode=external' -o hello

.PHONY: checksec
checksec: build ## run gdb checksec
	gdb -batch -ex "checksec ./hello"
