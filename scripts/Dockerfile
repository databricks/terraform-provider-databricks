FROM hashicorp/terraform:latest
RUN apk add jq \ 
    && apk add bash \
    && apk add go \
    && apk add python3 \
    && apk add make \
    && go get gotest.tools/gotestsum \
    && go get honnef.co/go/tools/cmd/staticcheck \
    && go get github.com/katbyte/terrafmt

RUN mkdir /src \
    && ln -s /root/go/bin/gotestsum /bin/gotestsum \
    && ln -s /usr/bin/python3 /bin/python

WORKDIR /src
COPY . .

RUN make install

ENTRYPOINT [ "/src/scripts/it.sh" ]