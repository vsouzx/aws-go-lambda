# Author: Alan Ramalho

# Configurações
BINARY_NAME=lambda-handler
ZIP_NAME=lambda.zip
GOOS=linux
GOARCH=arm64

.PHONY: all build clean zip

all: clean build zip mov

build:
	@echo "🔧 Buildando para $(GOOS)/$(GOARCH)..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -ldflags="-s -w" -o $(BINARY_NAME) main.go

zip:
	@echo "📦 Gerando $(ZIP_NAME)..."
	@zip -j $(ZIP_NAME) $(BINARY_NAME)

clean:
	@echo "🧹 Limpando arquivos antigos..."
	@rm -f $(BINARY_NAME) $(ZIP_NAME)

mov:
	@echo "📦 Movendo $(ZIP_NAME) para o diretório de destino..."
	@mv $(ZIP_NAME) infra/