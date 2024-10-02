
.PHONY: clean
clean:
	rm -f go-unit-tests.out go-blackbox-tests.out coverage.out
	rm -rf cover-report

.PHONY: go-unit-tests
go-unit-tests:
	@echo "\033[0;32mRun unit tests...\033[0m"
	go test -cover -coverprofile=go-unit-tests.out ./coverage/...

	@echo "\033[0;32mTest coverage...\033[0m"
	go tool cover -func go-unit-tests.out

.PHONY: blackbox-server-up
blackbox-server-up:
	mkdir -p cover-report

	@echo "\033[0;32mServer starting...\033[0m"
	GOCOVERDIR=./cover-report go run -cover ./coverage

.PHONY: go-blackbox-tests
go-blackbox-tests:
	@echo "\033[0;32mRun blackbox tests...\033[0m"
	curl -X POST http://localhost:8080/numbers:is_even -d '{ "number": 10 }'

.PHONY: go-blackbox-tests-error
go-blackbox-tests-error:
	@echo "\033[0;32mRun blackbox tests...\033[0m"
	curl -X POST http://localhost:8080/numbers:is_even -d '<number>10</number>'

.PHONY: go-blackbox-cover
go-blackbox-cover:
	@echo "\033[0;32mTest coverage...\033[0m"
	go tool covdata textfmt -i ./cover-report -o go-blackbox-tests.out
	go tool cover -func go-blackbox-tests.out

.PHONY: go-cover
go-cover:
	@echo "\033[0;32mTest coverage...\033[0m"
	find go-*-tests.out | head -n1 | xargs -I{} head -n1 {} > coverage.out
	find go-*-tests.out | xargs -I{} tail -n+2 {} >> coverage.out
	go tool cover -func coverage.out
