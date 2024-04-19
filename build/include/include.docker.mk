#
# Docker utility targets
#

$(info !docker make: The docker image version is: $(VERSION))

ifneq (${GODEBUG},)
    $(info !docker make: Building debug docker image with flags: ${GODEBUG})
    BUILD_ARGS=--build-arg VERSION=${VERSION} \
        --build-arg GODEBUG=${GODEBUG}
else
    BUILD_ARGS=--build-arg VERSION=${VERSION}
endif

PROJECT_DIR=build/compose

PROJECT_OPTIONS=--project-directory=${PROJECT_DIR} \
	--file ${PROJECT_DIR}/docker-compose.yml \
	--file ${PROJECT_DIR}/docker-compose-prometheus.yml \
	--file ${PROJECT_DIR}/docker-compose-grafana.yml \
	--file ${PROJECT_DIR}/docker-compose-mongodb.yml \
	--file ${PROJECT_DIR}/docker-compose-redis.yml

IMAGE=$(REPOSITORY)/$(TARGET):$(VERSION)
LATEST_IMAGE=$(REPOSITORY)/$(TARGET):latest

docker/image:
	DOCKER_BUILDKIT=1 docker build -f Dockerfile -t $(IMAGE) . $(BUILD_ARGS)
	docker tag $(REPOSITORY)/$(TARGET):$(VERSION) $(LATEST_IMAGE)

docker/image/publish: committed
	DOCKER_BUILDKIT=1 docker push $(LATEST_IMAGE)

docker/clean:
	docker rmi $(TARGET)

docker/compose/up:
	docker-compose ${PROJECT_OPTIONS} up --detach

docker/compose/down:
	docker-compose ${PROJECT_OPTIONS} down

docker/compose/logs:
	docker-compose ${PROJECT_OPTIONS} logs -f

docker/compose/ps:
	docker-compose ${PROJECT_OPTIONS} ps -a

docker/deploy:
	make -C build/deploy start

docker/deploy/stop:
	make -C build/deploy stop

docker/deploy/show:
	make -C build/deploy show

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
	@echo '    make docker/deploy         deploy to k8s'
	@echo '    make docker/deploy/stop    remove deployment of k8s'
	@echo '    make docker/deploy/show    show deployment of k8s'
