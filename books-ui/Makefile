HOOK_NAME=pre-push
HOOK_PATH=../.git/hooks/$(HOOK_NAME)

restart:
	docker build -t proxy-ui-vue .
	docker run -d --name proxy-ui-vue --rm -p 5000:80 proxy-ui-vue

stoprestart:
	docker stop proxy-ui-vue
	docker build -t proxy-ui-vue .
	docker run -d --name proxy-ui-vue --rm -p 5000:80 proxy-ui-vue

rmrestart:
	docker image rm proxy-ui-vue
	docker build -t proxy-ui-vue .
	docker stop proxy-ui-vue
	docker run -d --name proxy-ui-vue --rm -p 5000:80 proxy-ui-vue


prehook:
	@echo "Setting up pre-push hook..."
	@rm -f $(HOOK_PATH)
	@echo '#!/bin/sh'								>> $(HOOK_PATH)
	@echo 'echo "Running lint for books-ui"'		>> $(HOOK_PATH)
	@echo 'cd ./books-ui'							>> $(HOOK_PATH)
	@echo 'yarn lint'						    	>> $(HOOK_PATH)
	@echo 'if [ $$? -ne 0 ]; then'					>> $(HOOK_PATH)
	@echo '  echo "Tests failed, push aborted."'	>> $(HOOK_PATH)
	@echo '  exit 1'								>> $(HOOK_PATH)
	@echo 'fi' 										>> $(HOOK_PATH)
	@echo 'echo "Running tests for proxy-ui"'		>> $(HOOK_PATH)
	@echo 'yarn test_ones'						  	>> $(HOOK_PATH)
	@echo 'if [ $$? -ne 0 ]; then'					>> $(HOOK_PATH)
	@echo '  echo "Tests failed, push aborted."'	>> $(HOOK_PATH)
	@echo '  exit 1'								>> $(HOOK_PATH)
	@echo 'fi' 										>> $(HOOK_PATH)
	@echo 'exit 0' 									>> $(HOOK_PATH)
	@chmod +x $(HOOK_PATH)
	@echo "Pre-push hook set successfully."



push:
ifeq ($(commit),)
	$(error mn is not set)
endif
	make prehook
	git add . && git commit -m "feat($(commit)):" && git push

