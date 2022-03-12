#!/usr/bin/make -f

SUB_DIRS = $(shell find examples features pkg -type d)

define _run_make =
	for dir in ${SUB_DIRS}; do \
		if [ -f "$${dir}/Makefile" ]; then \
			cd $${dir} > /dev/null; \
			make $(1); \
			cd - > /dev/null; \
		fi; \
	done
endef

tidy:
	@$(call _run_make,tidy)

local:
	@$(call _run_make,local)

unlocal:
	@$(call _run_make,unlocal)

be-update:
	@$(call _run_make,be-update)
