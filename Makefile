GOFILES=int.go int_test.go int8.go int8_test.go int16.go int16_test.go int32.go int32_test.go int64.go int64_test.go
GOFILES+=uint.go uint_test.go uint8.go uint8_test.go uint16.go uint16_test.go uint32.go uint32_test.go uint64.go uint64_test.go

all: gen test cov

test: gen
	go test .

cov: test
	$(eval TEMP:=$(shell tempfile))
	go test -coverprofile=$(TEMP) .
	go tool cover -func=$(TEMP)
	rm $(TEMP)

gen: $(GOFILES)
	go fmt $^

int8.go: int.go
	gofmt -r 'MaxInt -> MaxInt8' $^ | gofmt -r 'MinInt -> MinInt8' | gofmt -r 'int -> int8' > $@

int16.go: int.go
	gofmt -r 'MaxInt -> MaxInt16' $^ | gofmt -r 'MinInt -> MinInt16' | gofmt -r 'int -> int16' > $@

int32.go: int.go
	gofmt -r 'MaxInt -> MaxInt32' $^ | gofmt -r 'MinInt -> MinInt32' | gofmt -r 'int -> int32' > $@

int64.go: int.go
	gofmt -r 'MaxInt -> MaxInt64' $^ | gofmt -r 'MinInt -> MinInt64' | gofmt -r 'int -> int64' > $@

int8_test.go: int_test.go
	gofmt -r 'MaxInt -> MaxInt8' $^ | gofmt -r 'MinInt -> MinInt8' | gofmt -r 'TestMaxInt -> TestMaxInt8' | gofmt -r 'TestMinInt -> TestMinInt8' | gofmt -r 'maxIntTests -> maxInt8Tests' | gofmt -r 'minIntTests -> minInt8Tests' | gofmt -r 'int -> int8' > $@

int16_test.go: int_test.go
	gofmt -r 'MaxInt -> MaxInt16' $^ | gofmt -r 'MinInt -> MinInt16' | gofmt -r 'TestMaxInt -> TestMaxInt16' | gofmt -r 'TestMinInt -> TestMinInt16' | gofmt -r 'maxIntTests -> maxInt16Tests' | gofmt -r 'minIntTests -> minInt16Tests' | gofmt -r 'int -> int16' > $@

int32_test.go: int_test.go
	gofmt -r 'MaxInt -> MaxInt32' $^ | gofmt -r 'MinInt -> MinInt32 ' | gofmt -r 'TestMaxInt -> TestMaxInt32' | gofmt -r 'TestMinInt -> TestMinInt32' | gofmt -r 'maxIntTests -> maxInt32Tests' | gofmt -r 'minIntTests -> minInt32Tests' | gofmt -r 'int -> int32' > $@

int64_test.go: int_test.go
	gofmt -r 'MaxInt -> MaxInt64' $^ | gofmt -r 'MinInt -> MinInt64' | gofmt -r 'TestMaxInt -> TestMaxInt64' | gofmt -r 'TestMinInt -> TestMinInt64' | gofmt -r 'maxIntTests -> maxInt64Tests' | gofmt -r 'minIntTests -> minInt64Tests' | gofmt -r 'int -> int64' > $@

uint.go: int.go
	gofmt -r 'MaxInt -> MaxUint' $^ | gofmt -r 'MinInt -> MinUint' | gofmt -r 'int -> uint' > $@

uint8.go: uint.go
	gofmt -r 'MaxUint -> MaxUint8' $^ | gofmt -r 'MinUint -> MinUint8' | gofmt -r 'uint -> uint8' > $@

uint16.go: uint.go
	gofmt -r 'MaxUint -> MaxUint16' $^ | gofmt -r 'MinUint -> MinUint16' | gofmt -r 'uint -> uint16' > $@

uint32.go: uint.go
	gofmt -r 'MaxUint -> MaxUint32' $^ | gofmt -r 'MinUint -> MinUint32' | gofmt -r 'uint -> uint32' > $@

uint64.go: uint.go
	gofmt -r 'MaxUint -> MaxUint64' $^ | gofmt -r 'MinUint -> MinUint64' | gofmt -r 'uint -> uint64' > $@

uint8_test.go: uint_test.go
	gofmt -r 'MaxUint -> MaxUint8' $^ | gofmt -r 'TestMaxUint -> TestMaxUint8' | gofmt -r 'maxUintTests -> maxUint8Tests' | gofmt -r 'uint -> uint8' > $@

uint16_test.go: uint_test.go
	gofmt -r 'MaxUint -> MaxUint16' $^ | gofmt -r 'TestMaxUint -> TestMaxUint16' | gofmt -r 'maxUintTests -> maxUint16Tests' | gofmt -r 'uint -> uint16' > $@

uint32_test.go: uint_test.go
	gofmt -r 'MaxUint -> MaxUint32' $^ | gofmt -r 'TestMaxUint -> TestMaxUint32' | gofmt -r 'maxUintTests -> maxUint32Tests' | gofmt -r 'uint -> uint32' > $@

uint64_test.go: uint_test.go
	gofmt -r 'MaxUint -> MaxUint64' $^ | gofmt -r 'TestMaxUint -> TestMaxUint64' | gofmt -r 'maxUintTests -> maxUint64Tests' | gofmt -r 'uint -> uint64' > $@

.PHONEY: gen test cov
