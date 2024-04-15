#
# Targets for handling minikube testing environment
#

NODES=3
CNI=calico

minikube/start:
	minikube start -n $(NODES) --cni=$(CNI)

minikube/delete:
	minikube delete

minikube/status:
	minikube status

minikube/help:
	@echo
	@echo 'Minikube utility targets'
	@echo
	@echo 'Usage:'
	@echo '    make minikube/start   Start test cluster'
	@echo '    make minikube/delete  Delete test cluster'
	@echo '    make minikube/status  Show test cluster status'
