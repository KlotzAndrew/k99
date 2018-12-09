default: login build

.PHONY: login
login:
	docker login -u klotzandrew

.PHONY: build
build:
	docker build -f Dockerfile -t klotzandrew/k99-base .
	docker push klotzandrew/k99-base
