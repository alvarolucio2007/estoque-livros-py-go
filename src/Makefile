.PHONY: test-all test-db test-verifiers test-api

# Comando principal que roda tudo na ordem
test-all: test-db test-verifiers test-api
	@echo "✅ Todos os testes passaram com sucesso!"

test-db:
	@echo "Running Database Connection tests..."
	go test ./testing/... -run TestDatabaseConnection

test-verifiers:
	@echo "Running Verifier tests..."
	go test ./testing/... -run TestVerifiers
test-api:
	@echo "Running API Route tests..."
	go test ./testing/... -run TestAPI
