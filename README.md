# My Info 1.1.0
* name： noel.zhang
* num：G20220797070069

# 🏗 Module Homeworks Description:
## Module2
* 编写一个 HTTP 服务器
## Module8
* Kubernetes based on Google GKE.
* Kubernetes [manifest](https://github.com/nozhang/homework/tree/master/manifests)
* The SSL certificate store in gcp secret manager. 
* https://noel.srenantong.site/healthz
* https://noel.srenantong.site/metrics
## Module10
```Bash
helm install -f values.yaml kube-prometheus-stack prometheus-community/kube-prometheus-stack
helm template grafana/grafana --output-dir ./
kubectl apply -f ./grafana --recursive
```
* Grafana [URL](https://grafana.srenantong.site)
* https://noel.srenantong.site/metrics
* Screenshots Grafana ![image](https://github.com/nozhang/homework/blob/master/images/grafana-screenshot.png)
              Prometheus ![image](https://github.com/nozhang/homework/blob/master/images/pre-screenshot.png)
## Module12
### Install istio
```Bash
curl -L https://istio.io/downloadIstio | sh -
cd istio-1.15.1
cp bin/istioctl /usr/local/bin
istioctl install --set profile=demo -y
```
### Deploy httpserver
```Bash
kubectl apply -f namespace.yaml
kubectl label ns securesvc istio-injection=enabled
kubectl apply -f httpserver.yaml
kubectl apply -f external-secrets.yaml
kubectl apply -f istio-specs.yaml
```
warning istio on GKE update [firewall](https://istio.io/latest/docs/setup/platform-setup/gke/?_ga=2.20382525.448473504.1665058722-1985876561.1665058722#:~:text=For%20private%20GKE%20clusters)
* https://httpsserver.srenantong.site/healthz
* https://httpsserver.srenantong.site/metrics
