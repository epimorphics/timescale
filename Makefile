all: frontend backend Dockerfile
	make vuedep
	make godep
	make vuebuild
	make gobuild
	make dockerbuild
clean:
	rm timescale || true
	rm -r static/* || true
vuedep: frontend
	cd frontend && yarn
vuedev: frontend
	cd frontend && yarn serve
vuebuild: frontend
	cd frontend && yarn build && cp -r dist/* ../static
vue: frontend
	make vuedep
	make vuebuild
godep: backend
	cd backend && dep ensure
godev: backend
	cd backend && JWT_SECRET=devsecret PORT=3001 DB_CONN=postgresql://timescale:testpassword@localhost:5432/timescale?sslmode=disable gin
gobuild: backend
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o timescale ./backend/*.go
dockerbuild: backend frontend Dockerfile
	docker build -t timescale .
dockerrun: backend frontend Dockerfile
	docker run -it --rm --net=host -e DB_CONN=postgresql://timescale:testpassword@localhost:5432/timescale?sslmode=disable timescale
dockerlogin:
	aws ecr get-login --no-include-email --region eu-west-1 | sh
