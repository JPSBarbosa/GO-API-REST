# GO-API-REST

Uma API em GO, com CRUDs de usuários e produtos, seguindo a arquitetura proposta no Clean Architecture de Robert Cecil Martin. Utiliza GORM, Gin Gonic e SQL.

## Índice

- [Descrição](#descrição)
- [Instalação](#instalação)
- [Licença](#licença)
- [Contato](#contato)

## Descrição

Este projeto é uma API RESTful desenvolvida em Go, que implementa operações CRUD (Create, Read, Update, Delete) para usuários e produtos. A API utiliza o Gin Gonic para o roteamento, GORM para interações com o banco de dados e SQL como o banco de dados subjacente.

## Instalação

Instruções para instalar e configurar o projeto.

```bash
# Clone o repositório
git clone https://github.com/JPSBarbosa/GO-API-REST.git

# Entre no diretório do projeto
cd ./cmd/api

# Instale as dependências
go mod tidy

# Configure seu banco de dados (ajuste o arquivo .env conforme necessário)
# Exemplo de configuração do .env:
# DB_HOST=localhost
# DB_USER=seu_usuario
# DB_PASSWORD=sua_senha
# DB_NAME=seu_banco_de_dados
# DB_PORT=3306

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

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

## Contato

João Pedro S. Barbosa - jpbarbosa.dev@gmail.com
