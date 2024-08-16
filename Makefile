# Define a variable for the year
YEAR=2023

# List of days to create targets for
DAYS=day01 day02

# Target to run a specific day
day%:
	@echo "Running day $*..."
	@cd $(YEAR)/day$* && go run .

# Target to run all days
run-all: $(addprefix run-,$(DAYS))

# Default target
all: run-all
