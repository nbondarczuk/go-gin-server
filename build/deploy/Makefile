SRC=go-gin-server.yaml

load:
	minikube image load go-gin-server:latest

create:
	kubectl apply -f $(SRC)

delete:
	kubectl delete -f $(SRC)

show:
	kubectl get ingress,svc,pods -o wide

tunnel:
	minikube tunnel

test:
	curl http://127.0.0.1:8080/health

