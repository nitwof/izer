MAPPINGS_DIR = mappings
GENICONS_DIR = icons
FONT_TEMPLATE = icons/font.go.tmpl
BINARY = iconizer
TEST_DIR = test
FIXTURES_DIR = $(TEST_DIR)/fixtures
TESTINPUT_TEMPLATE = $(FIXTURES_DIR)/input.tmpl

GO = go
GOFMT = gofmt
GOMPLATE = gomplate
LINTER = golangci-lint
BENCHTIME = 10x

MAPPINGS = $(wildcard $(MAPPINGS_DIR)/*_map.json)
GENICONS = $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(GENICONS_DIR)/%_font.go)
TESTINPUTS = $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/iconize.input)
TESTGOLDENS = $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/iconize.golden) \
							$(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/iconize_color.golden) \
							$(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/iconize_dir.golden) \
							$(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/iconize_dir_color.golden)

default: build

$(GENICONS): $(MAPPINGS) $(FONT_TEMPLATE)
	$(GOMPLATE) -d data=$< -f $(FONT_TEMPLATE) -o $@
	$(GOFMT) -w $@

$(TESTINPUTS): $(MAPPINGS) $(TESTINPUT_TEMPLATE)
	@mkdir -p $(@D)
	$(GOMPLATE) -d data=$< -f $(TESTINPUT_TEMPLATE) -o $@

$(FIXTURES_DIR)/%/iconize.golden: $(FIXTURES_DIR)/%/iconize.input
	@mkdir -p $(@D)
	(cat $< | ./$(BINARY) -f=$*) > $@

$(FIXTURES_DIR)/%/iconize_color.golden: $(FIXTURES_DIR)/%/iconize.input
	@mkdir -p $(@D)
	(cat $< | ./$(BINARY) -f=$* -c) > $@

$(FIXTURES_DIR)/%/iconize_dir.golden: $(FIXTURES_DIR)/%/iconize.input
	@mkdir -p $(@D)
	(cat $< | ./$(BINARY) -f=$* -d) > $@

$(FIXTURES_DIR)/%/iconize_dir_color.golden: $(FIXTURES_DIR)/%/iconize.input
	@mkdir -p $(@D)
	(cat $< | ./$(BINARY) -f=$* -d -c) > $@

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

golden: build $(TESTGOLDENS)

itest: build
	$(GO) test -v ./$(TEST_DIR)

bench: build
	$(GO) test -v -bench=. -benchtime=$(BENCHTIME) ./$(TEST_DIR)

clean:
	@rm -f $(BINARY)

.PHONY: default generate download lint test cover golden build itest bench clean
