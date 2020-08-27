FROM node:14-alpine AS frontend

ARG GITHUB_TOKEN

WORKDIR /src
COPY package.json yarn.lock .npmrc ./

RUN echo "//npm.pkg.github.com/:_authToken=${GITHUB_TOKEN}" >> ~/.npmrc
RUN yarn

#

FROM golang

RUN mkdir /src
WORKDIR /src

COPY main.go go.mod go.sum ./
COPY --from=frontend /src/node_modules ./node_modules

RUN cd /src && go build -o main && rm main.go go.mod go.sum

ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ./main
