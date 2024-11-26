FROM golang:1.23

WORKDIR /app

ARG TARGETPLATFORM

# Install Task as build system
# install script apparently does not work on 32-bit Raspi so this ugly thing is needed
# see https://github.com/go-task/task/issues/1516#issuecomment-2347395883
RUN echo "$TARGETPLATFORM"


COPY Taskfile.yml ./
COPY go.mod ./

COPY internal/ ./internal
COPY pkg/ ./pkg
COPY cmd/ ./cmd

RUN go mod download && go mod verify

RUN mkdir bin

RUN if [ "$TARGETPLATFORM" = "linux/amd64" ]; then \
    curl -L https://taskfile.dev/install.sh > install-task.sh \
    && chmod +x ./install-task.sh \
    && ./install-task.sh -b $HOME/.local/bin v3.40.0 \
    && $HOME/.local/bin/task all; \
    elif [ "$TARGETPLATFORM" = "linux/armv7" ]; then \
    wget https://github.com/go-task/task/releases/download/v3.40.0/task_linux_arm.deb \
    && sudo dpkg -i task_linux_arm.deb \
    && task all; \
  fi

RUN $HOME/.local/bin/task all

EXPOSE 8080

CMD ["bin/goto"]