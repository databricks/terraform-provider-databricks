ARG VARIANT=1.19-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/go:${VARIANT}

ARG TERRAFORM_VERSION=1.0.0
ARG TARGETARCH
ARG USERNAME=vscode

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
    && curl -sSL -o /tmp/docker-downloads/terraform.zip https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_${TARGETARCH}.zip \
    && unzip /tmp/docker-downloads/terraform.zip \
    && mv terraform /usr/local/bin \
    && terraform -install-autocomplete

USER $USERNAME

# Install Go tools
RUN \
    # --> Goimports
    go install -v golang.org/x/tools/cmd/goimports@latest \
    # # --> Gotestsum
    && go install -v gotest.tools/gotestsum@latest \
    # # --> Go symbols and outline for go to symbol support and test support
    && go install github.com/acroca/go-symbols@v0.1.1 \
    # # --> Static checker
    && go install honnef.co/go/tools/cmd/staticcheck@v0.3.2 \
    # # --> TF formatter
    && go install github.com/katbyte/terrafmt@latest
