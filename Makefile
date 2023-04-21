

build.static.image:
	docker build -t ca-static-img --rm .

create.static.container:
	docker run -t --name ca-static -p 80:80 -d ca-static-img


stop.static.container:
	docker stop ca-static

start.static.container:
	docker start ca-static

rm.static.container:
	docker rm ca-statc --force

hugo.build:
	hugo

hugo.serve:
	hugo serve . -D