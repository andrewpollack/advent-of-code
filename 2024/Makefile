# Makefile

# Default goal
.DEFAULT_GOAL := run

# Run a specific day module
run:
	@read -p "Enter the day number to run (e.g., 01): " DAY; \
	day_file="day$${DAY}.go"; \
	if [ ! -f "$$day_file" ]; then \
		echo "Error: $$day_file does not exist."; \
		exit 1; \
	fi; \
	echo "Running $$day_file..."; \
	go run $$day_file
