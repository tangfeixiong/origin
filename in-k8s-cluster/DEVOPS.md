# Trace

## Aug. 1

Logs

    [vagrant@localhost origin]$ docker ps -l
    CONTAINER ID        IMAGE                           COMMAND                  CREATED             STATUS              PORTS               NAMES
    66433e7c5f72        tangfeixiong/openshift-origin   "/usr/bin/openshift s"   11 seconds ago      Up 10 seconds                           k8s_origin.b003dfb1_openshift-f9xy1_openshift-origin_d57b4759-5820-11e6-ab10-0800274654e7_42977bbe
    [vagrant@localhost origin]$ docker logs 66433e7c5f72
    W0801 19:47:56.976456       1 start_master.go:272] assetConfig.loggingPublicURL: Invalid value: "": required to view aggregated container logs in the console
    W0801 19:47:56.979118       1 start_master.go:272] assetConfig.metricsPublicURL: Invalid value: "": required to view cluster metrics in the console
    I0801 19:47:57.001074       1 configgetter.go:129] using watch cache storage (capacity=1000) for clusterpolicies
    I0801 19:47:57.006436       1 configgetter.go:129] using watch cache storage (capacity=1000) for clusterpolicybindings
    I0801 19:47:57.006710       1 configgetter.go:129] using watch cache storage (capacity=1000) for policies
    I0801 19:47:57.007607       1 configgetter.go:129] using watch cache storage (capacity=1000) for policybindings
    I0801 19:47:57.007879       1 configgetter.go:129] using watch cache storage (capacity=1000) for groups
    I0801 19:47:57.008836       1 reflector.go:253] Listing and watching *api.ClusterPolicyBinding from pkg/storage/cacher.go:163
    I0801 19:47:57.017383       1 admission.go:36] Admission plugin "ProjectRequestLimit" is not configured so it will be disabled.
    I0801 19:47:57.018250       1 admission.go:35] Admission plugin "PodNodeConstraints" is not configured so it will be disabled.
    I0801 19:47:57.018614       1 reflector.go:202] Starting reflector *api.LimitRange (0) from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/limitranger/admission.go:154
    I0801 19:47:57.018667       1 admission.go:30] Admission plugin "RunOnceDuration" is not configured so it will be disabled.
    I0801 19:47:57.018688       1 reflector.go:202] Starting reflector *api.Namespace (5m0s) from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/namespace/lifecycle/admission.go:122
    I0801 19:47:57.018701       1 admission.go:35] Admission plugin "PodNodeConstraints" is not configured so it will be disabled.
    I0801 19:47:57.018710       1 admission.go:40] Admission plugin "ClusterResourceOverride" is not configured so it will be disabled.
    I0801 19:47:57.018736       1 reflector.go:202] Starting reflector *api.LimitRange (0) from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/limitranger/admission.go:154
    I0801 19:47:57.019439       1 reflector.go:213] Starting reflector *api.ServiceAccount (0) from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/serviceaccount/admission.go:103
    I0801 19:47:57.019466       1 reflector.go:213] Starting reflector *api.Secret (0) from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/serviceaccount/admission.go:119
    I0801 19:47:57.019564       1 reflector.go:213] Starting reflector *api.SecurityContextConstraints (0) from github.com/openshift/origin/pkg/security/admission/admission.go:51
    I0801 19:47:57.019588       1 admission.go:25] Initializing BuildDefaults plugin with config: &api.BuildDefaultsConfig{TypeMeta:unversioned.TypeMeta{Kind:"", APIVersion:""}, GitHTTPProxy:"", GitHTTPSProxy:"", Env:[]api.EnvVar(nil)}
    I0801 19:47:57.019616       1 admission.go:22] Initializing BuildOverrides plugin with config: &api.BuildOverridesConfig{TypeMeta:unversioned.TypeMeta{Kind:"", APIVersion:""}, ForcePull:false}
    I0801 19:47:57.019645       1 reflector.go:213] Starting reflector *api.SecurityContextConstraints (0) from github.com/openshift/origin/pkg/security/admission/admission.go:51
    I0801 19:47:57.024349       1 reflector.go:213] Starting reflector *api.ResourceQuota (0) from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/resourcequota/resource_access.go:83
    I0801 19:47:57.024404       1 reflector.go:253] Listing and watching *api.ResourceQuota from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/resourcequota/resource_access.go:83
    I0801 19:47:57.039686       1 reflector.go:253] Listing and watching *api.ClusterPolicy from pkg/storage/cacher.go:163
    I0801 19:47:57.042742       1 reflector.go:253] Listing and watching *api.Policy from pkg/storage/cacher.go:163
    I0801 19:47:57.044374       1 reflector.go:253] Listing and watching *api.PolicyBinding from pkg/storage/cacher.go:163
    I0801 19:47:57.045331       1 reflector.go:253] Listing and watching *api.Group from pkg/storage/cacher.go:163
    I0801 19:47:57.048287       1 reflector.go:253] Listing and watching *api.LimitRange from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/limitranger/admission.go:154
    I0801 19:47:57.053820       1 reflector.go:253] Listing and watching *api.Namespace from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/namespace/lifecycle/admission.go:122
    I0801 19:47:57.058344       1 reflector.go:253] Listing and watching *api.LimitRange from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/limitranger/admission.go:154
    I0801 19:47:57.060443       1 reflector.go:253] Listing and watching *api.ServiceAccount from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/serviceaccount/admission.go:103
    I0801 19:47:57.061677       1 reflector.go:253] Listing and watching *api.Secret from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/serviceaccount/admission.go:119
    I0801 19:47:57.062817       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    I0801 19:47:57.064473       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    
    E0801 19:47:57.062865       1 reflector.go:205] github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/namespace/lifecycle/admission.go:122: Failed to list *api.Namespace: Get https://10.3.0.53:8443/api/v1/namespaces?resourceVersion=0: dial tcp 10.3.0.53:8443: getsockopt: connection refused
    E0801 19:47:57.063027       1 reflector.go:205] github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/limitranger/admission.go:154: Failed to list *api.LimitRange: Get https://10.3.0.53:8443/api/v1/limitranges?resourceVersion=0: dial tcp 10.3.0.53:8443: getsockopt: connection refused
    E0801 19:47:57.067868       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: Get https://10.3.0.53:8443/api/v1/securitycontextconstraints?resourceVersion=0: dial tcp 10.3.0.53:8443: getsockopt: connection refused
    E0801 19:47:57.068483       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: Get https://10.3.0.53:8443/api/v1/securitycontextconstraints?resourceVersion=0: dial tcp 10.3.0.53:8443: getsockopt: connection refused
    E0801 19:47:57.069047       1 reflector.go:216] github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/serviceaccount/admission.go:119: Failed to list *api.Secret: Get https://10.3.0.53:8443/api/v1/secrets?fieldSelector=type%3Dkubernetes.io%2Fservice-account-token&resourceVersion=0: dial tcp 10.3.0.53:8443: getsockopt: connection refused
    E0801 19:47:57.069598       1 reflector.go:216] github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/resourcequota/resource_access.go:83: Failed to list *api.ResourceQuota: Get https://10.3.0.53:8443/api/v1/resourcequotas?resourceVersion=0: dial tcp 10.3.0.53:8443: getsockopt: connection refused

    I0801 19:47:57.069829       1 configgetter.go:129] using watch cache storage (capacity=1000) for oauthaccesstokens
    I0801 19:47:57.076505       1 configgetter.go:129] using watch cache storage (capacity=1000) for users
    I0801 19:47:57.077408       1 start_master.go:389] Starting master on 0.0.0.0:8443 (v1.3.0-alpha.2+9ac6923-dirty)
    I0801 19:47:57.078274       1 start_master.go:390] Public master address is https://172.17.4.50:30443
    I0801 19:47:57.079005       1 start_master.go:394] Using images from "openshift/origin-<component>:v1.3.0-alpha.2"
    I0801 19:47:57.079940       1 reflector.go:253] Listing and watching *api.OAuthAccessToken from pkg/storage/cacher.go:163
    I0801 19:47:57.080927       1 reflector.go:253] Listing and watching *api.User from pkg/storage/cacher.go:163

    E0801 19:47:57.077559       1 reflector.go:216] github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/serviceaccount/admission.go:103: Failed to list *api.ServiceAccount: Get https://10.3.0.53:8443/api/v1/serviceaccounts?resourceVersion=0: dial tcp 10.3.0.53:8443: getsockopt: connection refused
    E0801 19:47:57.086736       1 reflector.go:205] github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/limitranger/admission.go:154: Failed to list *api.LimitRange: Get https://10.3.0.53:8443/api/v1/limitranges?resourceVersion=0: dial tcp 10.3.0.53:8443: getsockopt: connection refused

    I0801 19:47:57.091207       1 reflector.go:202] Starting reflector *api.Group (2m0s) from github.com/openshift/origin/pkg/user/cache/groups.go:38
    I0801 19:47:57.093068       1 run_components.go:220] Using default project node label selector: 
    I0801 19:47:57.093633       1 reflector.go:202] Starting reflector *api.Namespace (0) from github.com/openshift/origin/pkg/project/cache/cache.go:84
    I0801 19:47:57.096577       1 configgetter.go:129] using watch cache storage (capacity=1000) for users
    I0801 19:47:57.097390       1 configgetter.go:129] using watch cache storage (capacity=1000) for identities
    I0801 19:47:57.108780       1 reflector.go:213] Starting reflector *api.ClusterPolicy (10m0s) from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:99
    I0801 19:47:57.108883       1 reflector.go:253] Listing and watching *api.ClusterPolicy from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:99
    I0801 19:47:57.109822       1 reflector.go:213] Starting reflector *api.ClusterResourceQuota (10m0s) from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:99
    I0801 19:47:57.109864       1 reflector.go:253] Listing and watching *api.ClusterResourceQuota from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:99
    I0801 19:47:57.111446       1 reflector.go:213] Starting reflector *api.Policy (10m0s) from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:99
    I0801 19:47:57.111616       1 reflector.go:253] Listing and watching *api.Policy from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:99
    I0801 19:47:57.111899       1 reflector.go:213] Starting reflector *api.PolicyBinding (10m0s) from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:99
    I0801 19:47:57.111921       1 reflector.go:253] Listing and watching *api.PolicyBinding from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:99
    
    E0801 19:47:57.112344       1 reflector.go:216] github.com/openshift/origin/pkg/controller/shared/shared_informer.go:99: Failed to list *api.ClusterResourceQuota: Get https://10.3.0.53:8443/oapi/v1/clusterresourcequotas?resourceVersion=0: dial tcp 10.3.0.53:8443: getsockopt: connection refused
    
    I0801 19:47:57.112617       1 reflector.go:253] Listing and watching *api.Group from github.com/openshift/origin/pkg/user/cache/groups.go:38
    I0801 19:47:57.112876       1 reflector.go:213] Starting reflector *api.ClusterPolicyBinding (10m0s) from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:99
    I0801 19:47:57.112899       1 reflector.go:253] Listing and watching *api.ClusterPolicyBinding from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:99
    I0801 19:47:57.112978       1 clusterquotamapping.go:304] Waiting for the caches to sync before starting the quota mapping controller workers
    I0801 19:47:57.113038       1 reflector.go:253] Listing and watching *api.Namespace from github.com/openshift/origin/pkg/project/cache/cache.go:84
    
    E0801 19:47:57.114614       1 reflector.go:205] github.com/openshift/origin/pkg/project/cache/cache.go:84: Failed to list *api.Namespace: Get https://10.3.0.53:8443/api/v1/namespaces?resourceVersion=0: dial tcp 10.3.0.53:8443: getsockopt: connection refused
    
    I0801 19:47:57.114660       1 reflector.go:253] Listing and watching *api.User from pkg/storage/cacher.go:163
    I0801 19:47:57.115314       1 reflector.go:253] Listing and watching *api.Identity from pkg/storage/cacher.go:163
    I0801 19:47:57.122248       1 configgetter.go:129] using watch cache storage (capacity=1000) for builds
    I0801 19:47:57.126585       1 configgetter.go:129] using watch cache storage (capacity=1000) for buildconfigs
    I0801 19:47:57.127156       1 reflector.go:253] Listing and watching *api.Build from pkg/storage/cacher.go:163
    I0801 19:47:57.128266       1 configgetter.go:129] using watch cache storage (capacity=1000) for deploymentconfigs
    I0801 19:47:57.128910       1 reflector.go:253] Listing and watching *api.BuildConfig from pkg/storage/cacher.go:163
    I0801 19:47:57.132687       1 plugin.go:27] Route plugin initialized with suffix=router.default.svc.cluster.local
    I0801 19:47:57.133289       1 configgetter.go:129] using watch cache storage (capacity=1000) for routes
    I0801 19:47:57.133934       1 configgetter.go:129] using watch cache storage (capacity=1000) for hostsubnets
    I0801 19:47:57.134307       1 reflector.go:253] Listing and watching *api.Route from pkg/storage/cacher.go:163
    I0801 19:47:57.132999       1 reflector.go:253] Listing and watching *api.DeploymentConfig from pkg/storage/cacher.go:163
    I0801 19:47:57.135299       1 configgetter.go:129] using watch cache storage (capacity=1000) for netnamespaces
    I0801 19:47:57.137144       1 reflector.go:253] Listing and watching *api.HostSubnet from pkg/storage/cacher.go:163
    I0801 19:47:57.141320       1 configgetter.go:129] using watch cache storage (capacity=1000) for clusternetworks
    I0801 19:47:57.141450       1 configgetter.go:129] using watch cache storage (capacity=1000) for users
    I0801 19:47:57.141496       1 configgetter.go:129] using watch cache storage (capacity=1000) for identities
    I0801 19:47:57.142047       1 configgetter.go:129] using watch cache storage (capacity=1000) for groups
    I0801 19:47:57.142272       1 configgetter.go:129] using watch cache storage (capacity=1000) for policies
    I0801 19:47:57.142315       1 configgetter.go:129] using watch cache storage (capacity=1000) for policybindings
    I0801 19:47:57.142355       1 configgetter.go:129] using watch cache storage (capacity=1000) for clusterpolicies
    I0801 19:47:57.142603       1 configgetter.go:129] using watch cache storage (capacity=1000) for clusterpolicybindings
    I0801 19:47:57.142707       1 configgetter.go:129] using watch cache storage (capacity=1000) for images
    I0801 19:47:57.149476       1 reflector.go:253] Listing and watching *api.Identity from pkg/storage/cacher.go:163
    I0801 19:47:57.154601       1 reflector.go:253] Listing and watching *api.NetNamespace from pkg/storage/cacher.go:163
    I0801 19:47:57.155993       1 reflector.go:253] Listing and watching *api.ClusterNetwork from pkg/storage/cacher.go:163
    I0801 19:47:57.154854       1 reflector.go:253] Listing and watching *api.Group from pkg/storage/cacher.go:163
    I0801 19:47:57.154897       1 reflector.go:253] Listing and watching *api.Policy from pkg/storage/cacher.go:163
    I0801 19:47:57.154921       1 reflector.go:253] Listing and watching *api.PolicyBinding from pkg/storage/cacher.go:163
    I0801 19:47:57.154942       1 reflector.go:253] Listing and watching *api.ClusterPolicy from pkg/storage/cacher.go:163
    I0801 19:47:57.155338       1 reflector.go:202] Starting reflector *api.LimitRange (0) from github.com/openshift/origin/pkg/image/admission/imagestream_limits.go:29
    I0801 19:47:57.161096       1 configgetter.go:129] using watch cache storage (capacity=1000) for imagestreams
    I0801 19:47:57.155356       1 reflector.go:253] Listing and watching *api.Image from pkg/storage/cacher.go:163
    I0801 19:47:57.156847       1 reflector.go:253] Listing and watching *api.User from pkg/storage/cacher.go:163
    I0801 19:47:57.154740       1 reflector.go:253] Listing and watching *api.ClusterPolicyBinding from pkg/storage/cacher.go:163
    I0801 19:47:57.172010       1 configgetter.go:129] using watch cache storage (capacity=1000) for oauthclients
    I0801 19:47:57.172084       1 reflector.go:253] Listing and watching *api.ImageStream from pkg/storage/cacher.go:163
    I0801 19:47:57.172284       1 reflector.go:253] Listing and watching *api.LimitRange from github.com/openshift/origin/pkg/image/admission/imagestream_limits.go:29
    I0801 19:47:57.173882       1 configgetter.go:129] using watch cache storage (capacity=1000) for oauthauthorizetokens
    I0801 19:47:57.173990       1 configgetter.go:129] using watch cache storage (capacity=1000) for oauthaccesstokens
    I0801 19:47:57.174101       1 configgetter.go:129] using watch cache storage (capacity=1000) for oauthclientauthorizations
    I0801 19:47:57.174167       1 reflector.go:253] Listing and watching *api.OAuthAccessToken from pkg/storage/cacher.go:163
    I0801 19:47:57.174474       1 configgetter.go:129] using watch cache storage (capacity=1000) for templates
    I0801 19:47:57.175250       1 configgetter.go:129] using watch cache storage (capacity=1000) for clusterresourcequotas
    
    E0801 19:47:57.183435       1 reflector.go:205] github.com/openshift/origin/pkg/image/admission/imagestream_limits.go:29: Failed to list *api.LimitRange: Get https://10.3.0.53:8443/api/v1/limitranges?resourceVersion=0: dial tcp 10.3.0.53:8443: getsockopt: connection refused
    
    I0801 19:47:57.184360       1 configgetter.go:129] using watch cache storage (capacity=1000) for clusterresourcequotas
    I0801 19:47:57.193111       1 reflector.go:253] Listing and watching *api.OAuthClient from pkg/storage/cacher.go:163
    I0801 19:47:57.193313       1 reflector.go:253] Listing and watching *api.OAuthAuthorizeToken from pkg/storage/cacher.go:163
    I0801 19:47:57.198246       1 reflector.go:253] Listing and watching *api.Template from pkg/storage/cacher.go:163
    I0801 19:47:57.200201       1 reflector.go:253] Listing and watching *api.OAuthClientAuthorization from pkg/storage/cacher.go:163
    I0801 19:47:57.203905       1 reflector.go:253] Listing and watching *api.ClusterResourceQuota from pkg/storage/cacher.go:163
    I0801 19:47:57.206109       1 reflector.go:253] Listing and watching *api.ClusterResourceQuota from pkg/storage/cacher.go:163
    I0801 19:47:57.213133       1 clusterquotamapping.go:304] Waiting for the caches to sync before starting the quota mapping controller workers
    I0801 19:47:57.220534       1 configgetter.go:129] using watch cache storage (capacity=1000) for oauthclients
    I0801 19:47:57.220704       1 configgetter.go:129] using watch cache storage (capacity=1000) for oauthaccesstokens
    I0801 19:47:57.220765       1 configgetter.go:129] using watch cache storage (capacity=1000) for oauthauthorizetokens
    I0801 19:47:57.220828       1 configgetter.go:129] using watch cache storage (capacity=1000) for oauthclientauthorizations
    I0801 19:47:57.220936       1 reflector.go:253] Listing and watching *api.OAuthAccessToken from pkg/storage/cacher.go:163
    I0801 19:47:57.221553       1 reflector.go:253] Listing and watching *api.OAuthClient from pkg/storage/cacher.go:163
    I0801 19:47:57.222549       1 reflector.go:253] Listing and watching *api.OAuthAuthorizeToken from pkg/storage/cacher.go:163
    I0801 19:47:57.227075       1 reflector.go:253] Listing and watching *api.OAuthClientAuthorization from pkg/storage/cacher.go:163
    W0801 19:47:57.256235       1 swagger.go:32] No API exists for predefined swagger description /api/v1
    I0801 19:47:57.267093       1 master.go:272] Started Kubernetes proxy at 0.0.0.0:8443/api/
    I0801 19:47:57.267792       1 master.go:272] Started Origin API at 0.0.0.0:8443/oapi/v1
    I0801 19:47:57.268303       1 master.go:272] Started OAuth2 API at 0.0.0.0:8443/oauth
    I0801 19:47:57.268912       1 master.go:272] Started Web Console 0.0.0.0:8443/console/
    I0801 19:47:57.269383       1 master.go:272] Started Swagger Schema API at 0.0.0.0:8443/swaggerapi/
    I0801 19:47:57.267210       1 net.go:105] Got error &net.OpError{Op:"dial", Net:"tcp4", Source:net.Addr(nil), Addr:(*net.TCPAddr)(0xc821a37230), Err:(*os.SyscallError)(0xc820f80020)}, trying again: "0.0.0.0:8443"
    I0801 19:47:57.313814       1 clusterquotamapping.go:304] Waiting for the caches to sync before starting the quota mapping controller workers
    I0801 19:47:57.398120       1 configgetter.go:129] using watch cache storage (capacity=1000) for clusterpolicies
    I0801 19:47:57.398570       1 reflector.go:213] Starting reflector *api.Namespace (10m0s) from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:87
    I0801 19:47:57.398611       1 reflector.go:253] Listing and watching *api.Namespace from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:87
    I0801 19:47:57.399667       1 reflector.go:253] Listing and watching *api.ClusterPolicy from pkg/storage/cacher.go:163
    I0801 19:47:57.403320       1 ensure.go:231] Ignoring bootstrap policy file because cluster policy found
    I0801 19:47:57.420891       1 clusterquotamapping.go:304] Waiting for the caches to sync before starting the quota mapping controller workers
    I0801 19:47:57.495190       1 ensure.go:193] Ignoring default security context constraints when running on external Kubernetes.
    I0801 19:47:57.521618       1 clusterquotamapping.go:304] Waiting for the caches to sync before starting the quota mapping controller workers
    I0801 19:47:57.641044       1 clusterquotamapping.go:304] Waiting for the caches to sync before starting the quota mapping controller workers
    I0801 19:47:57.743612       1 clusterquotamapping.go:304] Waiting for the caches to sync before starting the quota mapping controller workers
    I0801 19:47:57.825921       1 ensure.go:86] Added build-controller service accounts to the system:build-controller cluster role: <nil>
    I0801 19:47:57.884303       1 clusterquotamapping.go:304] Waiting for the caches to sync before starting the quota mapping controller workers
    I0801 19:47:57.946826       1 ensure.go:86] Added daemonset-controller service accounts to the system:daemonset-controller cluster role: <nil>
    I0801 19:47:57.986983       1 clusterquotamapping.go:304] Waiting for the caches to sync before starting the quota mapping controller workers
    I0801 19:47:58.060238       1 ensure.go:86] Added deployment-controller service accounts to the system:deployment-controller cluster role: <nil>
    I0801 19:47:58.067354       1 reflector.go:253] Listing and watching *api.Namespace from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/namespace/lifecycle/admission.go:122
    I0801 19:47:58.068870       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    I0801 19:47:58.069931       1 reflector.go:253] Listing and watching *api.LimitRange from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/limitranger/admission.go:154
    I0801 19:47:58.077379       1 reflector.go:253] Listing and watching *api.ResourceQuota from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/resourcequota/resource_access.go:83
    I0801 19:47:58.078087       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    I0801 19:47:58.078859       1 reflector.go:253] Listing and watching *api.Secret from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/serviceaccount/admission.go:119
    I0801 19:47:58.105891       1 clusterquotamapping.go:304] Waiting for the caches to sync before starting the quota mapping controller workers
    I0801 19:47:58.111169       1 reflector.go:253] Listing and watching *api.ServiceAccount from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/serviceaccount/admission.go:103
    I0801 19:47:58.111677       1 reflector.go:253] Listing and watching *api.LimitRange from github.com/openshift/origin/vendor/k8s.io/kubernetes/plugin/pkg/admission/limitranger/admission.go:154
    
    E0801 19:47:58.112493       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: the server could not find the requested resource
    
    I0801 19:47:58.135349       1 reflector.go:253] Listing and watching *api.Namespace from github.com/openshift/origin/pkg/project/cache/cache.go:84
    I0801 19:47:58.136810       1 reflector.go:253] Listing and watching *api.ClusterResourceQuota from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:99
    I0801 19:47:58.200634       1 reflector.go:253] Listing and watching *api.LimitRange from github.com/openshift/origin/pkg/image/admission/imagestream_limits.go:29
    
    E0801 19:47:58.209483       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: the server could not find the requested resource
    
    I0801 19:47:58.220768       1 clusterquotamapping.go:105] Starting workers for quota mapping controller workers
    I0801 19:47:58.380599       1 ensure.go:86] Added endpoint-controller service accounts to the system:endpoint-controller cluster role: <nil>
    I0801 19:47:58.497603       1 ensure.go:86] Added gc-controller service accounts to the system:gc-controller cluster role: <nil>
    I0801 19:47:58.613683       1 ensure.go:86] Added hpa-controller service accounts to the system:hpa-controller cluster role: <nil>
    I0801 19:47:58.740108       1 ensure.go:86] Added job-controller service accounts to the system:job-controller cluster role: <nil>
    I0801 19:47:58.869913       1 ensure.go:86] Added namespace-controller service accounts to the system:namespace-controller cluster role: <nil>
    I0801 19:47:58.989405       1 ensure.go:86] Added pv-binder-controller service accounts to the system:pv-binder-controller cluster role: <nil>
    I0801 19:47:59.111291       1 ensure.go:86] Added pv-provisioner-controller service accounts to the system:pv-provisioner-controller cluster role: <nil>
    I0801 19:47:59.112701       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    
    E0801 19:47:59.122335       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: the server could not find the requested resource
    
    I0801 19:47:59.211783       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    
    E0801 19:47:59.220863       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: the server could not find the requested resource
    
    I0801 19:47:59.233038       1 ensure.go:86] Added pv-recycler-controller service accounts to the system:pv-recycler-controller cluster role: <nil>
    I0801 19:47:59.356231       1 ensure.go:86] Added replication-controller service accounts to the system:replication-controller cluster role: <nil>
    I0801 19:47:59.531186       1 ensure.go:86] Added service-load-balancer-controller service accounts to the system:service-load-balancer-controller cluster role: <nil>
    I0801 19:47:59.704235       1 ensure.go:86] Added service-serving-cert-controller service accounts to the system:service-serving-cert-controller cluster role: <nil>
    I0801 19:47:59.748624       1 net.go:21] Failure while checking listen on listen tcp4 10.3.0.10:53: bind: cannot assign requested address: %!v(MISSING)
    W0801 19:47:59.749305       1 run_components.go:202] Could not start DNS: listen tcp4 10.3.0.10:53: bind: cannot assign requested address
    I0801 19:47:59.749353       1 reflector.go:202] Starting reflector *api.Namespace (2m0s) from github.com/openshift/origin/pkg/project/auth/cache.go:186
    I0801 19:47:59.749434       1 reflector.go:253] Listing and watching *api.Namespace from github.com/openshift/origin/pkg/project/auth/cache.go:186
    I0801 19:47:59.750680       1 start_master.go:544] Controllers starting (*)
    I0801 19:47:59.753044       1 run_components.go:106] Skipped starting Service Account Token Manager, no private key specified
    I0801 19:47:59.754136       1 reflector.go:213] Starting reflector *api.Namespace (0) from pkg/controller/serviceaccount/serviceaccounts_controller.go:143
    I0801 19:47:59.754187       1 reflector.go:253] Listing and watching *api.Namespace from pkg/controller/serviceaccount/serviceaccounts_controller.go:143
    I0801 19:47:59.755974       1 reflector.go:213] Starting reflector *api.ServiceAccount (0) from pkg/controller/serviceaccount/serviceaccounts_controller.go:142
    I0801 19:47:59.756025       1 reflector.go:253] Listing and watching *api.ServiceAccount from pkg/controller/serviceaccount/serviceaccounts_controller.go:142
    I0801 19:47:59.757061       1 reflector.go:213] Starting reflector *api.Secret (0) from github.com/openshift/origin/pkg/serviceaccounts/controllers/deleted_dockercfg_secrets.go:73
    I0801 19:47:59.757112       1 reflector.go:253] Listing and watching *api.Secret from github.com/openshift/origin/pkg/serviceaccounts/controllers/deleted_dockercfg_secrets.go:73
    I0801 19:47:59.758983       1 reflector.go:213] Starting reflector *api.Secret (0) from github.com/openshift/origin/pkg/serviceaccounts/controllers/deleted_token_secrets.go:67
    I0801 19:47:59.759928       1 reflector.go:253] Listing and watching *api.Secret from github.com/openshift/origin/pkg/serviceaccounts/controllers/deleted_token_secrets.go:67
    I0801 19:47:59.759654       1 reflector.go:213] Starting reflector *api.Secret (0) from github.com/openshift/origin/pkg/serviceaccounts/controllers/docker_registry_service.go:134
    I0801 19:47:59.765686       1 reflector.go:253] Listing and watching *api.Secret from github.com/openshift/origin/pkg/serviceaccounts/controllers/docker_registry_service.go:134
    I0801 19:47:59.759533       1 reflector.go:213] Starting reflector *api.Service (0) from github.com/openshift/origin/pkg/serviceaccounts/controllers/docker_registry_service.go:133
    I0801 19:47:59.766034       1 reflector.go:253] Listing and watching *api.Service from github.com/openshift/origin/pkg/serviceaccounts/controllers/docker_registry_service.go:133
    I0801 19:47:59.840210       1 reflector.go:202] Starting reflector *api.Namespace (10m0s) from github.com/openshift/origin/pkg/security/controller/factory.go:40
    I0801 19:47:59.841558       1 reflector.go:213] Starting reflector *api.SecurityContextConstraints (0) from github.com/openshift/origin/pkg/security/admission/admission.go:51
    I0801 19:47:59.841992       1 reflector.go:253] Listing and watching *api.Namespace from github.com/openshift/origin/pkg/security/controller/factory.go:40
    I0801 19:47:59.844032       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    I0801 19:47:59.888067       1 create_dockercfg_secrets.go:119] Dockercfg secret controller initialized, starting.
    I0801 19:47:59.896019       1 reflector.go:213] Starting reflector *api.ServiceAccount (0) from github.com/openshift/origin/pkg/serviceaccounts/controllers/create_dockercfg_secrets.go:121
    I0801 19:47:59.896697       1 reflector.go:253] Listing and watching *api.ServiceAccount from github.com/openshift/origin/pkg/serviceaccounts/controllers/create_dockercfg_secrets.go:121
    I0801 19:47:59.964597       1 create_dockercfg_secrets.go:73] Adding service account default
    I0801 19:47:59.964640       1 create_dockercfg_secrets.go:73] Adding service account hpa-controller
    I0801 19:47:59.964656       1 create_dockercfg_secrets.go:73] Adding service account pv-binder-controller
    I0801 19:47:59.964664       1 create_dockercfg_secrets.go:73] Adding service account builder
    I0801 19:47:59.964670       1 create_dockercfg_secrets.go:73] Adding service account deployer
    I0801 19:47:59.964676       1 create_dockercfg_secrets.go:73] Adding service account build-controller
    I0801 19:47:59.964682       1 create_dockercfg_secrets.go:73] Adding service account deployment-controller
    I0801 19:47:59.964688       1 create_dockercfg_secrets.go:73] Adding service account builder
    I0801 19:47:59.964695       1 create_dockercfg_secrets.go:73] Adding service account builder
    I0801 19:47:59.964702       1 create_dockercfg_secrets.go:73] Adding service account deployer
    I0801 19:47:59.964708       1 create_dockercfg_secrets.go:73] Adding service account endpoint-controller
    I0801 19:47:59.964714       1 create_dockercfg_secrets.go:73] Adding service account job-controller
    I0801 19:47:59.964720       1 create_dockercfg_secrets.go:73] Adding service account namespace-controller
    I0801 19:47:59.964728       1 create_dockercfg_secrets.go:73] Adding service account pv-provisioner-controller
    I0801 19:47:59.964735       1 create_dockercfg_secrets.go:73] Adding service account default
    I0801 19:47:59.964741       1 create_dockercfg_secrets.go:73] Adding service account daemonset-controller
    I0801 19:47:59.964747       1 create_dockercfg_secrets.go:73] Adding service account replication-controller
    I0801 19:47:59.964752       1 create_dockercfg_secrets.go:73] Adding service account service-serving-cert-controller
    I0801 19:47:59.964758       1 create_dockercfg_secrets.go:73] Adding service account deployer
    I0801 19:47:59.964765       1 create_dockercfg_secrets.go:73] Adding service account default
    I0801 19:47:59.964771       1 create_dockercfg_secrets.go:73] Adding service account default
    I0801 19:47:59.964777       1 create_dockercfg_secrets.go:73] Adding service account default
    I0801 19:47:59.964783       1 create_dockercfg_secrets.go:73] Adding service account deployer
    I0801 19:47:59.964789       1 create_dockercfg_secrets.go:73] Adding service account builder
    I0801 19:47:59.964810       1 create_dockercfg_secrets.go:73] Adding service account gc-controller
    I0801 19:47:59.964816       1 create_dockercfg_secrets.go:73] Adding service account pv-recycler-controller
    I0801 19:47:59.964824       1 create_dockercfg_secrets.go:73] Adding service account service-load-balancer-controller
    I0801 19:47:59.964831       1 create_dockercfg_secrets.go:73] Adding service account deployer
    I0801 19:47:59.964837       1 create_dockercfg_secrets.go:73] Adding service account builder
    
    E0801 19:47:59.965950       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: the server could not find the requested resource
    
    I0801 19:47:59.967713       1 reflector.go:213] Starting reflector *api.Build (2m0s) from github.com/openshift/origin/pkg/build/controller/factory/factory.go:77
    I0801 19:47:59.967895       1 reflector.go:213] Starting reflector *api.Build (5m0s) from github.com/openshift/origin/pkg/build/controller/factory/factory.go:130
    I0801 19:47:59.967914       1 reflector.go:213] Starting reflector *api.Build (2m0s) from github.com/openshift/origin/pkg/build/controller/factory/factory.go:191
    I0801 19:47:59.967945       1 reflector.go:213] Starting reflector *api.Pod (2m0s) from github.com/openshift/origin/pkg/build/controller/factory/factory.go:194
    I0801 19:47:59.967962       1 reflector.go:213] Starting reflector *api.Pod (5m0s) from github.com/openshift/origin/pkg/build/controller/factory/factory.go:246
    I0801 19:47:59.967978       1 reflector.go:213] Starting reflector *api.BuildConfig (2m0s) from github.com/openshift/origin/pkg/build/controller/factory/factory.go:324
    I0801 19:47:59.968036       1 reflector.go:213] Starting reflector *api.ImageStream (2m0s) from github.com/openshift/origin/pkg/build/controller/factory/factory.go:285
    I0801 19:47:59.968041       1 reflector.go:213] Starting reflector *api.BuildConfig (2m0s) from github.com/openshift/origin/pkg/build/controller/factory/factory.go:288
    I0801 19:47:59.968231       1 reflector.go:253] Listing and watching *api.Build from github.com/openshift/origin/pkg/build/controller/factory/factory.go:77
    I0801 19:47:59.968733       1 reflector.go:253] Listing and watching *api.Build from github.com/openshift/origin/pkg/build/controller/factory/factory.go:130
    I0801 19:47:59.968790       1 factory.go:480] Checking for deleted builds
    I0801 19:47:59.969104       1 reflector.go:253] Listing and watching *api.Build from github.com/openshift/origin/pkg/build/controller/factory/factory.go:191
    I0801 19:47:59.969324       1 reflector.go:253] Listing and watching *api.Pod from github.com/openshift/origin/pkg/build/controller/factory/factory.go:194
    I0801 19:47:59.969649       1 reflector.go:253] Listing and watching *api.Pod from github.com/openshift/origin/pkg/build/controller/factory/factory.go:246
    I0801 19:47:59.969709       1 factory.go:565] Checking for deleted build pods
    I0801 19:47:59.970051       1 reflector.go:253] Listing and watching *api.BuildConfig from github.com/openshift/origin/pkg/build/controller/factory/factory.go:324
    I0801 19:47:59.970387       1 reflector.go:253] Listing and watching *api.ImageStream from github.com/openshift/origin/pkg/build/controller/factory/factory.go:285
    I0801 19:47:59.970703       1 reflector.go:253] Listing and watching *api.BuildConfig from github.com/openshift/origin/pkg/build/controller/factory/factory.go:288
    I0801 19:48:00.139883       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    I0801 19:48:00.201257       1 reflector.go:202] Starting reflector *api.ImageStream (2m0s) from github.com/openshift/origin/pkg/deploy/controller/imagechange/factory.go:37
    I0801 19:48:00.201948       1 reflector.go:202] Starting reflector *api.DeploymentConfig (2m0s) from github.com/openshift/origin/pkg/deploy/controller/imagechange/factory.go:48
    I0801 19:48:00.202337       1 reflector.go:202] Starting reflector *api.ImageStream (10m0s) from github.com/openshift/origin/pkg/image/controller/factory.go:40
    I0801 19:48:00.202383       1 reflector.go:202] Starting reflector *api.Namespace (1m0s) from github.com/openshift/origin/pkg/project/controller/factory.go:37
    
    E0801 19:48:00.205275       1 util.go:45] Metric for replenishment_controller already registered
    E0801 19:48:00.207078       1 util.go:45] Metric for replenishment_controller already registered
    E0801 19:48:00.207112       1 util.go:45] Metric for replenishment_controller already registered
    E0801 19:48:00.207122       1 util.go:45] Metric for replenishment_controller already registered
    E0801 19:48:00.207123       1 util.go:45] Metric for replenishment_controller already registered
    E0801 19:48:00.207240       1 util.go:45] Metric for replenishment_controller already registered
    E0801 19:48:00.209158       1 util.go:45] Metric for replenishment_controller already registered
    E0801 19:48:00.209668       1 util.go:45] Metric for replenishment_controller already registered
    E0801 19:48:00.209686       1 util.go:45] Metric for replenishment_controller already registered
    E0801 19:48:00.209711       1 util.go:45] Metric for replenishment_controller already registered
    E0801 19:48:00.209720       1 util.go:45] Metric for replenishment_controller already registered
    E0801 19:48:00.209727       1 util.go:45] Metric for replenishment_controller already registered
    E0801 19:48:00.209734       1 util.go:45] Metric for replenishment_controller already registered
    
    I0801 19:48:00.229940       1 factory.go:91] Waiting for the rc and pod caches to sync before starting the deployment controller workers
    I0801 19:48:00.230685       1 factory.go:101] Waiting for the dc, rc, and pod caches to sync before starting the deployment config controller workers
    I0801 19:48:00.230732       1 factory.go:82] Waiting for the deployment config cache to sync before starting the trigger controller workers
    I0801 19:48:00.230751       1 reflector.go:253] Listing and watching *api.ImageStream from github.com/openshift/origin/pkg/deploy/controller/imagechange/factory.go:37
    I0801 19:48:00.230901       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    I0801 19:48:00.231212       1 reflector.go:253] Listing and watching *api.DeploymentConfig from github.com/openshift/origin/pkg/deploy/controller/imagechange/factory.go:48
    I0801 19:48:00.231955       1 reflector.go:253] Listing and watching *api.ImageStream from github.com/openshift/origin/pkg/image/controller/factory.go:40
    I0801 19:48:00.235251       1 scheduler.go:74] DEBUG: scheduler: queue (0):
     []controller.bucket{controller.bucket{}, controller.bucket{}, controller.bucket{}, controller.bucket{}, controller.bucket{}}
    I0801 19:48:00.235304       1 scheduler.go:79] DEBUG: scheduler: position: 1 5
    I0801 19:48:00.235310       1 scheduler.go:56] DEBUG: scheduler: waiting for limit
    I0801 19:48:00.235336       1 reflector.go:253] Listing and watching *api.Namespace from github.com/openshift/origin/pkg/project/controller/factory.go:37
    I0801 19:48:00.235802       1 resource_quota_controller.go:154] Resource quota controller queued all resource quota for full calculation of usage
    I0801 19:48:00.236230       1 reflector.go:213] Starting reflector *api.ResourceQuota (5m0s) from pkg/controller/resourcequota/resource_quota_controller.go:230
    I0801 19:48:00.236366       1 reflector.go:253] Listing and watching *api.ResourceQuota from pkg/controller/resourcequota/resource_quota_controller.go:230
    I0801 19:48:00.237113       1 reflector.go:213] Starting reflector *api.Service (12h0m0s) from pkg/controller/resourcequota/resource_quota_controller.go:233
    I0801 19:48:00.237144       1 reflector.go:253] Listing and watching *api.Service from pkg/controller/resourcequota/resource_quota_controller.go:233
    I0801 19:48:00.237918       1 reflector.go:213] Starting reflector *api.ReplicationController (12h0m0s) from pkg/controller/resourcequota/resource_quota_controller.go:233
    I0801 19:48:00.237964       1 reflector.go:253] Listing and watching *api.ReplicationController from pkg/controller/resourcequota/resource_quota_controller.go:233
    I0801 19:48:00.239111       1 reflector.go:213] Starting reflector *api.PersistentVolumeClaim (12h0m0s) from pkg/controller/resourcequota/resource_quota_controller.go:233
    I0801 19:48:00.239188       1 reflector.go:253] Listing and watching *api.PersistentVolumeClaim from pkg/controller/resourcequota/resource_quota_controller.go:233
    I0801 19:48:00.240093       1 reflector.go:213] Starting reflector *api.Secret (12h0m0s) from pkg/controller/resourcequota/resource_quota_controller.go:233
    I0801 19:48:00.240139       1 reflector.go:253] Listing and watching *api.Secret from pkg/controller/resourcequota/resource_quota_controller.go:233
    I0801 19:48:00.240819       1 reflector.go:213] Starting reflector *api.ConfigMap (12h0m0s) from pkg/controller/resourcequota/resource_quota_controller.go:233
    I0801 19:48:00.240873       1 reflector.go:253] Listing and watching *api.ConfigMap from pkg/controller/resourcequota/resource_quota_controller.go:233
    I0801 19:48:00.241605       1 reflector.go:213] Starting reflector *api.ImageStream (12h0m0s) from pkg/controller/resourcequota/resource_quota_controller.go:233
    I0801 19:48:00.241653       1 reflector.go:253] Listing and watching *api.ImageStream from pkg/controller/resourcequota/resource_quota_controller.go:233
    I0801 19:48:00.242801       1 reflector.go:213] Starting reflector *api.Service (12h0m0s) from github.com/openshift/origin/pkg/quota/controller/clusterquotareconciliation/reconcilation_controller.go:119
    I0801 19:48:00.242869       1 reflector.go:253] Listing and watching *api.Service from github.com/openshift/origin/pkg/quota/controller/clusterquotareconciliation/reconcilation_controller.go:119
    I0801 19:48:00.243654       1 reflector.go:213] Starting reflector *api.ReplicationController (12h0m0s) from github.com/openshift/origin/pkg/quota/controller/clusterquotareconciliation/reconcilation_controller.go:119
    I0801 19:48:00.243697       1 reflector.go:253] Listing and watching *api.ReplicationController from github.com/openshift/origin/pkg/quota/controller/clusterquotareconciliation/reconcilation_controller.go:119
    I0801 19:48:00.244360       1 reflector.go:213] Starting reflector *api.PersistentVolumeClaim (12h0m0s) from github.com/openshift/origin/pkg/quota/controller/clusterquotareconciliation/reconcilation_controller.go:119
    I0801 19:48:00.244542       1 reflector.go:253] Listing and watching *api.PersistentVolumeClaim from github.com/openshift/origin/pkg/quota/controller/clusterquotareconciliation/reconcilation_controller.go:119
    I0801 19:48:00.245270       1 reflector.go:213] Starting reflector *api.Secret (12h0m0s) from github.com/openshift/origin/pkg/quota/controller/clusterquotareconciliation/reconcilation_controller.go:119
    I0801 19:48:00.245312       1 reflector.go:253] Listing and watching *api.Secret from github.com/openshift/origin/pkg/quota/controller/clusterquotareconciliation/reconcilation_controller.go:119
    I0801 19:48:00.246034       1 reflector.go:213] Starting reflector *api.ConfigMap (12h0m0s) from github.com/openshift/origin/pkg/quota/controller/clusterquotareconciliation/reconcilation_controller.go:119
    I0801 19:48:00.246101       1 reflector.go:253] Listing and watching *api.ConfigMap from github.com/openshift/origin/pkg/quota/controller/clusterquotareconciliation/reconcilation_controller.go:119
    I0801 19:48:00.246568       1 reflector.go:213] Starting reflector *api.ImageStream (12h0m0s) from github.com/openshift/origin/pkg/quota/controller/clusterquotareconciliation/reconcilation_controller.go:119
    I0801 19:48:00.246568       1 reflector.go:253] Listing and watching *api.ImageStream from github.com/openshift/origin/pkg/quota/controller/clusterquotareconciliation/reconcilation_controller.go:119
    I0801 19:48:00.262748       1 scheduler.go:74] DEBUG: scheduler: queue (1):
     []controller.bucket{controller.bucket{}, controller.bucket{}, controller.bucket{}, controller.bucket{}, controller.bucket{}}
    I0801 19:48:00.262785       1 scheduler.go:79] DEBUG: scheduler: position: 2 5
    I0801 19:48:00.262791       1 scheduler.go:56] DEBUG: scheduler: waiting for limit
    
    E0801 19:48:00.305233       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: the server could not find the requested resource
    
    I0801 19:48:00.362976       1 factory.go:91] Waiting for the rc and pod caches to sync before starting the deployment controller workers
    I0801 19:48:00.362989       1 factory.go:101] Waiting for the dc, rc, and pod caches to sync before starting the deployment config controller workers
    I0801 19:48:00.362998       1 factory.go:82] Waiting for the deployment config cache to sync before starting the trigger controller workers
    I0801 19:48:00.516770       1 factory.go:82] Waiting for the deployment config cache to sync before starting the trigger controller workers
    I0801 19:48:00.537709       1 factory.go:91] Waiting for the rc and pod caches to sync before starting the deployment controller workers
    I0801 19:48:00.537760       1 factory.go:101] Waiting for the dc, rc, and pod caches to sync before starting the deployment config controller workers
    I0801 19:48:00.618168       1 factory.go:82] Waiting for the deployment config cache to sync before starting the trigger controller workers
    I0801 19:48:00.650143       1 factory.go:91] Waiting for the rc and pod caches to sync before starting the deployment controller workers
    I0801 19:48:00.650771       1 factory.go:101] Waiting for the dc, rc, and pod caches to sync before starting the deployment config controller workers
    I0801 19:48:00.659589       1 start_master.go:662] Started Origin Controllers
    I0801 19:48:00.660399       1 reflector.go:213] Starting reflector *api.ImageStream (10m0s) from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:87
    I0801 19:48:00.660950       1 reflector.go:253] Listing and watching *api.ImageStream from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:87
    
    E0801 19:48:00.661834       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: the server could not find the requested resource
    
    I0801 19:48:00.662575       1 reflector.go:213] Starting reflector *api.ReplicationController (10m0s) from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:87
    I0801 19:48:00.663034       1 reflector.go:253] Listing and watching *api.ReplicationController from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:87
    I0801 19:48:00.663839       1 reflector.go:213] Starting reflector *api.Pod (10m0s) from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:87
    I0801 19:48:00.664666       1 reflector.go:253] Listing and watching *api.Pod from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:87
    I0801 19:48:00.665711       1 reflector.go:213] Starting reflector *api.DeploymentConfig (10m0s) from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:87
    I0801 19:48:00.669581       1 reflector.go:253] Listing and watching *api.DeploymentConfig from github.com/openshift/origin/pkg/controller/shared/shared_informer.go:87
    I0801 19:48:00.675064       1 factory.go:147] Replication controller "openshift" added.
    I0801 19:48:00.675113       1 factory.go:147] Replication controller "kube-dns-v15" added.
    I0801 19:48:00.675138       1 factory.go:147] Replication controller "kubernetes-dashboard-v1.1.0" added.
    I0801 19:48:00.675146       1 factory.go:147] Replication controller "etcd-v3-single" added.
    I0801 19:48:00.675155       1 factory.go:147] Replication controller "gogs" added.
    I0801 19:48:00.976706       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    
    E0801 19:48:00.979307       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: the server could not find the requested resource
    
    I0801 19:48:01.307634       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    
    E0801 19:48:01.310712       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: the server could not find the requested resource
    
    I0801 19:48:01.690826       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    
    E0801 19:48:01.694921       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: the server could not find the requested resource
    
    I0801 19:48:01.979674       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    
    E0801 19:48:01.982170       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: the server could not find the requested resource
    
    I0801 19:48:02.312212       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    
    E0801 19:48:02.316490       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: the server could not find the requested resource
    
    I0801 19:48:02.702787       1 reflector.go:253] Listing and watching *api.SecurityContextConstraints from github.com/openshift/origin/pkg/security/admission/admission.go:51
    
    E0801 19:48:02.708134       1 reflector.go:216] github.com/openshift/origin/pkg/security/admission/admission.go:51: Failed to list *api.SecurityContextConstraints: the server could not find the requested resource

<repeated error>