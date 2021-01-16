MAKEFLAGS += -k

%: database-% backend-% frontend-%
	@

database-%:
	@$(MAKE) -C database $*

backend-%:
	@$(MAKE) -C backend $*

frontend-%:
	@$(MAKE) -C frontend $*

