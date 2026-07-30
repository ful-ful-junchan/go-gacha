FROM golang:1.22

RUN groupadd -r app && useradd -r -g app -m -d /home/app app

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN chown -R app:app /app

USER app

CMD ["go", "run", "./cmd/api"]
