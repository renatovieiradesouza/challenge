FROM node:14-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY app.js .

EXPOSE 5001

CMD ["node", "app.js"]

