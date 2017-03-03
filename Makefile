test:
	go test --race

cover:
	rm -f *.coverprofile
	go test -coverprofile=session.coverprofile
	go tool cover -html=session.coverprofile
	rm -f *.coverprofile

.PHONY: test cover
