swagger:
	swagger generate spec -o ./docs/swagger.yml --scan-models

build:
	go build cmd/api/api.go