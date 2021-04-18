.PHONY: foo bar

foo:
	kubectl apply -f foo/deployment.yaml && kubectl apply -f foo/service.yaml

bar:
	kubectl apply -f bar/deployment.yaml && kubectl apply -f bar/service.yaml

ingress:
	kubectl apply -f ingress.yaml

enable-ingress:
	minikube addons enable ingress

ip:
	minikube ip