go_files := $(wildcard src/*/*.go)
go_files += cmd/nyx/main.go
src_files := $(filter-out %_test.go, $(go_files))

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	FSWATCH := inotifywatch
else
	FSWATCH := fswatch
endif

TESTS?=.*

.PHONY: all run test cover watch fmt gpm

all: nyx

clean:
	@rm -rf src/golang.org logs nyx

nyx: $(src_files) .git/hooks/pre-commit
	@echo Build nyx
	go build cmd/nyx/main.go
	mv main nyx
	@echo Done

gpm:
	gpm install

nyx.test: $(go_files)
	go test -cover -covermode=count -c ./

test: nyx.test
	@echo Test nyx
	./nyx.test -test.run=$(TESTS) -test.v -test.coverprofile=nyx_cover.out | while read line; \
		do echo $$line | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''; \
	done
	@echo Done testing nyx

run: all
	@./nyx

rerun: all
	./bin/kill-dev-server
	./nyx &

nyx_cover.out: $(go_files)
	@make test

nyx_coverage.html: nyx_cover.out
	go tool cover -html=nyx_cover.out -o nyx_coverage.html

cover: nyx nyx_coverage.html nyx_cover.out
	@echo Command and double click: file:///`pwd`/nyx_coverage.html

watch: rerun
	-@make -s
	@$(FSWATCH) -r . -e='\.idea|target' | while read line; do make rerun -s; done

fmt: $(go_files)
	go fmt *.go

.git/hooks/pre-commit: bin/pre-commit-hook
	@mkdir -p .git/hooks/
	@cp ./bin/pre-commit-hook .git/hooks/pre-commit

gopath:
	mkdir -p $(GOPATH)/src/shpsec.com
	ln -sf $(shell pwd) $(GOPATH)/src/shpsec.com/nyx
