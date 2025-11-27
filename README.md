# api-students
API de cadastro de estudantes desenvolvida em Go (Golang), implementando arquitetura RESTful para operações CRUD completas e persistência de dados.

Rotas:

GET /students - Listar todos os alunos
POST /alunos - Criar aluno
GET /students/:id - Obter informações de um aluno específico
PUT /students/:id - Atualizar aluno
DELETE /students/:id - Excluir aluno
GET /students?active=<true/false> - Lista todos os alunos ativos/inativos

Estrutura dos alunos:

Nome (string)
CPF (int)
E-mail (string)
Idade (int)
Ativo (bool)