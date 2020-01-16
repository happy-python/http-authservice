# Ambassador http authentication service

Ambassador can authenticate incoming requests before routing them to a backing service. In this tutorial, we'll configure Ambassador to use an external third party authentication service.



```shell
$ kubectl apply -f ambassador.yaml
service/ambassador-admin created
clusterrole.rbac.authorization.k8s.io/ambassador unchanged
serviceaccount/ambassador unchanged
clusterrolebinding.rbac.authorization.k8s.io/ambassador unchanged
deployment.apps/ambassador created
service/ambassador created

$ kubectl apply -f httpbin.yaml
service/httpbin created

$ kubectl apply -f authservice.yaml
authservice.getambassador.io/authentication created
service/authservice created
deployment.apps/authservice configured

$ minikube service list
|-------------|----------------------|-----------------------------|
|  NAMESPACE  |         NAME         |             URL             |
|-------------|----------------------|-----------------------------|
| default     | ambassador           | http://192.168.99.100:31077 |
| default     | ambassador-admin     | http://192.168.99.100:31666 |
| default     | authservice          | No node port                |
| default     | httpbin              | No node port                |
| default     | kubernetes           | No node port                |
| kube-system | default-http-backend | http://192.168.99.100:30001 |
| kube-system | kube-dns             | No node port                |
| kube-system | kubernetes-dashboard | No node port                |
|-------------|----------------------|-----------------------------|

$ curl -i http://192.168.99.100:31077/httpbin/ip
HTTP/1.1 401 Unauthorized
content-length: 0
date: Thu, 16 Jan 2020 03:49:16 GMT
server: envoy

$ curl -i -H 'Authorization: 111' http://192.168.99.100:31077/httpbin/ip
HTTP/1.1 403 Forbidden
content-length: 0
date: Thu, 16 Jan 2020 03:49:43 GMT
server: envoy

$ curl -i -H 'Authorization: 123' http://192.168.99.100:31077/httpbin/ip
HTTP/1.1 200 OK
access-control-allow-credentials: true
access-control-allow-origin: *
content-type: application/json
date: Thu, 16 Jan 2020 03:49:51 GMT
referrer-policy: no-referrer-when-downgrade
server: envoy
x-content-type-options: nosniff
x-frame-options: DENY
x-xss-protection: 1; mode=block
content-length: 55
x-envoy-upstream-service-time: 592

{
  "origin": "172.17.0.1, 61.149.4.214, 172.17.0.1"
}


参考文档：

https://www.getambassador.io/user-guide/auth-tutorial/
