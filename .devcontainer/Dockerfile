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

# these are all the default values except for the last one (USE_MOBY) which is what we want to set.
# currently packages.microsoft.com does not have Moby packages for Debian Bullseye, which the
# devcontainer image is based on: https://github.com/microsoft/vscode-dev-containers/issues/1008
RUN bash /tmp/library-scripts/docker-debian.sh true /var/run/docker-host.sock /var/run/docker.sock automatic false

ENTRYPOINT ["/usr/local/share/docker-init.sh"]
CMD ["sleep", "infinity"]
