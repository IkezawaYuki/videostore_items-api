FROM golang:1.12.15

ENV REPO_URL=https://github.com/IkezawaYuki/videostore_items-api

ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL

ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

EXPOSE 8081
CMD ["./items-api"]