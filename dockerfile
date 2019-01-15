FROM golang



# create a working directory
WORKDIR /go/src/app


# add source code
ADD RestFull RestFull


EXPOSE 9191


# run main.go
CMD ["go", "run", "RestFull/src/webservice/main.go"]
