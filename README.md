# API CRUD em Go

Uma API REST construída com **Go**, utilizando **Clean Architecture**, **PostgreSQL**, **Docker**, **Migrations** e um design modular para escalabilidade. Atualmente, **Movie API**, permite que os usuários acompanhem filmes assistidos, avaliem e gerenciem coleções pessoais.

## Funcionalidades
- Autenticação e autorização de usuários
- Avaliação e gerenciamento de coleções de filmes
- Arquitetura modular e escalável
- Integração com banco de dados PostgreSQL
- Ambiente Dockerizado para fácil configuração

## Começando

### Pré-requisitos
- Go 1.18+
- Docker & Docker Compose
- PostgreSQL

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

