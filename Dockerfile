FROM golang:1.22.1 as builder
ARG CGO_ENABLED=0
WORKDIR /app

# Copier les fichiers go.mod et go.sum et télécharger les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le reste des fichiers/dossiers du projet
COPY . .

# Se déplacer dans le dossier où se trouve le main.go et construire l'application
WORKDIR /app/cmd/rcp-api-data
RUN go build -o server .

# Étape finale, configuration de l'image minimaliste
FROM scratch
COPY --from=builder /app/cmd/rcp-api-data/server /server
ENTRYPOINT ["/server"]
