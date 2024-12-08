# Makefile

# Default goal
.DEFAULT_GOAL := run

# Run a specific day module
run:
	@go run .

day:
	@read -p "Enter day number (e.g. 05): " daynumber; \
	if [ -f "day$${daynumber}/day$${daynumber}.go" ]; then \
		echo "Already found... skipping."; \
	else \
		mkdir -p day$${daynumber}; \
		cp template/day.go day$${daynumber}/day$${daynumber}.go; \
		sed -i "s/dayXX/day$${daynumber}/g" day$${daynumber}/day$${daynumber}.go; \
	fi
