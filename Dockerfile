FROM golang:latest
WORKDIR /app
COPY . /app
RUN go mod download
RUN go build -o main main.go
EXPOSE 8080
RUN chmod +x ./main
CMD [ "./main" ]