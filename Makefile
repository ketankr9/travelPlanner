bin=trainstatus
.PHONY: clean
all: clean $(bin)

trainstatus:
	CGO_ENABLED=0 go build -o $(bin)

clean:
	rm -fv $(bin)
