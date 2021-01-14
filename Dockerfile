FROM golang:1.15.6-alpine3.12
RUN mkdir /app
COPY . /app
WORKDIR /app
EXPOSE 8080
RUN go build -o server . 
CMD [ "/app/server" ]