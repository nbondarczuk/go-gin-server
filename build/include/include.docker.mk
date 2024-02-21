#
# Docker utility targets
#

BUILD_ARGS=--build-arg VERSION=${VERSION}

docker/image:
	docker build -f Dockerfile -t $(TARGET) . $(BUILD_ARGS)

docker/clean:
	docker rmi $(TARGET)

docker/help:
	@echo
	@echo 'Docker utility targets'
	@echo
	@echo 'Usage:'
	@echo '    make docker/image          create docker image'
	@echo '    make docker/clean          clean docker image'
