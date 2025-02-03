define extract_targets
	$(shell ls | grep -E '^lab|^session' | tr -d ' ' | xargs)
endef

TARGETS := $(strip $(call extract_targets))
CHOICE := ""

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

test:
	@make test-structure
	@make test-exercise
	@make test-solution


menu: ## Makefile Interactive Menu
	@# Check if fzf is installed
	@if command -v fzf >/dev/null 2>&1; then \
		echo "Using fzf for selection..."; \
		echo "$(TARGETS)" | tr ' ' '\n' | fzf > .selected_target; \
		target_choice=$$(cat .selected_target); \
	else \
		echo "fzf not found, using numbered menu:"; \
		echo "$(TARGETS)" | tr ' ' '\n' > .targets; \
		awk '{print NR " - " $$0}' .targets; \
		read -p "Enter choice: " choice; \
		target_choice=$$(awk 'NR == '$$choice' {print}' .targets); \
	fi; \
	if [ -n "$$target_choice" ]; then \
		$(MAKE) CHOICE=$$target_choice show-slides; \
	else \
		echo "Invalid choice"; \
	fi

slides: menu

show-slides:
	@cd $(CHOICE)/slides && npm install && npm run dev

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  test           Run all tests"
	@echo "  help           Show this help message"
	@echo "  show-slides    Show slides for each lab"
