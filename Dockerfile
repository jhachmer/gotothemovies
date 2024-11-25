FROM golang:1.23

WORKDIR /app

# Install Task as build system
RUN curl -L https://taskfile.dev/install.sh > install-task.sh
RUN chmod +x ./install-task.sh
RUN ./install-task.sh -b $HOME/.local/bin v3.40.0

COPY go.mod ./
COPY Taskfile.yml ./
COPY *.go ./

RUN mkdir bin
RUN $HOME/.local/bin/task all

EXPOSE 8080

CMD ["bin/goto"]