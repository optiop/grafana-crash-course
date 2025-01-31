test-structure:
	@echo "Each lab directory should have the following structure:"
	@echo ""
	for dir in lab*; do \
		if [ -d "$$dir" ]; then \
			if [ ! -f "$$dir/README.md" ]; then \
				echo "ERROR: $$dir/README.md not found"; \
			fi; \
			if [ ! -d "$$dir/solution" ]; then \
				echo "ERROR: $$dir/solution directory not found"; \
			fi; \
			if [ ! -d "$$dir/exercise" ]; then \
				echo "ERROR: $$dir/exercise directory not found"; \
			fi; \
		fi \
	done

test-exercise:
	@echo "Test instructions for each lab"
	./.scripts/test-exercise.sh

show-slides:
	@echo "Show slides for each lab"
	npm run dev lab.01-cadvisor-prometheus/slides.md 

test:
	@make test-structure
	@make test-exercise
	@make test-solution

help:
	@echo "test - Run all tests"
	@echo "help - Show this help message"