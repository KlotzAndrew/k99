default: login build

.PHONY: login
login:
	docker login -u klotzandrew

.PHONY: build
build:
	docker build -f Dockerfile-backend -t klotzandrew/k99-backend .
	docker push klotzandrew/k99-backend

	docker build -f Dockerfile-repo -t klotzandrew/k99-repo .
	docker push klotzandrew/k99-repo

	docker build -f Dockerfile-frontend -t klotzandrew/k99-frontend .
	docker push klotzandrew/k99-frontend

.PHONY: lint
lint:
	helm lint ./chart

.PHONY: dryrun
dryrun:
	helm install --debug --dry-run ./chart

.PHONY: deploy
deploy:
	@echo "deplying !"
	./deploy.sh
