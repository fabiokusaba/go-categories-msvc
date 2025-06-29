# Category Microservice

Este é um microserviço desenvolvido em **Go** utilizando o framework **Gin**, responsável pelo **gerenciamento de categorias**. Ele permite a realização de operações **CRUD (Create, Read, Update, Delete)** sobre a entidade `Category`.

## ✨ Tecnologias Utilizadas

- [Go](https://golang.org/) (Golang)
- [Gin](https://gin-gonic.com/) - Web framework para Go
- In Memory Database

## 📦 Estrutura da Entidade

A entidade `Category` possui os seguintes campos:

```go
type Category struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

## 🔧 Funcionalidades

- [X] Criar categoria POST /categories

- [X] Listar todas as categorias GET /categories

- [ ] Buscar categoria por ID GET /categories/:id

- [ ] Atualizar categoria PUT /categories/:id

- [ ] Excluir categoria DELETE /categories/:id

## 📮 Exemplos de Requisições

### Criar categoria
``` http
POST /categories
Content-Type: application/json

{
  "name": "Tecnologia"
}
```

### Resposta de sucesso
``` json
{
  "success": true
}
```
