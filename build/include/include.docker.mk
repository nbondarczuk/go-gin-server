#
# Docker utility targets
#

BUILD_ARGS=--build-arg VERSION=${VERSION}
PROJECT_DIR=build/compose

docker/image:
	DOCKER_BUILDKIT=1 docker build -f Dockerfile -t $(REPOSITORY)/$(TARGET):$(VERSION) . $(BUILD_ARGS)

docker/publish: committed
	docker tag $(REPOSITORY)/$(TARGET):$(VERSION) $(TARGET):$(VERSION)
	DOCKER_BUILDKIT=1 docker push $(REPOSITORY)/$(TARGET):$(VERSION)

docker/clean:
	docker rmi $(TARGET)

docker/compose/up:
	docker-compose --project-directory=${PROJECT_DIR} up --detach

docker/compose/down:
	docker-compose --project-directory=${PROJECT_DIR} down

docker/compose/logs:
	docker-compose --project-directory=${PROJECT_DIR} logs -f

docker/compose/ps:
	docker-compose --project-directory=${PROJECT_DIR} ps -a

docker/help:
	@echo
	@echo 'Docker utility targets'
	@echo
	@echo 'Usage:'
	@echo '    make docker/image          create docker image'
	@echo '    make docker/publish        publish docker image'
	@echo '    make docker/clean          clean docker image'
	@echo '    make docker/compose/up     start docker compose'
	@echo '    make docker/compose/down   stop docker compose'
	@echo '    make docker/compose/logs   show logs'
	@echo '    make docker/compose/ps     show processes'
