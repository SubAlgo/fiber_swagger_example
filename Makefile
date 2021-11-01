mkdoc:
	- swag init -g app/main.go --output doc/

serv:
	- go run app/main.go