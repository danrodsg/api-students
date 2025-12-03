# 📚 API de Cadastro de Estudantes (GoLang)

[![Go](https://github.com/golang/go/blob/master/assets/badge.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Esta é uma **API RESTful** completa desenvolvida em Go (Golang) para o gerenciamento de um cadastro de estudantes. O projeto utiliza uma **arquitetura robusta** para operações CRUD (Create, Read, Update, Delete) e emprega o **SQLite** como banco de dados embutido para garantir facilidade de configuração e portabilidade.

## 🚀 Tecnologias Utilizadas

| Tecnologia | Descrição | Propósito no Projeto |
| :--- | :--- | :--- |
| **Go (Golang)** | Linguagem de programação principal. | Alta performance e concorrência nativa para o servidor HTTP. |
| **Arquitetura RESTful** | Padrões de design para a comunicação da API. | Define *endpoints* claros e utiliza métodos HTTP padrão para operações CRUD. |
| **SQLite** | Banco de dados embutido e baseado em arquivo. | Persistência de dados leve e sem necessidade de servidores de banco de dados externos. |
| **Pacote SQL/ORM** | Ex: `gorm` ou `database/sql` + `mattn/go-sqlite3`. | Gerenciamento de conexões e mapeamento de objetos para o banco de dados. |

---

## 🛠️ Estrutura de Dados do Estudante

Os dados de cada estudante são persistidos no banco de dados com a seguinte estrutura:

| Campo | Tipo | Descrição |
| :--- | :--- | :--- |
| **Nome** | `string` | Nome completo do estudante. |
| **CPF** | `int` | Cadastro de Pessoa Física. |
| **E-Mail** | `string` | Endereço de e-mail. |
| **Idade** | `int` | Idade do estudante. |
| **Ativo** | `bool` | Status de atividade do estudante (`true` para ativo, `false` para inativo). |

---

## ✨ Funcionalidades (Endpoints RESTful)

A API expõe os seguintes *endpoints* para gerenciar o cadastro de estudantes:

| Método HTTP | Endpoint | Descrição |
| :--- | :--- | :--- |
| **`GET`** | `/students` | **Lista** todos os estudantes (ativos e inativos). |
| **`GET`** | `/students/{id}` | **Obtém** as informações de um estudante específico pelo seu ID. |
| **`GET`** | `/students?active=<true/false>` | **Filtra** a lista de estudantes por status de ativo (`true`) ou inativo (`false`). |
| **`POST`** | `/students` | **Cria** um novo estudante no cadastro, enviando o objeto no corpo da requisição. |
| **`PUT`** | `/students/{id}` | **Atualiza** as informações de um estudante existente, baseado no ID. |
| **`DELETE`** | `/students/{id}` | **Exclui** (ou marca como inativo) um estudante pelo seu ID. |

---

## ⚙️ Como Executar o Projeto

### 1. Pré-requisitos

Certifique-se de ter os seguintes softwares instalados:

* **Golang:** Versão **1.18 ou superior**.
* **Git:** Para clonar o repositório.

### 2. Clonar e Instalar Dependências

Abra seu terminal e siga os passos:

```bash
git clone [https://github.com/danrodsg/api-students.git](https://github.com/danrodsg/api-students.git)
cd api-estudantes
go mod tidy

### 3. Executar a API

O arquivo do banco de dados SQLite será criado automaticamente na primeira execução

```bash
go run main.go

A API estará rodando em http://localhost:8080.

### 🧪 Teste Rápido

Você pode verificar se a API está no ar fazendo uma requisição GET com o curl:

``bash
curl -X GET http://localhost:8080/students


