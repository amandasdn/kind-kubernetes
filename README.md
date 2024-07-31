# Kubernetes with Kind

A project for study purposes using Kubernetes.

## Technologies

The following tools and technologies were used:

- Docker
- Kubernetes
- Kind
- Go

## Helpful commands

Create cluster from a yaml file

    kind create cluster --config=k8s/kind.yaml --name=amandacluster 

Create app image

    docker build -t amandasdn/hello-amanda .

Run app image

    docker run --rm -p 80:80 amandasdn/hello-amanda

Push image to docker hub

    docker push amandasdn/hello-amanda:latest

Get deployment history

    kubectl rollout history deployment goserver

Rollout version

    kubectl rollout undo depolyment goserver

Create configs from a yaml file

    kubectl apply -f k8s/pod.yaml
    kubectl apply -f k8s/replicaset.yaml
	
    kubectl apply -f k8s/service.yaml
    kubectl apply -f k8s/deployment.yaml
    kubectl apply -f k8s/configmap/env.yaml
    kubectl apply -f k8s/configmap/pets.yaml
    kubectl apply -f k8s/secret.yaml
    kubectl apply -f k8s/metrics-server.yaml
    kubectl apply -f k8s/hpa.yaml

Foward Cluster IP

    kubectl port-forward svc/goserver-service 8000:80

Get API kubernetes

    kubectl proxy --port=8080
	
Delete deployment

	kubectl delete deploy goserver

Follow pods to see container restarting

	kubectl apply -f k8s/deployment.yaml

	while ($true) {
		kubectl get pods
		Start-Sleep -Seconds 5
	}
