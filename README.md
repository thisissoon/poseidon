# Compass

Compass is a release management tool (specific to SOON_\_) that handles managing namerd finagle delegation tables, mapping logical names to Kubernetets services deployed via Helm.

## Goals

* Manage HTTP/1.1 and gRPC services - these live in specific `namerd` namespaces
* Easy management of mapping service logical names (Host name for example) to Kubernetets services
* A UI appropriate for Humans to use (potentially interactive)
* Ability to run as part of automation for CI/CD
* gRPC Client / Server model - allowing for other UI's to be created

### CLI

Below are some initial CLI thoughts:

#### Managing Apps

To manage an application we would create on using the `upsert` command:

```
compass manage --logical-name frontend --namespace http --description "Frontend application"
```

This would take 3 arguments:

* `--logical-name`: A unique logical name for the application
* `--namespace`: The `namerd` namespace the delegation table entry should be managed for this app
* `--description`: An optional app description

#### Routing to a Kubernetets service

To update a services delegation table entry to map the logical name to a Kubernetes service
we would use the `route` command. This would instantly update the the applications delegation
table entry to route requests for the logical name to a Kubernetes service.

The labels on Kubernetes services are key to aid in this management, a `logical-name` label
must exist on the service, for version support a `version` label must also exist on the service.

``` yaml
apiVersion: v1
kind: Service
metadata:
  name: frontend-v1-2-3
  labels:
    app: helm-name
    logical-name: frontend
    version: v1.2.3
spec:
  type: NodePort
  selector:
    app: helm-name
    version: v1.2.3
  ports:
    - name: http
      port: 4010
```

This would take the following arguments:

* `--logical-name`: The application logical name e.g `frontend`
* `--version`: Optional version used to lookup the kubernetets service name
* `--service`: Optional specific kubernetets service name if known upfront

If neither `--service` or `--version` are provided an interactive UI will be
displayed showing the availible servies to route too.

```
compass route --logical-name frontend --version v1.2.3
compass route --logical-name frontend --service frontend-v1-2-3
compass route --logical-name frontend # spawns an interactive ui
```

## Development

### Dependencies

`protoc` is required to generate `go` code from `.proto` files.

Download the latest `protoc` release for your OS from https://github.com/google/protobuf/releases and follow the `readme.txt`.

#### `make` targets

* `make protoc`: runs protoc against the `_proto` directory
* `make migrations`: builds migrations into the application code
* `make ensure`: runs `dep esnure`
* `make test`: runs the test suite
* `make compass`: builds a `compass` binary in the `_bin` directory - CGO
  dynamically linked
* `make needle`: builds a `needle` binary in the `_bin` directory - CGO
  dynamically linked
* `APP=needle|compass make static`: builds a static binary for `needle` or
  `compass` in the `_bin` directory
* `APP=needle|compass make image` builds a docker image for `needle` or `compass`
