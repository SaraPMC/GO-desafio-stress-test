.PHONY: build run test clean docker-build docker-run docker-clean help

help:
	@echo "Stress Test - Makefile Commands"
	@echo ""
	@echo "Local Development:"
	@echo "  make build          - Compila a aplicação"
	@echo "  make run            - Compila e executa com exemplo padrão"
	@echo "  make test           - Executa testes"
	@echo "  make clean          - Remove binários compilados"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-build   - Build da imagem Docker"
	@echo "  make docker-run     - Executa container com exemplo"
	@echo "  make docker-clean   - Remove imagem Docker"
	@echo ""
	@echo "Development:"
	@echo "  make deps           - Baixa dependências"
	@echo "  make fmt            - Formata código"
	@echo "  make lint           - Executa linter"

build:
	@echo "Compilando aplicação..."
	@go build -o stress-test main.go
	@echo "✅ Compilação concluída: stress-test"

run: build
	@echo "Executando teste padrão..."
	@./stress-test --url=http://google.com --requests=100 --concurrency=5

test:
	@echo "Executando testes..."
	@go test -v ./...

clean:
	@echo "Limpando arquivos compilados..."
	@rm -f stress-test stress-test.exe
	@echo "✅ Limpeza concluída"

deps:
	@echo "Baixando dependências..."
	@go mod download
	@echo "✅ Dependências instaladas"

fmt:
	@echo "Formatando código..."
	@go fmt ./...
	@echo "✅ Código formatado"

lint:
	@echo "Executando linter..."
	@go vet ./...
	@echo "✅ Linter concluído"

docker-build:
	@echo "Building Docker image..."
	@docker build -t stress-test:latest .
	@echo "✅ Docker image built: stress-test:latest"

docker-run: docker-build
	@echo "Executando container..."
	@docker run stress-test:latest --url=http://google.com --requests=100 --concurrency=5

docker-clean:
	@echo "Removendo imagem Docker..."
	@docker rmi stress-test:latest
	@echo "✅ Docker image removida"
