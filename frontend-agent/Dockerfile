FROM node:12.18-alpine
WORKDIR /usr/src/app
COPY package.json ./
COPY package-lock.json ./
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent
#RUN mv node_modules ../
COPY . .
EXPOSE 3000
CMD ["npm", "start"]
