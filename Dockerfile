FROM golang:1.15 AS builder

ENV GOOS linux

WORKDIR ${GOPATH}/src/github.com/jefersonhuan/action-jira-transition
COPY . .

RUN go mod download
RUN go build main.go

FROM builder

COPY --from=builder go/src/github.com/jefersonhuan/action-jira-transition/main /usr/bin/jira-transition

ENTRYPOINT ["/usr/bin/jira-transition"]