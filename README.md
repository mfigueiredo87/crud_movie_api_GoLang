Uma API RESTful em Go (Golang) para gerenciar filmes sem banco de dados.
O projeto usa Gorilla Mux para roteamento e armazena os filmes em memória usando slices.

🛠 Tecnologias

Golang 1.21+

Gorilla Mux (github.com/gorilla/mux)

JSON para comunicação entre cliente e servidor

HTTP Server padrão do Go (net/http)

🚀 Funcionalidades

Listar todos os filmes (GET /movies)

Obter filme por ID (GET /movies/{id})

Criar um novo filme (POST /movies)

Atualizar um filme existente (PUT /movies/{id})

Deletar um filme (DELETE /movies/{id})

Mensagens de sucesso retornadas em JSON após criar ou deletar

Estrutura do projecto:
go-movies-crud/
│
├── main.go          # Código principal do servidor e endpoints
├── go.mod           # Módulo Go
└── README.md        # Documentação do projeto

🏁 Como executar

1. Clonar o repositório
git clone https://github.com/teu-usuario/go-movies-crud.git
cd go-movies-crud

2. Inicializar o módulo Go (caso ainda não exista)
   go mod init go-movies-crud

3. Instalar dependências
   go get github.com/gorilla/mux@latest
   go mod tidy

4 Rodar o servidor
go run main.go

📡 Endpoints e Exemplos
1. Listar todos os filmes

URL: /movies

Método: GET

cURL:

curl -X GET http://localhost:8000/movies


Response:

[
  {
    "id": "1",
    "isbn": "438227",
    "title": "Movie One",
    "director": {
      "firstname": "John",
      "lastname": "Doe"
    }
  }
]

2. Obter filme por ID

URL: /movies/{id}

Método: GET

cURL:

curl -X GET http://localhost:8000/movies/1


Response:

{
  "id": "1",
  "isbn": "438227",
  "title": "Movie One",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}

3. Criar filme

URL: /movies

Método: POST

Request JSON:

{
  "isbn": "123456",
  "title": "Novo Filme",
  "director": {
    "firstname": "Oscar",
    "lastname": "Paul"
  }
}


cURL:

curl -X POST http://localhost:8000/movies \
-H "Content-Type: application/json" \
-d '{
    "isbn": "123456",
    "title": "Novo Filme",
    "director": {"firstname": "Oscar", "lastname": "Paul"}
}'


Response JSON:

{
  "message": "Movie adicionado com sucesso",
  "movie": {
    "id": "789456123",
    "isbn": "123456",
    "title": "Novo Filme",
    "director": {
      "firstname": "Oscar",
      "lastname": "Paul"
    }
  }
}

4. Atualizar filme

URL: /movies/{id}

Método: PUT

Request JSON: Mesma estrutura do POST

cURL:

curl -X PUT http://localhost:8000/movies/1 \
-H "Content-Type: application/json" \
-d '{
    "isbn": "999999",
    "title": "Filme Atualizado",
    "director": {"firstname": "John", "lastname": "Doe"}
}'


Response JSON: Filme atualizado com os novos dados.

5. Deletar filme

URL: /movies/{id}

Método: DELETE

cURL:

curl -X DELETE http://localhost:8000/movies/1


Response JSON:

{
  "message": "Movie apagado com sucesso",
  "moviews": [
    ... lista de filmes restantes ...
  ]
}

🖥 Testando no Postman

Abrir o Postman e criar uma nova requisição.

Selecionar o método HTTP (GET, POST, PUT, DELETE) conforme o endpoint.

Inserir a URL: http://localhost:8000/movies ou http://localhost:8000/movies/{id}.

No caso de POST ou PUT, selecionar Body > raw > JSON e colar o JSON.

Adicionar header:

Content-Type: application/json


Enviar requisição e verificar o JSON de resposta.

⚡ Observações

Todos os dados são armazenados em memória, portanto serão perdidos ao reiniciar o servidor.

O campo ID é gerado automaticamente pelo servidor.

Se algum campo não for enviado no JSON, ele será inicializado com valor padrão ("Sem Título", "N/A", etc).

Mensagens de sucesso são retornadas em JSON para facilitar integração com front-end.
