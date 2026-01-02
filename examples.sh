#!/bin/bash

# Script de exemplo para executar o stress-test

# Teste 1: Teste simples no Google
echo "=== Teste 1: Google (100 requests, concorrência 5) ==="
./stress-test --url=http://google.com --requests=100 --concurrency=5

# Teste 2: Teste mais intenso
# echo ""
# echo "=== Teste 2: Teste Intenso (1000 requests, concorrência 20) ==="
# ./stress-test --url=http://google.com --requests=1000 --concurrency=20

# Teste 3: Localhost (descomente após subir um servidor local)
# echo ""
# echo "=== Teste 3: Localhost (500 requests, concorrência 10) ==="
# ./stress-test --url=http://localhost:8080 --requests=500 --concurrency=10
