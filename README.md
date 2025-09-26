# Leilão em Go (GoExpert)

Este é um projeto de uma API de leilões desenvolvida em Go.

## Como Executar o Projeto

Existem duas maneiras de executar este projeto: utilizando Docker (recomendado) ou em um ambiente de desenvolvimento local.

### 1. Rodando com Docker (Recomendado)

Esta é a maneira mais simples de subir a aplicação e todas as suas dependências.

**Pré-requisitos:**
- Docker
- Docker Compose

**Passos:**

1.  Clone o repositório:
    ```bash
    git clone https://github.com/devfullcycle/labs-auction-goexpert.git
    cd labs-auction-goexpert
    ```

2.  Execute o Docker Compose:
    ```bash
    docker-compose up -d
    ```

A aplicação estará disponível em `http://localhost:8080`.

### 2. Rodando em Ambiente de Desenvolvimento Local

Para executar a aplicação localmente, você precisará ter o Go e o MongoDB instalados e configurados.

**Pré-requisitos:**
- Go (versão 1.20 ou superior)
- MongoDB

**Passos:**

1.  **Inicie o MongoDB:** Certifique-se de que uma instância do MongoDB esteja em execução.

2.  **Configure as Variáveis de Ambiente (Opcional):**
    A aplicação utiliza a biblioteca Viper para gerenciar as configurações, com valores padrão definidos no código. Se precisar sobrescrevê-los, você pode exportar as seguintes variáveis de ambiente no seu terminal:

    ```bash
    export MONGODB_URL="mongodb://admin:admin@localhost:27017/auctions?authSource=admin"
    export MONGODB_DB="auctions"
    export AUCTION_INTERVAL="1m"
    export BATCH_INSERT_INTERVAL="30s"
    export MAX_BATCH_SIZE="5"
    ```

3.  **Execute a Aplicação:**
    Navegue até a raiz do projeto e execute o seguinte comando:
    ```bash
    go run cmd/auction/main.go
    ```

A aplicação estará disponível em `http://localhost:8080`.
