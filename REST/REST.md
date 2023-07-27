# REST

REST (Representational State Transfer) é um estilo arquitetural para projetar sistemas de software distribuídos, mais comumente usado para projetar serviços web. Foi proposto por Roy Fielding em sua tese de doutorado em 2000 e rapidamente se tornou uma abordagem amplamente adotada para o desenvolvimento de APIs (Application Programming Interfaces) na World Wide Web.

O REST é baseado em alguns princípios fundamentais:

1. Recursos (Resources): Os recursos são as entidades que o sistema expõe, como objetos, dados ou serviços. Em uma API RESTful, cada recurso é identificado por um URI (Uniform Resource Identifier).

2. Verbos HTTP (HTTP Verbs): O REST utiliza os principais verbos HTTP, como GET, POST, PUT e DELETE, para definir as ações que podem ser realizadas nos recursos.

    - GET: É usado para obter um recurso existente.
    - POST: É usado para criar um novo recurso.
    - PUT: É usado para atualizar um recurso existente ou criar um novo recurso se ele não existir.
    - DELETE: É usado para excluir um recurso.

3. Representações (Representations): Os recursos podem ter várias representações, como JSON, XML ou texto, que são usadas para transferir informações entre cliente e servidor.

4. Sem estado (Stateless): Cada solicitação para o servidor deve conter todas as informações necessárias para compreender e processar a solicitação. O servidor não mantém informações de estado sobre o cliente entre as solicitações. Isso significa que cada solicitação é independente das outras.

5. Interface uniforme (Uniform Interface): O REST enfatiza uma interface uniforme entre componentes, promovendo a simplicidade e a independência dos sistemas envolvidos. Uma interface uniforme facilita a comunicação entre diferentes partes do sistema.

Vantagens do REST:

- É amplamente adotado e compreendido, facilitando a interoperabilidade entre sistemas.

- Utiliza os padrões HTTP existentes, o que torna sua implementação mais fácil em várias plataformas.

- É escalável e pode ser usado para criar sistemas distribuídos em grande escala.

- Permite o uso de diferentes formatos de dados, como JSON e XML, para troca de informações.

## Niveis de Maturidade (Richardson Maturity Model)

O Richardson Maturity Model, nomeado em homenagem a Leonard Richardson, é uma classificação que descreve diferentes níveis de maturidade para projetos de APIs (Application Programming Interfaces) RESTful. Ele oferece uma maneira de avaliar o quão bem uma API RESTful foi projetada e implementada com base em seus princípios fundamentais. O modelo possui quatro níveis de maturidade, do nível 0 ao nível 3.

Vamos ver cada nível de maturidade em detalhes:

1. Nível 0 - The Swamp of POX:

Nesse nível, a API é basicamente apenas uma coleção de endpoints que mapeiam operações HTTP diretas para funções de negócios. A API não segue os princípios RESTful e usa HTTP apenas como um transporte para outras operações. Normalmente, a comunicação é baseada em chamadas de procedimento remoto (RPC) usando HTTP POST, e não há utilização adequada dos métodos HTTP como GET, POST, PUT e DELETE.

2. Nível 1 - Resources:

Neste nível, a API começa a utilizar os métodos HTTP corretamente, como GET, POST, PUT e DELETE, para ações específicas sobre recursos. No entanto, a API ainda não aproveita os recursos no sentido pleno. Em geral, é possível encontrar várias URLs que apontam para recursos diferentes, mas as URLs não seguem uma estrutura consistente e significativa.

3. Nível 2 - HTTP Verbs:

No nível 2, a API utiliza corretamente os métodos HTTP para operações específicas em recursos. Além disso, ela começa a adotar uma estrutura mais consistente para suas URLs, tornando-as mais legíveis e compreensíveis. A API também pode começar a usar códigos de status HTTP apropriados para indicar o resultado de uma solicitação (por exemplo, 200 OK, 404 Not Found, 201 Created).

4. Nível 3 - Hypermedia Controls:

No nível 3, também conhecido como "HATEOAS" (Hypermedia as the Engine of Application State), a API atinge o nível mais alto de maturidade. Nesse nível, a API fornece links hipermídia (hiperlinks) nos dados de resposta, permitindo que o cliente descubra e navegue de forma autônoma pela API sem depender de conhecimento prévio sobre suas URLs. A API se torna verdadeiramente autoexplicativa e dinâmica, tornando a integração com ela mais flexível e resiliente às mudanças.

O objetivo final é alcançar o nível 3, permitindo que a API seja altamente flexível, autoexplicativa e fácil de evoluir. No entanto, nem todas as APIs precisam ou devem alcançar o nível 3, pois a complexidade adicional pode não ser justificada em todos os cenários. A escolha do nível de maturidade dependerá dos requisitos e objetivos do projeto.

## REST vs Restful

A diferença entre "REST" e "RESTful" está relacionada ao uso da terminologia e ao nível de adesão aos princípios do estilo arquitetural REST.

- REST (Representational State Transfer):REST é um estilo arquitetural para projetar sistemas de software distribuídos, mais comumente usado para projetar serviços web. Como mencionado anteriormente, o REST é baseado em alguns princípios fundamentais, como recursos identificados por URIs, utilização dos métodos HTTP corretamente (GET, POST, PUT, DELETE) e uma interface uniforme, entre outros.

- RESTful: "RESTful" é um termo usado para descrever uma API que segue os princípios e as diretrizes do estilo arquitetural REST. Quando dizemos que uma API é "RESTful", estamos enfatizando que ela foi projetada e implementada de acordo com as melhores práticas e padrões REST.

## HAL, Collection+JSON e Siren

