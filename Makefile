lint:
	spectral lint openapi.yaml

publish:
	yarn npm publish

testspec:
	go test -v

buildspec:
	go run .
