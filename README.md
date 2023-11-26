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
  - name: rabbitmq-cluster-operator
    repository: bitnami
```


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
$ docker run -it --rm -v ${PWD}/config.yaml:/config/charts.yaml germis/helm-util charts show -d
```


## Development

Dependencies that were added to the module

```bash
$ go get github.com/spf13/cobra
$ go get github.com/sirupsen/logrus
$ go get gopkg.in/yaml.v3
```