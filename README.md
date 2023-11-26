## Description

Helm wrapper utility to add some useful functionality to helm

* find latest versions of helm charts

## Usage

```bash
$ docker run -it --rm germis/helm-util version
0.0.1
```

A default config file is provided inside the image


### Config

The docker image expects a `YAML` config file (by default in `/config/charts.yaml`). A default config file is provided. The
config file should have the following structure

```yaml
repositories:
  - name: autoscaler
    url: https://kubernetes.github.io/autoscaler
  - name: bitnami
    url: https://charts.bitnami.com/bitnami

charts:
  - name: cluster-autoscaler
    repository: autoscaler
    aliases:
      - alternate-name-1
      - alternate-name-2    
  - name: rabbitmq-cluster-operator
    repository: bitnami
```

In the `aliases` section is optional. You can provide alternative names for the chart as it is known on your clusters (ie. you changed the name during installation). This information is used when charts are listed from a live cluster so that we can match them

You can show what config file is currently loaded by using the following command


```bash
$ docker run -it --rm germis/helm-util config show
```

You can override this config file either by mounting a local file over it (ie. `/config/charts.yaml`) 

```bash
$ docker run -it --rm -v ${PWD}/config.yaml:/config/charts.yaml germis/helm-util config show
```

or by mounting it somewhere else and referencing it using the `-c` flag during runtime

```bash
$ docker run -it --rm -v ${PWD}/config.yaml:/some/path/to/config.yaml germis/helm-util -c /some/path/to/config.yaml config show
```

### Charts

Actions can be performed against helm charts

* find the latest version of all charts defined in the config file
* find all helm charts installed on a cluster and check whether a newer version is available

```bash
$ docker run -it --rm \
  -v ${PWD}/config.yaml:/config/charts.yaml \
  germis/helm-util charts versions -d
```
To be able to list installed helm charts we need to pass a valid `kubeconfig` file to the docker instance and point the `$KUBECONFIG` environment variable to it

```bash
$ docker run -it --rm \
  -e KUBECONFIG=/config/kubeconfig \
  -v ${HOME}/.kube/kubeconfig:/config/kubeconfig \
  -v ${PWD}/config.yaml:/config/charts.yaml \
  germis/helm-util charts versions --live -d
```

## Development

Dependencies that were added to the module

```bash
$ go get github.com/spf13/cobra
$ go get github.com/sirupsen/logrus
$ go get gopkg.in/yaml.v3
$ go get github.com/google/uuid
```