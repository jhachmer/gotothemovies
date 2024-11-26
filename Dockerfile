FROM golang:1.23

WORKDIR /app

# Install Task as build system
RUN curl -L https://taskfile.dev/install.sh > install-task.sh
RUN chmod +x ./install-task.sh
RUN ./install-task.sh -b $HOME/.local/bin v3.40.0

COPY Taskfile.yml ./
COPY go.mod ./
COPY internal/ ./internal
COPY pkg/ ./pkg
COPY cmd/ ./cmd

RUN mkdir bin
RUN $HOME/.local/bin/task all

EXPOSE 8080

CMD ["bin/goto"]