const express = require('express');
const cors = require('cors');

const app = express();
app.use(cors());

app.get('/texto', (req, res) => {
  res.json({ mensagem: 'Aplicação 2 - Texto fixo' });
});

app.get('/horario', (req, res) => {
  const dataAtual = new Date();
  const horaFormatada = `${dataAtual.getHours()}:${dataAtual.getMinutes()}:${dataAtual.getSeconds()} - ${dataAtual.getDate()}/${dataAtual.getMonth() + 1}/${dataAtual.getFullYear()}`;
  res.json({ horario: horaFormatada });
});

const PORT = 5001;
app.listen(PORT, '0.0.0.0', () => {
  console.log(`Servidor rodando na porta ${PORT}`);
});

