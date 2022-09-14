FROM        golang:1.18.6
RUN         mkdir -p /studygroup
WORKDIR     /studygroup
COPY        . .
RUN         cd cmd; go mod download; go build -o app
ENTRYPOINT  ["./cmd/app"]
