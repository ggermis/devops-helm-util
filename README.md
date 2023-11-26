## Description

Helm wrapper utility to add some useful functionality to helm

* find latest versions of helm charts

## Usage

```bash
$ docker run -it --rm germis/helm-util version
0.0.1
```

```bash
$ docker run -it --rm germis/helm-util config show
```

```bash
$ docker run -it --rm germis/helm-util charts show -d
```


## Development

```bash
$ go get github.com/spf13/cobra
$ go get github.com/sirupsen/logrus
$ go get gopkg.in/yaml.v3
```