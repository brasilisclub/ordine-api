# Usar a imagem oficial do Go
FROM golang:1.23.3
WORKDIR /app
# ENV GOPROXY=direct

# # Baixar as dependências do Go
RUN go install github.com/spf13/cobra-cli@latest
RUN go install github.com/gin-gonic/gin@latest


