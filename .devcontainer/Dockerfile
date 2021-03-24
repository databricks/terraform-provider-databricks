FROM ubuntu:20.04

ARG TERRAFORM_VERSION=0.14.8

# Configure apt, install packages and tools
RUN \
    apt-get update \
    && apt-get -y install --no-install-recommends apt-utils bash-completion curl icu-devtools jq \
    # Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
    && apt-get -y install git iproute2 procps lsb-release unzip wget gcc build-essential \
    # Github actions has Azure CLI installed
    && curl -sL https://aka.ms/InstallAzureCLIDeb | bash \
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Install terraform 
RUN \
    mkdir -p /tmp/docker-downloads \
    && curl -sSL -o /tmp/docker-downloads/terraform.zip https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
    && unzip /tmp/docker-downloads/terraform.zip \
    && mv terraform /usr/local/bin \
    && terraform -install-autocomplete

# Save command line history 
RUN \
    echo "export HISTFILE=/root/commandhistory/.bash_history" >> "/root/.bashrc" \
    && echo "export PROMPT_COMMAND='history -a'" >> "/root/.bashrc" \
    && mkdir -p /root/commandhistory \
    && touch /root/commandhistory/.bash_history \
    && echo "source /usr/share/bash-completion/bash_completion" >> "/root/.bashrc"

# GoLang
RUN \
    wget -c https://dl.google.com/go/go1.15.2.linux-amd64.tar.gz -O - | tar -xz -C /usr/local \
    && export PATH=$PATH:/usr/local/go/bin \
    && echo 'export PATH="/root/go/bin:/usr/local/go/bin:$PATH"' >> /root/.bashrc \
    && ln -s /usr/local/go/bin/go /bin/go

ENV GO111MODULE=on

# Install Go tools
RUN \
    # --> Delve for debugging
    go get github.com/go-delve/delve/cmd/dlv \
    # --> Go language server
    && go get -v golang.org/x/tools/gopls \
    # --> Goimports
    && go get golang.org/x/tools/cmd/goimports \
    # --> Gotestsum
    && go get gotest.tools/gotestsum \
    # --> Go symbols and outline for go to symbol support and test support 
    && go get github.com/acroca/go-symbols@v0.1.1 && go get github.com/ramya-rao-a/go-outline@7182a932836a71948db4a81991a494751eccfe77 \
    # --> GolangCI-lint
    && go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.25.0 \

USER developer