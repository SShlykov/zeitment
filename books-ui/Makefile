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

// feat():
// git add . && git commit -m "feat(): " && git push
push:
    git add .
    git commit -m "${commit}"
    git push