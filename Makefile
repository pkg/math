GOFILES=int.go int_test.go int8.go int8_test.go int16.go int16_test.go int32.go int32_test.go int64.go int64_test.go extra.go extra_test.go
GOFILES+=uint.go uint_test.go uint8.go uint8_test.go uint16.go uint16_test.go uint32.go uint32_test.go uint64.go uint64_test.go
GOFILES+=bigint.go bigint_test.go bigrat.go bigrat_test.go

all: gen test cov README.md

test: gen
	go test .

cov: test
	$(eval TEMP:=$(shell mktemp cover.XXXXXXX))
	go test -coverprofile=$(TEMP) .
	go tool cover -func=$(TEMP)
	rm $(TEMP)

README.md: gen
	godoc2md github.com/pkg/math > $@

gen: $(GOFILES)
	go fmt $^

int8.go: int.go
	sed -e 's/Int/Int8/g' -e 's/int/int8/g' $^ > $@

int16.go: int.go
	sed -e 's/Int/Int16/g' -e 's/int/int16/g' $^ > $@

int32.go: int.go
	sed -e 's/Int/Int32/g' -e 's/int/int32/g' $^ > $@

int64.go: int.go
	sed -e 's/Int/Int64/g' -e 's/int/int64/g' $^ > $@

extra.go: int.go
	sed -e 's/Int//g' $^ > $@

extra_test.go: int_test.go
	sed -e 's/Int//g' $^ > $@

int8_test.go: int_test.go
	sed -e 's/Int/Int8/g' -e 's/int/int8/g' $^ > $@

int16_test.go: int_test.go
	sed -e 's/Int/Int16/g' -e 's/int/int16/g' $^ > $@

int32_test.go: int_test.go
	sed -e 's/Int/Int32/g' -e 's/int/int32/g' $^ > $@

int64_test.go: int_test.go
	sed -e 's/Int/Int64/g' -e 's/int/int64/g' $^ > $@

uint.go: int.go
	sed -e 's/Int/Uint/g' -e 's/ int/ uint/g' -e 's/\.int/\.uint/g' -e 's/\]int/\]uint/g' $^ > $@

uint8.go: uint.go
	sed -e 's/Uint/Uint8/g' -e 's/ uint/ uint8/g' -e 's/\.uint/\.uint8/g' -e 's/\]uint/\]uint8/g' $^ > $@

uint16.go: uint.go
	sed -e 's/Uint/Uint16/g' -e 's/ uint/ uint16/g' -e 's/\.uint/\.uint16/g' -e 's/\]uint/\]uint16/g' $^ > $@

uint32.go: uint.go
	sed -e 's/Uint/Uint32/g' -e 's/ uint/ uint32/g' -e 's/\.uint/\.uint32/g' -e 's/\]uint/\]uint32/g' $^ > $@

uint64.go: uint.go
	sed -e 's/Uint/Uint64/g' -e 's/ uint/ uint64/g' -e 's/\.uint/\.uint64/g' -e 's/\]uint/\]uint64/g' $^ > $@

uint8_test.go: uint_test.go
	sed -e 's/Uint/Uint8/g' -e 's/ uint/ uint8/g' -e 's/\.uint/\.uint8/g' -e 's/\]uint/\]uint8/g' $^ > $@

uint16_test.go: uint_test.go
	sed -e 's/Uint/Uint16/g' -e 's/ uint/ uint16/g' -e 's/\.uint/\.uint16/g' -e 's/\]uint/\]uint16/g' $^ > $@

uint32_test.go: uint_test.go
	sed -e 's/Uint/Uint32/g' -e 's/ uint/ uint32/g' -e 's/\.uint/\.uint32/g' -e 's/\]uint/\]uint32/g' $^ > $@

uint64_test.go: uint_test.go
	sed -e 's/Uint/Uint64/g' -e 's/ uint/ uint64/g' -e 's/\.uint/\.uint64/g' -e 's/\]uint/\]uint64/g' $^ > $@

bigrat.go: bigint.go
	sed -e 's/BigInt/BigRat/g' -e 's/big\.Int/big\.Rat/g' $^ > $@ 

.PHONEY: gen test cov
