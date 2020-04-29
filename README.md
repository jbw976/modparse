# modparse

This `modparse` tool downloads and parses the `go.mod` files for the given set
of remote repositories, then prints out the unique set (no duplicates) of all
required dependencies found. It is expected that all **input** repositories are
hosted on `github.com`, but dependencies that are parsed and discovered by this tool
can be hosted anywhere.

### Help output

```console
> modparse
expecting at least 1 arg: ./modparse [repos...]
[repo] is in [org/repo] format, and expected to be on github.com, e.g. crossplane/crossplane
```

### Simple usage

```console
> modparse jbw976/modparse
downloading 'https://raw.githubusercontent.com/jbw976/modparse/master/go.mod' to '/var/folders/ry/8ngky47n1r38lkl65hjfclmr0000gn/T/modparse-*068997027/jbw976-modparse.mod'
found 1 dependencies
github.com/sirkon/goproxy
```

### Complex usage

```console
> modparse crossplane/crossplane crossplane/crossplane-runtime crossplane/tbs
downloading 'https://raw.githubusercontent.com/crossplane/crossplane/master/go.mod' to '/var/folders/ry/8ngky47n1r38lkl65hjfclmr0000gn/T/modparse-*448903559/crossplane-crossplane.mod'
downloading 'https://raw.githubusercontent.com/crossplane/crossplane-runtime/master/go.mod' to '/var/folders/ry/8ngky47n1r38lkl65hjfclmr0000gn/T/modparse-*448903559/crossplane-crossplane-runtime.mod'
downloading 'https://raw.githubusercontent.com/crossplane/tbs/master/go.mod' to '/var/folders/ry/8ngky47n1r38lkl65hjfclmr0000gn/T/modparse-*448903559/crossplane-tbs.mod'
[WARN]: not found, skipping
found 32 dependencies
github.com/Azure/go-autorest/autorest
github.com/Azure/go-autorest/autorest/adal
github.com/crossplane/crossplane-runtime
github.com/crossplane/crossplane-tools
github.com/crossplane/oam-kubernetes-runtime
github.com/docker/distribution
github.com/evanphx/json-patch
github.com/ghodss/yaml
github.com/go-logr/logr
github.com/go-logr/zapr
github.com/golang/groupcache
github.com/google/go-cmp
github.com/google/uuid
github.com/gophercloud/gophercloud
github.com/hashicorp/go-getter
github.com/hashicorp/golang-lru
github.com/imdario/mergo
github.com/onsi/gomega
github.com/opencontainers/go-digest
github.com/pkg/errors
github.com/prometheus/client_golang
github.com/spf13/afero
golang.org/x/net
gopkg.in/alecthomas/kingpin.v2
k8s.io/api
k8s.io/apiextensions-apiserver
k8s.io/apimachinery
k8s.io/client-go
k8s.io/utils
sigs.k8s.io/controller-runtime
sigs.k8s.io/controller-tools
sigs.k8s.io/yaml
```
