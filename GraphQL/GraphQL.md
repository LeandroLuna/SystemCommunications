# GraphQL

O que é o GraphQL?

GraphQL é uma linguagem de consulta de dados e uma runtime de execução criada pelo Facebook em 2012 e lançada em 2015 como uma alternativa ao modelo tradicional de APIs REST. Em vez de seguir a abordagem RESTful de ter endpoints específicos para cada recurso, o GraphQL permite que os clientes requisitem exatamente os dados que desejam e nada mais. Isso torna as APIs mais eficientes, flexíveis e fáceis de evoluir.

## Principais Conceitos do GraphQL

Aqui estão alguns dos principais conceitos do GraphQL que você precisa conhecer:

### 1. Esquema (Schema)

O esquema é um componente central no GraphQL. É uma descrição do tipo de dados que sua API pode retornar e das operações que podem ser executadas nesses dados. O esquema define os tipos de objetos, as relações entre eles e as operações disponíveis.

### 2. Tipos de Dados (Data Types)

O GraphQL define tipos de dados para representar os objetos que podem ser consultados. Alguns tipos de dados básicos incluem `String`, `Int`, `Float`, `Boolean` e `ID`, além de tipos personalizados que você pode definir.

### 3. Consultas (Queries)

As consultas são usadas pelos clientes para buscar dados no servidor GraphQL. Uma consulta especifica os campos que o cliente deseja obter e pode incluir argumentos para filtrar, paginar ou ordenar os resultados.

Exemplo de consulta GraphQL:
```graphql
query {
  user(id: "123") {
    id
    name
    email
  }
}
```

### 4. Mutações (Mutations)

As mutações são usadas para modificar dados no servidor GraphQL. Elas são semelhantes às consultas, mas são usadas para criar, atualizar ou excluir informações.

Exemplo de mutação GraphQL:
```graphql
mutation {
  createUser(input: {
    name: "John Doe",
    email: "john.doe@example.com"
  }) {
    id
    name
    email
  }
}
```
### 5. Resolvers

Os resolvers são funções responsáveis por fornecer os dados para os campos especificados no esquema. Cada campo do esquema deve ter um resolver associado que retorna o valor desse campo quando é consultado.

Quando uma consulta é enviada ao servidor GraphQL, o mecanismo de execução do GraphQL resolve os campos consultados, invocando os resolvers correspondentes para obter os dados reais. Os resolvers são responsáveis por buscar os dados de uma fonte de dados, como um banco de dados, um serviço externo ou qualquer outra fonte.

### 6. Tipos de Relacionamento (Relationship Types)

O GraphQL permite definir relacionamentos entre os tipos de dados. Por exemplo, você pode ter um tipo `User` que possui uma lista de tarefas (`Todo`) relacionadas a esse usuário.

### 7. Fragments

Os fragments são uma maneira de reutilizar partes de consultas em várias partes do código. Eles ajudam a evitar a repetição de código e a simplificar consultas complexas.

### 8. Diretivas

As diretivas são usadas para modificar a execução da consulta ou para fornecer metadados adicionais. A diretiva mais comum é `@include` e `@skip`, que permite que o cliente inclua ou exclua campos com base em condições.

## Como o GraphQL funciona?

1. O cliente envia uma query GraphQL para o servidor.
2. O servidor analisa a query e valida seu esquema.
3. O servidor executa a query, buscando os dados necessários de suas fontes de dados (por exemplo, bancos de dados, serviços externos, etc.).
4. O servidor retorna os dados solicitados ao cliente em uma estrutura JSON.