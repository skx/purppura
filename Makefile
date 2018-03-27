#
# Trivial Makefile for the project
#


#
# Build our binary by default
#
all: p


#
# Rebuild our bindata.go file from the assets beneath data/
#
bindata.go: data/
	go-bindata -nomemcopy data/
	go fmt bindata.go


#
# Explicitly update all dependencies
#
deps:
	@for i in `grep -H github.com *.go | awk '{print $$NF}' | sort -u | tr -d \"`; do \
		echo "Updating $$i .." ; go get -u $$i ;\
	done


#
# Build our main binary
#
p: bindata.go $(wildcard *.go)
	go build .


#
# Run our tests
#
test:
	go test -coverprofile fmt

#
# Clean our build
#
clean:
	rm puppet-summary || true

#
# Build and serve
#
serve: all
	./puppet-summary serve -host 0.0.0.0


html:
	go test -coverprofile=cover.out
	go tool cover -html=cover.out -o foo.html
	firefox foo.html