HAL, Collection+JSON e Siren são formatos e padrões que podem ser utilizados para melhorar a semântica, a estrutura e a navegação de APIs REST. Cada um deles oferece uma abordagem diferente para alcançar esses objetivos:

1. HAL (Hypertext Application Language): HAL é um formato que segue o princípio do HATEOAS (Hypermedia as the Engine of Application State), que é um dos princípios fundamentais do REST. Ele permite que a API inclua links hipermídia nos dados de resposta, tornando a API autoexplicativa e permitindo que o cliente descubra e navegue pela API de forma autônoma.

2. Collection+JSON: Collection+JSON é um formato de mídia projetado para simplificar o tratamento de coleções de recursos em APIs RESTful. Ele fornece uma estrutura padrão para representar dados em uma coleção e inclui recursos, URLs e itens de dados associados a um conjunto de links e consultas.

3. Siren: Siren é outro formato que segue o princípio HATEOAS e foi projetado para aumentar a semântica das APIs RESTful. Ele utiliza um formato JSON mais estruturado para representar recursos e relacionamentos. Além dos links de hipermídia, o Siren também suporta entidades e ações, o que torna a descrição dos recursos mais rica e detalhada.

Esses formatos têm como objetivo fornecer uma maneira padronizada de descrever e navegar em APIs RESTful. Eles podem ser úteis em casos em que a simplicidade e a descoberta autônoma da API são fundamentais para melhorar a usabilidade e a experiência do desenvolvedor. A escolha entre HAL, Collection+JSON, Siren ou outros formatos dependerá do contexto específico e dos requisitos do projeto. É importante verificar a compatibilidade com as ferramentas e bibliotecas que serão utilizadas pelo cliente e pelo servidor para processar esses formatos.

## HTTP Method Negotiation e Content Negotiation

1. HTTP Method Negotiation: HTTP Method Negotiation refere-se à capacidade de um cliente e um servidor negociarem qual método HTTP será utilizado em uma determinada requisição. Os principais métodos HTTP são GET, POST, PUT, DELETE, entre outros. O cliente pode indicar ao servidor qual método ele deseja usar, e o servidor pode aceitar ou rejeitar a solicitação com base nas permissões, configurações e lógica de negócios da API. A negociação de método HTTP é essencial para garantir que os clientes utilizem as ações corretas em cada contexto da API.

Por exemplo, um cliente pode fazer uma requisição com o método GET para obter informações de um recurso, enquanto outra requisição com o método POST pode ser utilizada para criar um novo recurso no servidor. A negociação de método HTTP permite que o servidor entenda a intenção do cliente em cada solicitação.

2. Content Negotiation: Content Negotiation, também conhecida como Content-Type Negotiation, refere-se à capacidade de um cliente e um servidor negociarem o formato de representação dos dados que serão utilizados em uma resposta. Quando o cliente faz uma requisição, ele pode especificar quais tipos de mídia ele é capaz de entender através do cabeçalho "Accept". O servidor, por sua vez, escolhe uma das opções fornecidas pelo cliente ou retorna um código de status "406 Not Acceptable" se não puder atender nenhuma das opções.

Da mesma forma, quando o cliente envia dados em uma requisição (por exemplo, em uma requisição POST ou PUT), ele pode especificar o formato dos dados enviados através do cabeçalho "Content-Type". O servidor utiliza essa informação para processar os dados corretamente.

A Content Negotiation permite que a API seja flexível e possa fornecer respostas em diferentes formatos, como JSON, XML, HTML, entre outros, de acordo com as preferências do cliente. Isso facilita a interoperabilidade e a evolução da API ao longo do tempo, permitindo que novos formatos sejam adicionados sem afetar a compatibilidade com os clientes existentes.

### Content-type Negotiation

Content-Type Negotiation, também conhecido como negociação de tipo de conteúdo, é uma forma de Content Negotiation no contexto das requisições HTTP. Ele permite que o cliente e o servidor negociem o tipo de mídia (formato de conteúdo) que será utilizado para representar os dados na resposta de uma requisição.

Quando o cliente faz uma solicitação HTTP, ele pode incluir um cabeçalho chamado "Accept" para indicar ao servidor quais tipos de mídia ele é capaz de entender na resposta. O cabeçalho "Accept" contém um ou mais tipos de mídia (MIME types) separados por vírgulas, listando os formatos de conteúdo que o cliente é capaz de processar. 

Por exemplo, um cabeçalho "Accept" pode ser assim:

```
Accept: application/json, application/xml;q=0.9, text/html;q=0.8
```

Neste exemplo, o cliente informa que prefere receber a resposta em JSON ("application/json"), mas também pode aceitar XML ("application/xml") com uma qualidade (q) de 0.9 (90%), ou HTML ("text/html") com uma qualidade de 0.8 (80%).

O servidor, ao receber a requisição, verifica os tipos de mídia especificados no cabeçalho "Accept" e escolhe o formato de conteúdo mais apropriado que pode atender ao pedido do cliente. Se o servidor não puder fornecer nenhum dos tipos de mídia especificados no cabeçalho "Accept", ele poderá retornar um código de status "406 Not Acceptable".

A Content-Type Negotiation permite que a API forneça respostas em diferentes formatos, de acordo com as preferências do cliente. Isso torna a API mais flexível e adaptável a diferentes tipos de clientes, como aplicativos móveis, navegadores web ou outros sistemas, proporcionando uma melhor experiência para os usuários finais. Além disso, a negociação do tipo de conteúdo é útil para evoluir a API ao longo do tempo, pois novos formatos de conteúdo podem ser adicionados sem afetar a compatibilidade com clientes existentes.