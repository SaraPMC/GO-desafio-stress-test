# 🚀 Stress Test - Teste de Carga em Go

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![Docker](https://img.shields.io/badge/Docker-Required-blue.svg)](https://docker.com)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Production%20Ready-brightgreen.svg)]()

## 📋 Sobre o Projeto

**Stress Test** é uma ferramenta CLI (Command Line Interface) poderosa e eficiente desenvolvida em **Go** para realizar testes de carga em serviços web. Com suporte a múltiplas requisições simultâneas e relatórios detalhados, ela permite avaliar a capacidade e performance de seus servidores.

### 🎯 Objetivo

Criar um sistema CLI robusto que permite aos desenvolvedores e DevOps engineers:
- ✅ Testar a capacidade de carga de um serviço web
- ✅ Simular múltiplas requisições simultâneas (concorrência)
- ✅ Gerar relatórios detalhados com métricas de performance
- ✅ Identificar gargalos e limites de capacidade

### 🏆 Funcionalidades Implementadas

- ✅ **Requisições HTTP** - Suporte completo a requisições GET
- ✅ **Concorrência** - Controle total do número de requisições simultâneas
- ✅ **Relatório Completo** - Métricas detalhadas de performance
- ✅ **Docker Support** - Containerização pronta para uso
- ✅ **CLI Intuitiva** - Flags simples e diretas
- 📊 **Estatísticas Avançadas** - Min, Máx, Média de latência
- 🔢 **Distribuição de Status** - Análise dos códigos HTTP retornados

---

## 📝 Requisitos do Projeto

### Entrada de Parâmetros via CLI

| Flag | Descrição | Tipo | Obrigatório |
|------|-----------|------|------------|
| `--url` | URL do serviço a ser testado | string | ✅ Sim |
| `--requests` | Número total de requests | int | ✅ Sim |
| `--concurrency` | Número de chamadas simultâneas | int | ❌ Não (padrão: 1) |

### Execução do Teste

O sistema realiza:
- 🌐 Requisições HTTP para a URL especificada
- ⚡ Distribuição de requests de acordo com nível de concorrência
- 🔄 Garantia de cumprimento do número total de requests
- ⏱️ Medição precisa de tempos de resposta

### Geração de Relatório

O relatório final contém:
- ⏱️ **Tempo total** gasto na execução
- 📊 **Quantidade total** de requests realizados
- ✅ **Requests com status 200** (sucesso)
- 🔢 **Distribuição de códigos HTTP** (404, 500, etc.)
- 📈 **Estatísticas avançadas**:
  - Tempo mínimo, máximo e médio de resposta
  - Taxa de requisições por segundo

---

## 🚀 Como Executar

### Pré-requisitos

- [Go 1.21+](https://golang.org/dl/)
- [Docker](https://www.docker.com/) (opcional, para execução containerizada)

### Instalação Local

```bash
# Clonar repositório
git clone https://github.com/SaraPMC/GO-desafio-stress-test.git
cd stress-test

# Baixar dependências
go mod download

# Compilar
go build -o stress-test main.go
```

### Execução Local

```bash
# Exemplo básico
./stress-test --url=https://google.com --requests=1000

# Com concorrência
./stress-test --url=https://google.com --requests=1000 --concurrency=10

# Windows
stress-test.exe --url=https://google.com --requests=1000 --concurrency=10
```

### Execução via Docker

```bash
# Build da imagem
docker build -t stress-test:latest .

# Executar container
docker run stress-test:latest --url=https://google.com --requests=1000 --concurrency=10

# Exemplo com Google.com (1000 requisições, 10 simultâneas)
docker run stress-test:latest --url=http://google.com --requests=1000 --concurrency=10

# Exemplo com localhost (100 requisições, 5 simultâneas)
docker run stress-test:latest --url=http://localhost:8080 --requests=100 --concurrency=5
```

---

## 📊 Exemplos de Uso

### Teste Simples

```bash
./stress-test --url=https://api.example.com/health --requests=100
```

**Output esperado:**
```
🚀 Iniciando teste de carga
📍 URL: https://api.example.com/health
📊 Total de requests: 100
⚡ Concorrência: 1

============================================================
📋 RELATÓRIO DE TESTE DE CARGA
============================================================

⏱️  Tempo total: 5.234s
📊 Total de requests: 100
✅ Requests com status 200: 100 (100.00%)
❌ Requests com falha: 0

📈 Estatísticas de Duração:
   ⚡ Mínimo: 45ms
   ⏱️  Médio: 52.34ms
   🐢 Máximo: 120ms

🔢 Distribuição de Códigos HTTP:
   HTTP 200: 100 requisições

📊 Taxa de requisições por segundo: 19.10 req/s

============================================================
```

### Teste com Alta Concorrência

```bash
./stress-test --url=https://api.example.com/endpoint --requests=5000 --concurrency=50
```

Este comando enviará 5000 requisições com 50 requisições simultâneas (5000 / 50 = 100 lotes).

---

## 🏗️ Arquitetura do Projeto

```
stress-test/
├── main.go              # Ponto de entrada
├── cmd/
│   └── root.go          # Lógica CLI e orquestração
├── go.mod               # Dependências Go
├── go.sum               # Checksums
├── Dockerfile           # Build containerizado
├── .gitignore           # Exclusões Git
└── README.md            # Esta documentação
```

### Fluxo de Execução

```
┌─────────────────────────────────────────┐
│      Entrada de Flags da CLI            │
│  (url, requests, concurrency)           │
└──────────────┬──────────────────────────┘
               ↓
┌─────────────────────────────────────────┐
│    Inicializar Pool de Workers          │
│    (goroutines concorrentes)            │
└──────────────┬──────────────────────────┘
               ↓
┌─────────────────────────────────────────┐
│    Enviar Requisições HTTP              │
│    (distribuídas entre workers)         │
└──────────────┬──────────────────────────┘
               ↓
┌─────────────────────────────────────────┐
│    Coletar Resultados                   │
│    (status, latência, etc)              │
└──────────────┬──────────────────────────┘
               ↓
┌─────────────────────────────────────────┐
│    Calcular Estatísticas                │
│    (min, max, média, taxa)              │
└──────────────┬──────────────────────────┘
               ↓
┌─────────────────────────────────────────┐
│    Exibir Relatório Formatado           │
│    (visual e informativo)               │
└─────────────────────────────────────────┘
```

---

## 🛠️ Tecnologias Utilizadas

- **Go 1.21** - Linguagem de programação
- **Cobra** - Framework para CLI robusta
- **HTTP Client** - Biblioteca padrão do Go
- **Goroutines** - Concorrência nativa
- **Sync** - Sincronização de goroutines
- **Docker** - Containerização

---

## 📈 Métricas Geradas

O relatório fornece as seguintes métricas:

| Métrica | Descrição |
|---------|-----------|
| Tempo Total | Duração total do teste |
| Total de Requests | Número absoluto de requisições |
| Sucesso (200) | Quantidade e percentual de respostas 200 |
| Falhas | Requisições com status diferente de 200 |
| Mínimo | Menor tempo de resposta |
| Máximo | Maior tempo de resposta |
| Médio | Tempo médio de resposta |
| Taxa/s | Requisições por segundo |
| Distribuição HTTP | Contagem de cada código de status |

---

## 🔍 Casos de Uso

### 1️⃣ Teste de Capacidade de API

```bash
./stress-test --url=https://api.myservice.com/v1/users --requests=10000 --concurrency=100
```

### 2️⃣ Teste de Load Balancer

```bash
./stress-test --url=http://load-balancer.internal:8080 --requests=5000 --concurrency=50
```

### 3️⃣ Teste de Database Query Endpoint

```bash
./stress-test --url=http://localhost:3000/api/products --requests=1000 --concurrency=20
```

### 4️⃣ Teste de Microserviço em Produção

```bash
docker run stress-test:latest \
  --url=https://api.prod.com/service \
  --requests=50000 \
  --concurrency=200
```

---

## 🚨 Dicas de Performance

- 📈 **Comece pequeno**: Teste com 100 requests antes de aumentar
- ⚡ **Incremente gradualmente**: Aumente concorrência em passos (10 → 50 → 100)
- 🌐 **Considere latência de rede**: Testes remotos são mais lentos
- 💾 **Monitore recursos**: Use `docker stats` ao testar containers
- 🔍 **Analise resultados**: Procure por picos ou padrões de falha

---

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

## 👨‍💻 Contribuindo

Contribuições são bem-vindas! Por favor:

1. Faça um Fork do repositório
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

---

## ❓ FAQ

**P: Qual é o máximo de requisições que posso fazer?**
R: Não há limite fixo, depende dos recursos do seu sistema e da URL testada.

**P: Posso testar URLs internas/localhost?**
R: Sim, absolutamente. Use `http://localhost:8080` ou o IP interno.

**P: Como saber se meu serviço aguenta a carga?**
R: Se 90% das requisições retornam 200 e a latência é aceitável, seu serviço está bem.

**P: O Docker é obrigatório?**
R: Não, você pode compilar e rodar localmente com Go instalado.

---

**Desenvolvido com ❤️ em Go**