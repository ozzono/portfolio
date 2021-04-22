# Backend Assessment

Olá! 🖖🏽

Nossa intenção é, através deste (breve) desafio, avaliar a habilidade técnica percebida ao empregar e desenvolver uma solução para o problema aqui descrito.

## Domínio Problema

Uma instituição financeira contratou os serviços da _api/request_ buscando maior **agilidade dos dados** através da metrificação de processos que, até então, não eram _observados_ (apropriadamente). Um dos processos é a solicitação do produto débito automático de empresas parceiras.
A operação é realizada manualmente e vai ser automatizada por este serviço, que vai permitir que outros serviços consumam, de forma livre, de seus eventos operacionais.

## Escopo

## Casos de Uso

1. Autenticação e acesso a plataforma

Um usuário autenticado,

2. solicita uma ativação de débito automático
3. cancela uma solicitação de ativação
4. aprova uma solicitação de ativação
5. rejeita uma solicitação de ativação
6. visualiza uma solicitação

Diagrama do [modelo de eventos](img/model.jpg).

Observações **importantes** sobre o modelo:

- É uma representação do domínio _exclusivamente_.

- Não é mandatório ser modelado usando CQRS nem event-driven.

- Não é mandatório implementar o EmailServer

## Requisitos

Especifica o contexto em que a aplicação será operacionalizada

### Não funcionais

1. 30 empresas parceiras
1. 5000 usuários simultâneos
1. 100 reqs/s

### Funcionais

#### Tecnologias

- implementação: `golang | elixir | python`
- armazenamento: `postgres | mongodb`
- **não-mandatório** broker: `kafka | rabbitmq`

#### Protocolos

- pontos de entrada: `http`
- autenticação: `simple jwt`

#### Padrões

Bonus points:

- arquitetural: `cqrs & hexagonal`
- design: `ddd & solid`
- message bus as stream

### 3rd parties

O uso de bibliotecas externas é **livre**.

### Deployment

A forma como a aplicação será disponibilizada é **livre**. Fica a critério do candidato, por exemplo, usar algum PaaS a fim de reduzir a complexidade bem como utilizar receitas prontas através de ferramentas de automatização e.g. `ansible+dockercompose`.

No entanto, é esperado bom senso na documentação caso sejam usadas soluções @ `localhost`.

## Entrega

A _Release_ 0.1 🚀 consiste na implementação de um servidor web que implementa os casos de uso listados acima respeitando os requisitos funcionais e não funcionais. Fica a critério do desenvolvedor como os testes serão escritos, os scripts de _data migration_, os _schemas_ de entrada e saída da api e todas as outras definições que não foram listadas neste documento.

## Avaliação

Critérios ordenados por ordem de peso decrescente:

1. Correção (_correctness_) da solução

   - a fim de solucionar o [domínio-problema](#domínio-problema)
   - a fim de cumprir os [casos de uso](#casos-de-uso)
   - ao implementar os [requisitos](#requisitos) especificados

1. Testes
1. Organização, documentação e clareza na estruturação do projeto
1. Estilo, legibilidade e simplicidade no código
1. Escolhas e uso de 3rd parties
1. Padrões de segurança

### Bonus points 🏆

1. Teste de stress
1. Boas práticas na modelagem e armazenamento de dados

### Submissão

> Método de submissão do código omitido
