MAPPINGS_DIR = mappings
GENICONS_DIR = icons
ICONS_TEMPLATE = icons/icons.go.tmpl
BINARY = iconizer
TEST_DIR = test
TESTINPUTS_TEMPLATE = test/fixtures/input.tmpl
TESTINPUTS_DIR = test/fixtures
TESTGOLDENS_DIR = test/fixtures

GO = go
GOFMT = gofmt
GOMPLATE = gomplate
LINTER = golangci-lint
BENCHTIME = 10x

MAPPINGS = $(wildcard $(MAPPINGS_DIR)/*_map.json)
GENICONS = $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(GENICONS_DIR)/%_icons.go)
TESTINPUTS = $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(TESTINPUTS_DIR)/%.input)
TESTGOLDENS = $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(TESTGOLDENS_DIR)/%.golden) \
							$(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(TESTGOLDENS_DIR)/%_color.golden)

default: build

$(GENICONS): $(MAPPINGS) $(ICONS_TEMPLATE)
	$(GOMPLATE) -d data=$< -f $(ICONS_TEMPLATE) -o $@
	$(GOFMT) -w $@

$(TESTINPUTS): $(MAPPINGS) $(TESTINPUTS_TEMPLATE)
	$(GOMPLATE) -d data=$< -f $(TESTINPUTS_TEMPLATE) -o $@

$(TESTGOLDENS_DIR)/%.golden: $(TESTINPUTS)
	cat $< | ./$(BINARY) -f=$(<F:%.input=%) > $@

$(TESTGOLDENS_DIR)/%_color.golden: $(TESTINPUTS)
	cat $< | ./$(BINARY) -f=$(<F:%.input=%) -c > $@

generate: $(GENICONS) $(TESTINPUTS)

download: $(GENICONS)
	$(GO) mod download

lint: $(GENICONS)
	$(LINTER) run

build: download
	$(GO) build -o $(BINARY) .

test: download
	$(GO) test -cover ./...

cover: download
	$(GO) test -cover -coverprofile=coverage.out ./...
	$(GO) tool cover -func=coverage.out
	@rm coverage.out

gengolden: build $(TESTGOLDENS)

itest: build
	$(GO) test -v ./$(TEST_DIR)

bench: build
	$(GO) test -v -bench=. -benchtime=$(BENCHTIME) ./$(TEST_DIR)

clean:
	@rm -f $(BINARY)

.PHONY: default generate download lint test cover build itest bench clean
