SRC=go-gin-server.yaml

load:
	minikube image load go-gin-server:latest

start:
	kubectl apply -f $(SRC)

stop:
	kubectl delete -f $(SRC)

show:
	kubectl get svc,pods -o wide

tunnel:
	minikube tunnel

test:
	curl http://127.0.0.1:8080/health

