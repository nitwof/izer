MAPPINGS_DIR = mappings
GENICONS_DIR = icons
ICONS_TEMPLATE = icons/icons.go.tmpl
BINARY = iconizer
TEST_DIR = test
FIXTURES_DIR = test/fixtures

GO = go
GOMPLATE = gomplate
GOFMT = gofmt
LINTER = golangci-lint
BENCHTIME = 10x

MAPPINGS = $(wildcard $(MAPPINGS_DIR)/*_map.json)
GENICONS = $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(GENICONS_DIR)/%_icons.go)

default: build

$(GENICONS): $(MAPPINGS) $(ICONS_TEMPLATE)
	@gomplate -d data=$< -f $(ICONS_TEMPLATE) -o $@
	@gofmt -w $@

generate: $(GENICONS)

download: generate
	@$(GO) mod download

lint: generate
	@$(LINTER) run

build: download
	@$(GO) build -o $(BINARY) .

test: download
	@$(GO) test -cover ./...

cover: download
	@$(GO) test -cover -coverprofile=coverage.out ./...
	@$(GO) tool cover -func=coverage.out
	@rm coverage.out

itest: build
	@$(GO) test -v ./$(TEST_DIR)

bench: build
	@$(GO) test -v -bench=. -benchtime=$(BENCHTIME) ./$(TEST_DIR)

clean:
	@rm -f $(BINARY)

.PHONY: default generate download lint test cover build itest bench clean
