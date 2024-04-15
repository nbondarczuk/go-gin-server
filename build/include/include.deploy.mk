#
# Targets for handling deployment to minikube testing environment
#

deploy/start:
	make -C build/deploy start

deploy/delete:
	make -C build/deploy delete

deploy/show:
	make -C build/deploy show

deploy/tunnel:
	make -C build/deploy tunnel

deploy/test:
	make -C build/deploy test

deploy/help:
	@echo
	@echo 'Deployment run targets'
	@echo
	@echo 'Usage:'
	@echo '    make deploy/start   Start deployment'
	@echo '    make deploy/delete  Delete deployment'
	@echo '    make deploy/show    Show deployment'
	@echo '    make deploy/tunnel  Make tunnel to access the cluster'
	@echo '    make deploy/test    Test deployment with health check via tunnel'
