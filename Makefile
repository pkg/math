GOFILES=int.go uint.go

all: $(GOFILES)

uint.go: int.go
	gofmt -r 'MaxInt -> MaxUint' $^ | gofmt -r 'int -> uint' > $@
