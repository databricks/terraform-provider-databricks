FROM golang:1.13
WORKDIR /go/src/github.com/databrickslabs/databricks-terraform/
RUN curl -sSL "https://github.com/gotestyourself/gotestsum/releases/download/v0.4.2/gotestsum_0.4.2_linux_amd64.tar.gz" | tar -xz -C /usr/local/bin gotestsum
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.25.0
COPY . .
RUN CGO_ENABLED=0 make vendor build

FROM hashicorp/terraform:latest
COPY --from=0 /go/src/github.com/databrickslabs/databricks-terraform/terraform-provider-databricks /root/.terraform.d/plugins/
RUN ls ~/.terraform.d/plugins
