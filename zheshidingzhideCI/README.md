Openshift Origin in cluster CI
================================




Deploy into external kubernetes
--------------------------------

Environment

    [vagrant@localhost origin]$ echo $KUBECONFIG
    /data/src/github.com/openshift/origin/etc/kubeconfig

### v1.3.0-alpha.2 ( Aug. 1 )

Create Openshift Origin namespace

    [vagrant@localhost origin]$ _output/local/bin/linux/amd64/oc create -f etc/openshift/openshift-origin-namespace.yaml 
    namespace "openshift-origin" created
    [vagrant@localhost origin]$ _output/local/bin/linux/amd64/oc get namespaces
    NAME               STATUS    AGE
    default            Active    7h
    kube-system        Active    7h
    openshift-origin   Active    5s

Create Openshift Origin service

    [vagrant@localhost origin]$ _output/local/bin/linux/amd64/oc create -f etc/openshift/openshift-service.yaml 
    service "openshift" created
    [vagrant@localhost origin]$ _output/local/bin/linux/amd64/oc get service --namespace=openshift-origin
    NAME        CLUSTER-IP   EXTERNAL-IP   PORT(S)    AGE
    openshift   10.3.0.53    <nodes>       8443/TCP   30s
    [{openshift TCP 8443 {0 8443 } 30443}][vagrant@localhost origin]$ 
    [vagrant@localhost origin]$ _output/local/bin/linux/amd64/oc get service/openshift --namespace=openshift-origin -o jsonpath={.spec.ports[0].nodePort}
    30443

Using Etcd3

    [vagrant@localhost orig_output/local/bin/linux/amd64/oc get service --namespace=default
    NAME             CLUSTER-IP   EXTERNAL-IP   PORT(S)           AGE
    etcd-v3-single   10.3.0.189   <nodes>       2379/TCP          1h
    gogs             10.3.0.46    <nodes>       3000/TCP,22/TCP   1h
    kubernetes       10.3.0.1     <none>        443/TCP           7h

Create master configuration

    [vagrant@localhost origin]$ openshift start master --kubeconfig=etc/kubeconfig --write-config=openshift.local.config/master --dns=tcp://10.3.0.10:53 --host-subnet-length=7 --network-cidr=172.17.0.1/22 --portal-net=10.3.0.0/24 --etcd=http://10.3.0.189:2379 --master=https://10.3.0.53:8443 --public-master=https://172.17.4.50:30443
    Generated new key pair as openshift.local.config/master/serviceaccounts.public.key and openshift.local.config/master/serviceaccounts.private.key
    Wrote master config to: openshift.local.config/master/master-config.yaml

    [vagrant@localhost origin]$ cp etc/kubernetes/cacerts/server.key openshift.local.config/master/kube-apiserver.key

Create Replication Controller of master

    [vagrant@localhost origin]$ _output/local/bin/linux/amd64/oc create -f etc/openshift/openshift-controller-dev.yaml 
    replicationcontroller "openshift" created
    [vagrant@localhost origin]$ _output/local/bin/linux/amd64/oc get pods --namespace=openshift-origin
    NAME              READY     STATUS    RESTARTS   AGE
    openshift-fnehe   0/1       Error     2          43s
    [vagrant@localhost origin]$ docker ps -l
    CONTAINER ID        IMAGE                           COMMAND                  CREATED             STATUS                        PORTS               NAMES
    cadae9412c9c        tangfeixiong/openshift-origin   "/usr/bin/openshift s"   16 seconds ago      Exited (255) 12 seconds ago                       k8s_origin.5c58aaae_openshift-fnehe_openshift-origin_47b5846a-5819-11e6-ab10-0800274654e7_355215ba
    [vagrant@localhost origin]$ docker logs cadae9412c9c
    W0801 18:55:38.477653       1 start_master.go:272] assetConfig.loggingPublicURL: Invalid value: "": required to view aggregated container logs in the console
    W0801 18:55:38.480418       1 start_master.go:272] assetConfig.metricsPublicURL: Invalid value: "": required to view cluster metrics in the console
    W0801 18:55:38.480451       1 start_master.go:272] serviceAccountConfig.publicKeyFiles: Invalid value: "": no service account tokens will be accepted by the API, which will prevent builds and deployments from working
    Invalid MasterConfig /data/src/github.com/openshift/origin/openshift.local.config/master/master-config.yaml
      masterClients.externalKubernetesKubeConfig: Invalid value: "/data/src/github.com/openshift/origin/etc/kubeconfig": could not read file

    [vagrant@localhost origin]$ vim etc/openshift/openshift-controller-dev.yaml 
    [vagrant@localhost origin]$ _output/local/bin/linux/amd64/oc create -f etc/openshift/openshift-controller-dev.yaml 
    replicationcontroller "openshift" created
    [vagrant@localhost origin]$ _output/local/bin/linux/amd64/oc get pods --namespace=openshift-origin
    NAME              READY     STATUS    RESTARTS   AGE
    openshift-mplpp   1/1       Running   0          23s

Validation

    [vagrant@localhost origin]$ _output/local/bin/linux/amd64/oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig get projects
    NAME               DISPLAY NAME   STATUS
    kube-system                       Active
    openshift                         Active
    openshift-infra                   Active
    openshift-origin                  Active
    default                           Active
    [vagrant@localhost origin]$ _output/local/bin/linux/amd64/oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig get sa
    NAME       SECRETS   AGE
    builder    2         6m
    default    2         7h
    deployer   2         6m
    [vagrant@localhost origin]$ _output/local/bin/linux/amd64/oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig get secrets
    NAME                       TYPE                                  DATA      AGE
    builder-dockercfg-rzdwd    kubernetes.io/dockercfg               1         6m
    builder-token-6diro        kubernetes.io/service-account-token   3         6m
    builder-token-nx62t        kubernetes.io/service-account-token   3         6m
    default-dockercfg-vf9kz    kubernetes.io/dockercfg               1         6m
    default-token-awl1h        kubernetes.io/service-account-token   3         6m
    default-token-xriex        kubernetes.io/service-account-token   3         7h
    deployer-dockercfg-e4drv   kubernetes.io/dockercfg               1         6m
    deployer-token-f4auc       kubernetes.io/service-account-token   3         6m
    deployer-token-rg9r6       kubernetes.io/service-account-token   3         6m

### v1.3.0-alpha.0 ( May. 12 )

Create Openshift Origin namespace

    [vagrant@localhost origin]$ oc --config=openshift.local.config/master/kubeconfig create -f openshift-origin-namespace.yaml 
    namespace "openshift-origin" created
    
    [vagrant@localhost origin]$ kubectl --kubeconfig=etc/kubernetes/kubeconfig get ns
    NAME               STATUS    AGE
    default            Active    29m
    kube-system        Active    26m
    openshift-origin   Active    14s

Create Openshift Origin service

    [vagrant@localhost origin]$ oc --config=openshift.local.config/master/kubeconfig create -f openshift-service.yaml --namespace=openshift-origin
    service "openshift" created
    
    [vagrant@localhost origin]$ oc --config=openshift.local.config/master/kubeconfig --namespace=openshift-origin get svc
    NAME        CLUSTER-IP   EXTERNAL-IP   PORT(S)    AGE
    openshift   10.3.0.179                 8443/TCP   29s
    
    [vagrant@localhost origin]$ oc --config=openshift.local.config/master/kubeconfig --namespace=openshift-origin get svc openshift -o jsonpath={.spec.ports[0].nodePort};echo
    30448

