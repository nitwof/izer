MAPPINGS_DIR = mappings
FONTS_DIR = icons
FONT_TEMPLATE = icons/font.go.tmpl
BINARY = izer
TEST_DIR = test
FIXTURES_DIR = $(TEST_DIR)/fixtures
TESTINPUT_TEMPLATE = $(FIXTURES_DIR)/input.tmpl

GO = go
GOFMT = gofmt
GOMPLATE = gomplate
LINTER = golangci-lint
GORELEASER = goreleaser
BENCHTIME = 10x

MAPPINGS = $(wildcard $(MAPPINGS_DIR)/*_map.json)
FONTS = $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FONTS_DIR)/%_font.go)

TESTINPUTS = $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/iconize.input) \
						 $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/deiconize.input) \
						 $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/deiconize_color.input) \
						 $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/deiconize_dir.input) \
						 $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/deiconize_dir_color.input) \

TESTGOLDENS = $(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/iconize.golden) \
							$(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/iconize_color.golden) \
							$(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/iconize_dir.golden) \
							$(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/iconize_dir_color.golden) \
							$(MAPPINGS:$(MAPPINGS_DIR)/%_map.json=$(FIXTURES_DIR)/%/deiconize.golden) \

default: build

$(FONTS): $(MAPPINGS) $(FONT_TEMPLATE)
	$(GOMPLATE) -d data=$< -f $(FONT_TEMPLATE) -o $@
	$(GOFMT) -w $@

$(FIXTURES_DIR)/%/iconize.input: $(MAPPINGS) $(TESTINPUT_TEMPLATE)
	$(GOMPLATE) -d data=$< -f $(TESTINPUT_TEMPLATE) -o $@

$(FIXTURES_DIR)/%/deiconize.input: $(MAPPINGS) $(TESTINPUT_TEMPLATE)
	ICONS=1 $(GOMPLATE) -d data=$< -f $(TESTINPUT_TEMPLATE) -o $@

$(FIXTURES_DIR)/%/deiconize_color.input: $(MAPPINGS) $(TESTINPUT_TEMPLATE)
	ICONS=1 COLORS=1 $(GOMPLATE) -d data=$< -f $(TESTINPUT_TEMPLATE) -o $@

$(FIXTURES_DIR)/%/deiconize_dir.input: $(MAPPINGS) $(TESTINPUT_TEMPLATE)
	ICONS=1 DIRICONS=1 $(GOMPLATE) -d data=$< -f $(TESTINPUT_TEMPLATE) -o $@

$(FIXTURES_DIR)/%/deiconize_dir_color.input: $(MAPPINGS) $(TESTINPUT_TEMPLATE)
	ICONS=1 DIRICONS=1 COLORS=1 $(GOMPLATE) -d data=$< -f $(TESTINPUT_TEMPLATE) -o $@

$(FIXTURES_DIR)/%/iconize.golden: $(FIXTURES_DIR)/%/iconize.input
	(cat $< | ./$(BINARY) iconize -f=$*) > $@

$(FIXTURES_DIR)/%/iconize_color.golden: $(FIXTURES_DIR)/%/iconize.input
	(cat $< | ./$(BINARY) iconize -f=$* -c) > $@

$(FIXTURES_DIR)/%/iconize_dir.golden: $(FIXTURES_DIR)/%/iconize.input
	(cat $< | ./$(BINARY) iconize -f=$* -d) > $@

$(FIXTURES_DIR)/%/iconize_dir_color.golden: $(FIXTURES_DIR)/%/iconize.input
	(cat $< | ./$(BINARY) iconize -f=$* -d -c) > $@

$(FIXTURES_DIR)/%/deiconize.golden: $(FIXTURES_DIR)/%/iconize.input
	(cat $< | ./$(BINARY) deiconize) > $@

generate: $(FONTS) $(TESTINPUTS)

download: $(FONTS)
	$(GO) mod download

lint: $(FONTS)
	$(LINTER) run

build: download
	$(GO) build -o $(BINARY) .

test: download
	$(GO) test -cover ./...

cover: download
	$(GO) test -cover -coverprofile=coverage.out ./...
	$(GO) tool cover -func=coverage.out
	@rm coverage.out

goldens: build $(TESTGOLDENS)

itest: build $(TESTINPUTS)
	$(GO) test -v ./$(TEST_DIR)

bench: build $(TESTINPUTS)
	$(GO) test -v -bench=. -benchtime=$(BENCHTIME) ./$(TEST_DIR)

clean:
	@rm -f $(BINARY)

release:
	$(GORELEASER) release --rm-dist

.PHONY: default generate download lint build test cover goldens itest bench clean release
