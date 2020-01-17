bin=trainstatus
.PHONY: clean
all: clean $(bin)

trainstatus:
	echo ${CURDIR}
	GOPATH=${CURDIR} CGO_ENABLED=0 go build -o $(bin)

clean:
	rm -fv $(bin)
