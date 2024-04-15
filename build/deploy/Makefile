SRC=go-gin-server.yaml

load:
	minikube image load go-gin-server:latest

start: load
	kubectl apply -f $(SRC)

stop:
	kubectl delete -f $(SRC)

show:
	kubectl get svc,pods -o wide
