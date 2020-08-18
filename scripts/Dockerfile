ARG TERRAFORM_VERSION=0.12.29

FROM hashicorp/terraform:${TERRAFORM_VERSION}
RUN apk add jq \ 
    && apk add bash \
    && apk add go \
    && apk add python3 \
    && apk add make \
    && go get gotest.tools/gotestsum

RUN mkdir /src \
    && ln -s /root/go/bin/gotestsum /bin/gotestsum \
    && ln -s /usr/bin/python3 /bin/python

WORKDIR /src
COPY . .

ENTRYPOINT [ "/src/scripts/run.sh" ]