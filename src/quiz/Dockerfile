# Build the frontend. This stage is only needed when web/dist is out of date;
# the checked-in dist/ means a plain `go build` also works without Node.js.
FROM node:22-alpine AS web
WORKDIR /web
COPY src/quiz/web/package.json src/quiz/web/package-lock.json* ./
RUN npm install
COPY src/quiz/web/ ./
RUN npm run build

# Build the server. The repository root is the build context, because the quiz
# packages import through the module path of the whole course repository.
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=web /web/dist ./src/quiz/web/dist
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /quizserver ./src/quiz/cmd/quizserver

# The final image carries the static binary and the quiz files, nothing else.
FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /quizserver /app/quizserver
COPY --from=build /src/src/quiz/quizzes /app/quizzes
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/quizserver"]
CMD ["-addr", ":8080", "-dir", "/app/quizzes"]
