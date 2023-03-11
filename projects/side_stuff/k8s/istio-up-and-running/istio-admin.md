# Administration of Istio

## Installing Istio

Installing istio on a local cluster is pretty easy with `istioctl`. Just run `istioctl` and it will set up the control plane and ingress gateway. You can uninstall the instio configurations with `istioctl uninstall --purge`. To get an overview of the mesh that is installed, run `istioctl proxy-status`. This will return something like:
```
NAME                                                  CLUSTER        CDS        LDS        EDS        RDS          ECDS         ISTIOD                      VERSION
istio-ingressgateway-6785fcd48-pr5xc.istio-system     Kubernetes     SYNCED     SYNCED     SYNCED     NOT SENT     NOT SENT     istiod-65448977c9-r8hwd     1.16.1
```
This information of understanding how istio if managing the configuration of Envoy deployed as gateways helps in relation to how Envoy instances are managed in the data plane. The data plane, to review, is the sidecar proxies that are deployed with the application. The only side car proxies deployed in the above example belong to the ingress gateway.

## Networking in Istio

The following ports are used by the istio sidecar proxy(Envoy):

| Port	| Protocol	| Description | Pod-internal only |
| :--- | :--- | :--- | :--- |
| 15000	| TCP | Envoy admin port (commands/diagnostics) | Yes |
| 15001	| TCP | Envoy outbound | No |
| 15004	| HTTP | Debug port	| Yes |
| 15006	| TCP | Envoy inbound | No | 
| 15008	| H2 | HBONE mTLS tunnel port | No |
| 15009	| H2C | HBONE port for secure networks | No |
| 15020	| HTTP | Merged Prometheus telemetry from Istio agent, Envoy, and application | No |
| 15021	| HTTP | Health checks | No |
| 15053	| DNS |DNS port, if capture is enabled | Yes |
| 15090	| HTTP |	Envoy Prometheus telemetry | No |

And the following ports are used by the istio control plane(istiod):

| Port	| Protocol	| Description | Local host only |
| :--- | :--- | :--- | :--- |
| 443	| HTTPS | Webhooks service port	| No |
| 8080	| HTTP | Debug interface (deprecated, container port only) | No |
| 15010	| GRPC | XDS and CA services (Plaintext, only for secure networks) | No |
| 15012	| GRPC | XDS and CA services (TLS and mTLS, recommended for production use) | No |
| 15014	| HTTP | Control plane monitoring | No |
| 15017 | HTTPS | Webhook container port, forwarded from 443 | No |

## Service Proxy