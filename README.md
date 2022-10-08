## :student: My Info
* name：张洪宇（noel.zhang）
* num：G20220797070069

## :writing_hand: Homework Modules:
<details>
  <summary>Module 2</summary>

* 编写一个 HTTP 服务器
</details>

<details>
  <summary>Module 8</summary>

* Kubernetes based on Google GKE.
* [manifest](https://github.com/nozhang/homework/tree/master/manifests)
* The SSL certificate store in gcp secret manager. 
* https://noel.srenantong.site/healthz
* https://noel.srenantong.site/metrics
</details>

<details>
  <summary>Module 10</summary>

```Bash
helm install -f values.yaml kube-prometheus-stack prometheus-community/kube-prometheus-stack
helm template grafana/grafana --output-dir ./
kubectl apply -f ./grafana --recursive
```
* Grafana [URL](https://grafana.srenantong.site)
* https://noel.srenantong.site/metrics
* Screenshots Grafana ![image](https://github.com/nozhang/homework/blob/master/images/grafana-screenshot.png)
            Prometheus ![image](https://github.com/nozhang/homework/blob/master/images/pre-screenshot.png)
</details>

<details>
  <summary>Module 12</summary>

### Install istio
```Bash
curl -L https://istio.io/downloadIstio | sh -
cd istio-1.15.1
cp bin/istioctl /usr/local/bin
istioctl install --set profile=demo -y
```
### Install Jaeger
```Bash
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.15/samples/addons/jaeger.yaml
istioctl dashboard jaeger
```
### Deploy httpserver
```Bash
kubectl apply -f namespace.yaml
kubectl label ns securesvc istio-injection=enabled
kubectl apply -f httpserver.yaml
kubectl apply -f external-secrets.yaml
kubectl apply -f istio-specs.yaml
```
* [Manifest](https://github.com/nozhang/homework/tree/master/istio)
* :warning: Istio on GKE should update [firewall](https://istio.io/latest/docs/setup/platform-setup/gke/?_ga=2.20382525.448473504.1665058722-1985876561.1665058722#:~:text=For%20private%20GKE%20clusters)
* https://httpsserver.srenantong.site/healthz
* https://httpsserver.srenantong.site/metrics
* Screenshots Jaeger-tracing ![image](https://github.com/nozhang/homework/blob/master/images/jaeger-tracing.png)
</details>
