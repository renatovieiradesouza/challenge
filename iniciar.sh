#!/bin/bash

echo "Iniciando o desafio DevOps..."
docker compose up -d

echo ""
echo "Serviços disponíveis:"
echo "- Aplicação 1 (Go): http://localhost:8081/app1/texto e http://localhost:8081/app1/horario"
echo "- Aplicação 2 (Node.js): http://localhost:8081/app2/texto e http://localhost:8081/app2/horario"
echo ""
echo "Para verificar os logs: docker-compose logs -f"
echo "Para parar os serviços: docker-compose down"
