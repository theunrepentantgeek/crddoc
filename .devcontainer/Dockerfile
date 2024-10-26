# See here for image contents: https://github.com/devcontainers/images/blob/main/src/go/.devcontainer/Dockerfile

# This is pinned to a particular version of go:
FROM mcr.microsoft.com/devcontainers/go:1.23

# https://docs.docker.com/engine/reference/builder/#automatic-platform-args-in-the-global-scope
ARG TARGETARCH

# APT dependencies
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get -y install --no-install-recommends bash-completion software-properties-common lsb-release

COPY install-dependencies.sh .
RUN ./install-dependencies.sh devcontainer && rm install-dependencies.sh

# Setup go-task completions
RUN curl -sL "https://raw.githubusercontent.com/go-task/task/refs/tags/v3.39.2/completion/bash/task.bash" > "/etc/.task.completion.sh" \
    && echo 'source /etc/.task.completion.sh' >> "/etc/bash.bashrc"

ENTRYPOINT ["tail"]
CMD ["-f","/dev/null"]
