FROM golang:1.11-stretch AS build

WORKDIR ${GOPATH}/src/github.com/watercompany/cx-tracker/

ADD cmd/ cmd/
ADD src/ src/
ADD vendor/ vendor/

ENV GOARCH="amd64" \
    CGO_ENABLED="0" \
    GOOS="linux"

RUN go install ./cmd/
RUN sh -c "mkdir -p /tmp/files/usr/bin"
RUN cp ${GOPATH}/bin/cmd /tmp/files/usr/bin/cx-tracker

FROM busybox

VOLUME /root/.cx-tracker/

COPY --from=build /tmp/files /
COPY script/cx-tracker-migration/ /etc/cx-tracker/cx-tracker-migration/

EXPOSE 8085

ENTRYPOINT ["cx-tracker"]
