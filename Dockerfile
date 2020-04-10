FROM golang:1.13-alpine
WORKDIR /go/src/github.com/databrickslabs/databricks-terraform/
COPY . .
RUN apk add --update make
RUN make build

FROM hashicorp/terraform:latest
COPY --from=0 /go/src/github.com/databrickslabs/databricks-terraform/terraform-provider-db /root/.terraform.d/plugins/
RUN ls ~/.terraform.d/plugins
