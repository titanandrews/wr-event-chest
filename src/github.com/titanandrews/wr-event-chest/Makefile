EXE_NAME	    = wr-event-chest
BIN_SUBDIR		= ./bin
EXE_BIN		    = $(BIN_SUBDIR)/$(EXE_NAME)
TEST_SUBDIR		= ./test

.PHONY: all test test-clean clean

all: $(EXE_BIN) $(TEST_SUBDIR)

$(BIN_SUBDIR)/:
	mkdir -p $@

$(EXE_BIN): test | $(BIN_SUBDIR)/
	go build -i -o $(EXE_BIN) .

$(TEST_SUBDIR):
	@mkdir -p $(dir $@)
	go test

test-clean:
	rm -rf $(TEST_SUBDIR)

clean: test-clean
	rm -rf $(BIN_SUBDIR)
	go clean -i ./...