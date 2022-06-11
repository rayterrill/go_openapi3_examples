lint:
	spectral lint openapi.yaml

publish: install
	yarn npm publish

testspec:
	go test -v