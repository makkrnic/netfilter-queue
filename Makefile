# Build all dependencies
deps:
	@echo "==== Getting all dependencies ===="
	@dep ensure -v

rules: 
	@echo "==== Setting IPTABLES rule ===="
	@sudo iptables -A INPUT -p udp --dport 67 -j NFQUEUE --queue-num 0
	@sudo iptables -S
	@echo "Setting configuration file"
	@sudo mkdir /etc/netfilter-queue && sudo cp config.json /etc/netfilter-queue

build: 
	@echo "==== Generating binary files ===="
	@echo "Building main binary" 
	@go build .
	@echo "Building UDP client"
	@cd client/ && go build .
	@echo "Building UDP server"
	@cd server/ && go build .
	
queue: 
	@echo "Building netfilter queue"
	@go build .
	@echo "Running"
	@sudo ./netfilter-queue

client: 
	@echo "Building UDP client"
	@cd client/ && go build .
	@echo "Running client"
	@sudo ./client/client

server: 
	@echo "Running server"
	@echo "Building UDP server"
	@cd server/ && go build .
	@echo "Running client"
	@sudo ./server/server
