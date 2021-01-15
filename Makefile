MAKEFLAGS += -s -k

%: database-% backend-% frontend-%
	@

database-%:
	@$(MAKE) -s -C database $*

backend-%:
	@$(MAKE) -s -C backend $*

frontend-%:
	@$(MAKE) -s -C frontend $*

