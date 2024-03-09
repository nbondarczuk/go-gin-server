#
# Docker utility targets
#

BUILD_ARGS=--build-arg VERSION=${VERSION}

docker/image:
	DOCKER_BUILDKIT=1 docker build -f Dockerfile -t $(REPOSITORY)/$(TARGET):$(VERSION) . $(BUILD_ARGS)

docker/publish: committed
	docker tag $(REPOSITORY)/$(TARGET):$(VERSION) $(TARGET):$(VERSION)
	DOCKER_BUILDKIT=1 docker push $(REPOSITORY)/$(TARGET):$(VERSION)

docker/clean:
	docker rmi $(TARGET)

docker/help:
	@echo
	@echo 'Docker utility targets'
	@echo
	@echo 'Usage:'
	@echo '    make docker/image          create docker image'
	@echo '    make docker/publish        publish docker image'
	@echo '    make docker/clean          clean docker image'
