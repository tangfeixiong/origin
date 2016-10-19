Tutorial
=========

Target Platform - Aug. 1
-------------------------

* Project and User

Create and Get by System Admin

    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig new-project tangfeixiong
    Now using project "tangfeixiong" on server "https://10.3.0.53:8443".
    
    You can add applications to this project with the 'new-app' command. For example, try:
    
        $ oc new-app centos/ruby-22-centos7~https://github.com/openshift/ruby-ex.git
    
    to build a new example application in Ruby.
    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig get projects
    NAME               DISPLAY NAME   STATUS
    tangfeixiong                      Active
    default                           Active
    kube-system                       Active
    openshift                         Active
    openshift-infra                   Active
    openshift-origin                  Active
    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig get projects/tangfeixiong -o yaml
    apiVersion: v1
    kind: Project
    metadata:
      annotations:
        openshift.io/description: ""
        openshift.io/display-name: ""
        openshift.io/requester: system:admin
        openshift.io/sa.scc.mcs: s0:c7,c4
        openshift.io/sa.scc.supplemental-groups: 1000050000/10000
        openshift.io/sa.scc.uid-range: 1000050000/10000
      creationTimestamp: 2016-08-03T02:46:13Z
      name: tangfeixiong
      resourceVersion: "9673"
      selfLink: /oapi/v1/projects/tangfeixiong
      uid: 714e1d17-5924-11e6-ab10-0800274654e7
    spec:
      finalizers:
      - openshift.io/origin
      - kubernetes
    status:
      phase: Active

* About login

Login with System Admin

    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kube --username=system:admin --namespace=default
    You have access to the following projects and can switch between them with 'oc project <projectname>':
    
      * default (current)
      * kube-system
      * openshift
      * openshift-infra
      * openshift-origin
      * tangfeixiong
    
    Using project "default".
    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig login --username=system:admin --namespace=tangfeixiong
    Logged into "https://10.3.0.53:8443" as "system:admin" using existing credentials.
    
    You have access to the following projects and can switch between them with 'oc project <projectname>':
    
      * default
      * kube-system
      * openshift
      * openshift-infra
      * openshift-origin
      * tangfeixiong (current)
    
    Using project "tangfeixiong".

Login and operate project with user

    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig login
    Authentication required for https://10.3.0.53:8443 (openshift)
    Username: tangfeixiong
    Password: 
    Login successful.
    
    You don't have any projects. You can try to create a new project, by running
    
        $ oc new-project <projectname>
    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig new-project tangfeixiong
    Error from server: project "tangfeixiong" already exists
    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig new-project gogogo
    Now using project "gogogo" on server "https://10.3.0.53:8443".
    
    You can add applications to this project with the 'new-app' command. For example, try:
    
        $ oc new-app centos/ruby-22-centos7~https://github.com/openshift/ruby-ex.git
    
    to build a new example application in Ruby.
    
    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig get projects
    NAME      DISPLAY NAME   STATUS
    gogogo                   Active
    
    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig login --username=system:admin --namespace=gogogo
    You have access to the following projects and can switch between them with 'oc project <projectname>':
    
      * default
      * gogogo (current)
      * kube-system
      * openshift
      * openshift-infra
      * openshift-origin
      * tangfeixiong
    
    Using project "gogogo".
    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig get projects
    NAME               DISPLAY NAME   STATUS
    default                           Active
    gogogo                            Active
    kube-system                       Active
    openshift                         Active
    openshift-infra                   Active
    openshift-origin                  Active
    tangfeixiong                      Active
    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig get projects/gogogo -o yaml
    apiVersion: v1
    kind: Project
    metadata:
      annotations:
        openshift.io/description: ""
        openshift.io/display-name: ""
        openshift.io/requester: tangfeixiong
        openshift.io/sa.scc.mcs: s0:c8,c2
        openshift.io/sa.scc.supplemental-groups: 1000060000/10000
        openshift.io/sa.scc.uid-range: 1000060000/10000
      creationTimestamp: 2016-08-03T02:58:35Z
      name: gogogo
      resourceVersion: "9769"
      selfLink: /oapi/v1/projects/gogogo
      uid: 2b943fcb-5926-11e6-ab10-0800274654e7
    spec:
      finalizers:
      - openshift.io/origin
      - kubernetes
    status:
      phase: Active

Show users by System:Admin and user self

    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig get users
    NAME           UID                                    FULL NAME   IDENTITIES
    tangfeixiong   f2c85765-5923-11e6-aafd-860e000d794a               anypassword:tangfeixiong
    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig get users/tangfeixiong -o yaml
    apiVersion: v1
    groups: null
    identities:
    - anypassword:tangfeixiong
    kind: User
    metadata:
      creationTimestamp: 2016-08-03T02:42:41Z
      name: tangfeixiong
      resourceVersion: "116"
      selfLink: /oapi/v1/users/tangfeixiong
      uid: f2c85765-5923-11e6-aafd-860e000d794a
    
    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig login --username=tangfeixiong --namespace=gogogo
    Using project "gogogo".
    [vagrant@localhost go-to-cloud-1]$ oc --config=/data/src/github.com/openshift/origin/openshift.local.config/master/admin.kubeconfig get users
    Error from server: User "tangfeixiong" cannot list all users in the cluster
