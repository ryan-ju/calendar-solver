.PHONY : test

clean:
	rm -rf ./dist ./build

fix:
	bash ./scripts/build.sh fix

list:
	bash ./scripts/build.sh list

lint:
	golangci-lint run

test:
	go test -cover -count=1 ./... -coverprofile test-coverage.out
