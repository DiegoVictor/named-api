FROM golang:latest
WORKDIR /app
RUN apt-get update -y && apt-get install python3 python3-pip -y \
  && ln -s /usr/bin/python3 /usr/bin/python
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
RUN go build -tags netgo -ldflags '-s -w' -o app
EXPOSE 8080
CMD ["./app"]