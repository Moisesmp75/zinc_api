FROM node:20.9.0 as frontend-builder

WORKDIR /app/web

COPY web/package*.json ./

RUN npm install

COPY web .

RUN npm run build

FROM golang:1.21.3 as backend-builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

COPY --from=frontend-builder /app/web/dist /app/web/dist

RUN go build -o main .

FROM golang:1.21.3

WORKDIR /app

COPY --from=backend-builder /app/main .

EXPOSE 3000

CMD ["./main"]