FROM hashicorp/terraform:latest
RUN apk add jq \ 
    && apk add bash \
    && apk add go \
    && apk add python3 \
    && apk add make \
    && go install gotest.tools/gotestsum@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest

RUN mkdir /src \
    && ln -s /root/go/bin/gotestsum /bin/gotestsum \
    && ln -s /usr/bin/python3 /bin/python

WORKDIR /src
COPY . .

RUN make install

ENTRYPOINT [ "/src/scripts/it.sh" ]
