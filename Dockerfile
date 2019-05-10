FROM golang:1.11-stretch AS build

COPY . ${GOPATH}/src/github.com/watercompany/cx-tracker

WORKDIR ${GOPATH}/src/github.com/watercompany/cx-tracker

ENV GOARCH="amd64" \
		CGO_ENABLED="0" \
		GOOS="linux"

RUN go install ./cmd/
RUN sh -c "mkdir -p /tmp/files/usr/bin"
RUN cp ${GOPATH}/bin/cmd /tmp/files/usr/bin/cx-tracker

FROM busybox

VOLUME /root/.skywire-uptime-system/

COPY --from=build /tmp/files /

EXPOSE 8085

ENTRYPOINT ["cx-tracker"]
