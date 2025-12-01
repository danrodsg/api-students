# 📚 API de Cadastro de Estudantes (GoLang)

- Uma API RESTful completa desenvolvida em Go (Golang), utilizando uma arquitetura robusta para operações CRUD (Create, Read, Update, Delete) e persistência de dados de estudantes. Este projeto utiliza o SQLite para armazenamento de dados, garantindo facilidade de setup.

## 🚀 Tecnologias

-Go (Golang): Linguagem de programação principal.

-Arquitetura RESTful: Padrões de design para a comunicação da API.

-SQLite: Banco de dados embutido e baseado em arquivo para persistência.

-Pacote SQL/ORM: (Ex: gorm, database/sql + mattn/go-sqlite3).

## ✨ Funcionalidades (Endpoints)

GET	 -   /students	 -  Lista todos os estudantes (ativos e inativos)

GET	 - /students/{id} - Obtém as informações de um estudante específico pelo seu ID.

GET	 - /students?active=<true/false>	- Lista os estudantes, filtrando por status de ativo (true) ou inativo (false).

POST	- /students -	Cria um novo estudante no cadastro.

PUT	- /students/{id}	- Atualiza as informações de um estudante existente.

DELETE - /students/{id}	- Exclui (ou marca como inativo) um estudante pelo seu ID.

## 💾 Estrutura do Estudante 

Nome - string - Nome completo do estudante.

CPF - int - Cadastro de Pessoa Física.

E-Mail - string - Endereço de e-mail.

Idade - int - Idade do estudante.

Ativo - bool - Status de atividade do estudante (true para ativo, false para inativo).

# 🛠️ Como Executar o Projeto

## 1. Pré-requisitos
- Golang: Versão 1.18 ou superior.
- Git: Para clonar o repositório.
  
 ## 2. Clonar e Instalar Dependências 
 
 ## 🤖 Bash

- git clone github.com/danrodsg/api-students/api
- cd api-students
- go mod tidy

## 3. Executar a API 

## 🤖 Bash (O arquivo do banco de dados SQLite será criado automaticamente na primeira execução)

- go run main.go
- A API estará rodando em http://localhost:8080.

## 🧪 Testes 

## 🤖 Bash : Para testar se a API está no ar:)

-curl -X GET http://localhost:8080/students



## 🤝 Contribuições são bem-vindas! Abra uma Issue ou envie um Pull Request!

# ✉️ Contato

- Daniel Rodrigues / https://github.com/danrodsg
- E-mail: danielrods2004@gmail.com

  
  



