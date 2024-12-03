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

solve-01:
	cat data/day01.txt  | awk '{print $1}' | sort
	cat data/day01.txt  | awk '{print $2}' | sort
