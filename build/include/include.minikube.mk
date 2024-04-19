#
# Targets for handling minikube testing environment
#

NODES=1
CNI=calico
IMAGE=$(REPOSITORY)/$(TARGET):$(VERSION)
LATEST_IMAGE=$(REPOSITORY)/$(TARGET):latest

minikube/start:
	minikube start -n $(NODES) --cni=$(CNI)

minikube/delete:
	minikube delete

minikube/status:
	minikube status

minikube/deploy/build:
	minikube image build -t $(LATEST_IMAGE) .

minikube/deploy/create:
	make -C build/deploy create

minikube/deploy/delete:
	make -C build/deploy delete

minikube/deploy/show:
	make -C build/deploy show

minikube/deploy/tunnel:
	make -C build/deploy tunnel

minikube/deploy/test:
	make -C build/deploy test

minikube/deploy/clean:
	minikube image rm $(LATEST_IMAGE)

minikube/help:
	@echo
	@echo 'Minikube utility targets'
	@echo
	@echo 'Usage:'
	@echo '    make minikube/start          Start test cluster'
	@echo '    make minikube/delete         Delete test cluster'
	@echo '    make minikube/status         Show test cluster status'
	@echo '    make minikube/deploy         Create deployment'
	@echo '    make minikube/deploy/delete  Delete deployment'
	@echo '    make minikube/deploy/show    Show deployment'
	@echo '    make minikube/deploy/tunnel  Make tunnel to access the cluster'
	@echo '    make minikube/deploy/test    Test deployment pinging it health check'
	@echo '    make minikube/deploy/clean   Clean deployment images'
