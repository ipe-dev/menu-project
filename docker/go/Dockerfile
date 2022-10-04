FROM golang:1.19.1-alpine3.16

RUN apk --update-cache --no-cache add \
    git 

RUN apk --update-cache --no-cache add \
    vim 

CMD [ "go", "run", "/go/src/app/cmd/api/main.go" ] 