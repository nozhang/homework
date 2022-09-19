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
* https://noel.srenantong.site/headers
## Module10
```Bash
helm install -f values.yaml kube-prometheus-stack prometheus-community/kube-prometheus-stack
helm template grafana/grafana --output-dir ./
kubectl apply -f ./grafana --recursive
```
* Grafana [URL](https://grafana.srenantong.site)
* https://noel.srenantong.site/metrics
* Screenshots ![image](https://github.com/nozhang/homework/images/grafana-screenshot.png)
![image](https://github.com/nozhang/homework/images/pre-screenshot.png)
