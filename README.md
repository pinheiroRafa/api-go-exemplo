# API REST GOLANG

Esse é um repositório para trazer o básico de uma implementação REST em GOLANG.
Nesse repositório tem uma API que trata os conceitos:

Acesso a base de dados (POSTGRESQL)
Variáveis de ambiente 
Separação de códigos em arquivos
Separação de funções globais
Uso de Middlewares
Criação e validação de tokens JWT
Traduções de respostas de erro

## Para rodar 

Ter o go instalado na máquina, no exemplo usado a versão 1.20
Ter o postgresql instalado na máquina
Rodar os scripts de inicialização da pasta scripts no banco

Para acesso nos endpoints de status, necessário atualizar na base, o seu usuário criado para status 2(admin)

o Endpoint de criação por padrão cria no status 1

Comando para rodar: 

    go run .

vai abrir na porta especificada no arquivo .env

para acessar o swagger da aplicação:

    http://localhost:port/swagger

### Variáveis de ambiente

criar um arquivo .env na raiz com as variáveis:

`
PORT=8080
USER_DB="usuariodb"
HOST_DB="localhost"
PORT_DB="5432"
PASS_DB="senhadb"
NAME_DB="estudos"
JWT_KEY="golangpwtkey"
`

## Idioma

Para trocar o idioma das respostas de erros e tratativas, passar o header: 

    Content-Language: 'pt-Br' | 'en'

por padrão responde em português