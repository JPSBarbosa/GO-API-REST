# GO-API-REST

Uma API em GO, com CRUDs de usuários e produtos, seguindo a arquitetura proposta no Clean Architecture de Robert Cecil Martin. Utiliza GORM, Gin Gonic e SQL.

## Índice

- [Descrição](#descrição)
- [Instalação](#instalação)
- [Uso](#uso)
- [Contribuição](#contribuição)
- [Licença](#licença)
- [Contato](#contato)

## Descrição

Este projeto é uma API RESTful desenvolvida em Go, que implementa operações CRUD (Create, Read, Update, Delete) para usuários e produtos. Ele segue os princípios do Clean Architecture de Robert Cecil Martin, garantindo um código modular e fácil de manter. A API utiliza o Gin Gonic para o roteamento, GORM para interações com o banco de dados e SQL como o banco de dados subjacente. Ideal para aplicações que exigem uma estrutura robusta e escalável.

## Instalação

Instruções para instalar e configurar o projeto.

```bash
# Clone o repositório
git clone [https://github.com/usuario/projeto.git](https://github.com/JPSBarbosa/GO-API-REST.git)

# Entre no diretório do projeto
cd ./cmd/api

# Instale as dependências
go mod tidy

# Inicie a aplicação
go run main.go
```

## Endpoints Disponíveis

### Usuários

- `POST /users` - Cria um novo usuário
- `GET /users/:id` - Obtém um usuário específico
- `PUT /users/:id` - Atualiza um usuário específico
- `DELETE /users/:id` - Exclui um usuário específico

### Produtos

- `POST /products` - Cria um novo produto
- `GET /products/:id` - Obtém detalhes de um produto específico
- `PUT /products/:id` - Atualiza um produto específico
- `DELETE /products/:id` - Exclui um produto específico
