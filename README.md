# Desafio DevOps

Este projeto implementa um desafio de DevOps com duas aplicações em linguagens diferentes (Go e Node.js), uma camada de cache configurada com diferentes tempos de expiração e uma infraestrutura fácil de iniciar.

## Estrutura do Projeto

```
desafio-devops/
├── app1-go/           # Aplicação 1 em Python/Flask
│   ├── main.go            # Código da aplicação
│   ├── Dockerfile         # Configuração do container
├── app2-nodejs/           # Aplicação 2 em Node.js
│   ├── app.js             # Código da aplicação
│   ├── Dockerfile         # Configuração do container
│   └── package.json       # Dependências
├── nginx/                 # Configuração do proxy reverso e cache
│   ├── Dockerfile         # Configuração do container
│   └── nginx.conf         # Configuração do Nginx
├── docker-compose.yml     # Orquestração dos serviços
├── iniciar.sh             # Script para iniciar o projeto
└── README.md              # Este arquivo
```

## Requisitos

- Docker
- Docker Compose

## Como Executar

1. Clone o repositório
2. Execute o script de inicialização:

```bash
chmod +x iniciar.sh
chmod +x parar.sh
./iniciar.sh
```

Isso iniciará todos os serviços em containers Docker.

## Serviços Disponíveis

### Aplicação 1 (Go)
- Texto fixo: http://localhost:8081/app1/texto
- Horário atual: http://localhost:8081/app1/horario
- Cache configurado para 10 segundos

### Aplicação 2 (Node.js)
- Texto fixo: http://localhost:8081/app2/texto
- Horário atual: http://localhost:8081/app2/horario
- Cache configurado para 1 minuto

## Verificando o Cache

Para verificar se o cache está funcionando corretamente, observe o cabeçalho `X-Proxy-Cache` nas respostas:
- `MISS`: Conteúdo não estava em cache
- `HIT`: Conteúdo foi servido do cache

Exemplo de verificação com curl:
```bash
curl -I http://localhost:8081/app1/texto
```

## Comandos Úteis

- Ver logs dos serviços: `docker-compose logs -f`
- Parar os serviços: `docker-compose down`
- Reconstruir os serviços: `./parar.sh && ./iniciar.sh`

### Aplicação 1 (Go)
- Linguagem: Go
- Framework: net/http
- Porta: 5000

### Aplicação 2 (Node.js)
- Linguagem: JavaScript
- Framework: Express
- Porta: 5001

### Proxy Reverso e Cache
- Nginx configurado como proxy reverso
- Cache configurado com diferentes tempos de expiração:
  - Aplicação 1: 10 segundos
  - Aplicação 2: 1 minuto

