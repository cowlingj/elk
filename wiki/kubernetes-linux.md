Setup For arch linux

1. Install `virtualbox` (virtualbox is a minikube peer dependency), `minikube`, `kubernetes`, and `docker` (docker is onlt needed for the frontend, the docker daemon is running on minikube)
> Known good VIRTUALBOX-HOST-MODULES provider: virtualbox-host-modules-arch-6.0.10-6
2. Ensure kvm is setup in bios settings (may be labled as some proprietary virtualisation setting)
3. Enable virtualbox kernel drivers
```
sudo modprobe vboxdrv
```
4. start the systemd docker daemon
5. start minikube
6. `eval minikube docker-env` in current shell
> Step 6 is questionably unecessary

## Debugging services

### Check the service and pod are created successfully
kubectl get/describe

### Check the pod for any logs indicating it's at fault
kubectl logs

### Use the dnsutils and busybox images

kubectl run -i --tty dnsutils --generator=run-pod/v1 --image=tutum/dnsutils -- /bin/sh
kubectl run -i --tty busybox --generator=run-pod/v1 --image=busybox -- /bin/sh

### Check pod and service through proxy
kubectl proxy
curl https://<proxy_addr>/api/v1/namespaces/<namespace>/services/<service>/proxy
curl https://<proxy_addr>/api/v1/namespaces/<namespace>/pods/<pod>/proxy

without /proxy suffix we get configuration, with /proxy with call the endpoint
calling .../namespaces, .../services, .../pods without the name path param following lists all of the given resource

> the pod's proxy endpoint should work because we're not going through the service

### Check the service has automatically created endpoints
kubectl get endpoints <service>

### Check the ip address is correct
kubectl cluster-info
minikube ip

## Dashboards

minikube dashboard

kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta1/aio/deploy/recommended.yaml
kubectl proxy
visit http://<proxy_addr>/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login
