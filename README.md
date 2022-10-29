# CactusBank - Monolítico

CactusBank e uma API desenvolvida para o trabalho de sistemas distribuídos. Onde a ideia e simular transferências de pix, utilizando métodos e padrões de projetos para o desenvolvimento de suas funcionalidades, a fim de aplicar os conteúdos visto em sala de aula.

A aplicação desenvolvida não está totalmente completa, pois a primeiro tempo, foi necessário apenas realizar a utilização de técnicas ensinado em sala de aula, para determinar um padrão de desenvolvimento que seja escalável, a fim de futuramente realizar a distribuição de responsabilidades e ter um baixo acoplamento em relação a dependências, tornando a aplicação flexível e de fácil manutenibilidade.


# Linguagem de programação

A linguagem de programação utilizada foi o GoLang, desenvolvida pela google e lançada em 2009, é uma linguagem focada na produtividade e na programação concorrente, ela foi criada com soluções para problemas da computação mais relacionado ao sistema operacional em que as linguagens da época não tinha uma solução pré-definida, por trás da criação dessa linguagem estão pessoas relevantes que contribuíram para projetos de destaques no ramo da computação, tais como, Kenneth Thompson, Robert Griesemer, Rob Pike.

# Padrões de projeto

Utilizamos o clássico padrão MVC para separar as responsabilidades das "classes" do projeto, além de manter a preocupação constante com o código fonte, de modo que seja possível ter um baixo acoplamento na utilização de seus objetos e preservar o código para que fique limpo e seguindo padrão de desenvolvimento definido pelo arquiteto do software (Pedro Ítalo).

# Banco de Dados

No armazenamento de dados foi utilizado o banco de dados relacional MySql em sua última versão, porém a forma como realizamos a conexão e utilizamos os códigos SQL é totalmente atribuído a uma biblioteca chamada de GORM (Go Object-Relational Mapping), ela implementa as especificações de um ORM (Object-Relational Mapping), um padrão de desenvolvimento para integração e manipulação do banco de dados em que o desenvolvedor não precisa se preocupar com códigos SQL, deixando a aplicação totalmente responsável por gerar esses códigos, o que chamamos de IoC (Inversion of Control).

# API REST

Esse tipo de API, REST (Representational State Transfer), é utilizado para criação de endpoints de acesso que permitam a interação com determinada funcionalidade da aplicação através do protocolo HTTP. Ao realizar as implementações utilizando suas especificações podemos desenvolver aplicações com alto nível de maturidade, simplicidade, manutenibilidade e segurança. 

Podemos também avaliar se nossa API está de um certo modo madura o suficiente, utilizando um modelo desenvolvido por Richardson que é utilizado para medir em níveis a [maturidade de uma API Rest](https://www.brunobrito.net.br/richardson-maturity-model/). 

# Integração contínua e Entrega contínua 
Nessa seção será explicado e apresentado as ferramentas utilizadas para realizar o "build" e testes da aplicação durante o processo de desenvolvimento.

## Docker
Docker é uma ferramenta utilizada para o gerenciamento de containers, muito popular no mercado, com ele é possível criar diversos ambientes para cada contexto em que você precise utilizar sua aplicação.

O Docker já é capaz de realizar o CI (Continous Integration) e o CD (Continous Delivery) por isso não foi necessário a utilização de uma ferramenta apenas para realizar o CI da aplicação como o Jenkins, Travis CI ou qualquer outro, mas em um cenário de vários ambientes a utilização de uma ferramenta de CI seria utilizada, como o Jenkins, assim ele iria realizar o processo até gerar um artefato (Imagem Docker) e iria realizar o deploy dentro do container ou criando um novo container.