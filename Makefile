# Setup:
prerequire:
	curl -fsSL https://moonrepo.dev/install/proto.sh | bash
	proto install go
	yarn install
	yarn add knex -g

install:
	go mod download

# Run application
run: 
	air -c .air.toml

# Build application
build: 
	go build -o build/main