#
# Targets managing life cycle of the k8s objects. They are generic enouh so
# that they can be handled by Minikube deploy and Kubernetes as well (bold assumption).
#

SRC=go-gin-server.yaml

create:
	kubectl apply -f $(SRC)

delete:
	kubectl delete -f $(SRC)

show:
	kubectl get ingress,svc,pods -o wide

test:
	curl http://127.0.0.1:8080/health
