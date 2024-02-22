#
# This file provides all common test targets.
#

.PHONY: test
test:
	go test ./...

test/verbose:
	go test -v -cover ./...

test/cover:
	go test -v -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

test/run:
	make -C ./test/run

test/integration:
	make -C ./test/integration all

test/performance:
	make -C ./test/performance all

test/help:
	@echo
	@echo 'Test utility targets'
	@echo
	@echo 'Usage:'
	@echo '    make test/unit             run unit tests'
	@echo '    make test/unit/verbose     run unit tests with verbose level'
	@echo '    make test/unit/cover       run unit tests with coverage and start default broweser'
	@echo '    make test/run              run run tests'
	@echo '    make test/integration      run integration tests with docker compose'
	@echo '    make test/performance      run performance tests with docker compose'
