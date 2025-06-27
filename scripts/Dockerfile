FROM hashicorp/terraform:latest
RUN apk add jq \ 
    && apk add bash \
    && apk add go \
    && apk add python3 \
    && apk add make

RUN mkdir /src \
    && ln -s /usr/bin/python3 /bin/python

WORKDIR /src
COPY . .

RUN make install

ENTRYPOINT [ "/src/scripts/it.sh" ]
