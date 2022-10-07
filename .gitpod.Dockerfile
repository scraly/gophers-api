ARG GITPOD_IMAGE=gitpod/workspace-base:latest
FROM ${GITPOD_IMAGE}

RUN brew install go-task/tap/go-task

RUN brew tap go-swagger/go-swagger && brew install go-swagger
