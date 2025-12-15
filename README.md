# API CRUD em Go

Uma API REST construída com **Go**, utilizando **Clean Architecture**, **PostgreSQL**, **Docker**, **Migrations** e um design modular para escalabilidade. Atualmente, a API suporta operações CRUD para *Tasks* e está sendo expandida para uma **Movie API**, permitindo que os usuários acompanhem filmes assistidos, avaliem e gerenciem coleções pessoais.

## Funcionalidades
- Autenticação e autorização de usuários
- Gerenciamento de tarefas (operações CRUD)
- Avaliação e gerenciamento de coleções de filmes
- Arquitetura modular e escalável
- Integração com banco de dados PostgreSQL
- Ambiente Dockerizado para fácil configuração

## Começando

### Pré-requisitos
- Go 1.18+
- Docker & Docker Compose
- PostgreSQL

### Variáveis de Ambiente
Defina as seguintes variáveis de ambiente (ou use um arquivo `.env`):
- `DB_HOST`: Host do banco de dados
- `DB_PORT`: Porta do banco de dados
- `DB_USER`: Usuário do banco de dados
- `DB_PASSWORD`: Senha do banco de dados
- `DB_NAME`: Nome do banco de dados

### Executando com Docker Compose
```sh
docker compose up -d
```

### Executando Localmente
1. Exporte as variáveis de ambiente ou crie um arquivo `.env`.
2. Inicie o PostgreSQL.
3. Execute a aplicação:
   ```sh
   go run main.go
   ```

## Estrutura e Arquitetura do Projeto

Este projeto segue os princípios da Clean Architecture, inspirado pelo Domain Driven Design (DDD). O projeto está organizado em camadas:

- **domain/**: Entidades e regras de negócio
- **repository/**: Interfaces e implementações de acesso a dados
- **usecase/**: Casos de uso da aplicação
- **config/**: Configuração de banco de dados e ambiente
- **main.go**: Ponto de entrada da aplicação

### Exemplo de Estrutura
```
├── domain/
│   └── movie.go
├── repository/
│   └── movie_repository.go
├── usecase/
│   └── movie_usecase.go
├── config/
│   └── database.go
├── main.go
├── go.mod
├── docker-compose.yml
```

## Endpoints da API

### Autenticação
- **POST** `/login`: Login de usuário
- **POST** `/user`: Criar um novo usuário

### Tarefas
- **GET** `/tasks`: Listar todas as tarefas
- **POST** `/tasks`: Criar uma nova tarefa

### Filmes
- **GET** `/movies/search`: Buscar filmes
- **POST** `/movie/rate`: Avaliar um filme
- **GET** `/movie/rates/list`: Listar todas as avaliações de filmes
- **GET** `/movie/rate/:id/details`: Obter detalhes de uma avaliação específica

## Migrations
Use os scripts fornecidos para gerenciar as migrações do banco de dados:
- `migrate-up.sh`: Aplicar migrações
- `migrate-down.sh`: Reverter migrações

## Contribuindo
Contribuições são bem-vindas! Faça um fork do repositório e envie um pull request.

## Licença
Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

