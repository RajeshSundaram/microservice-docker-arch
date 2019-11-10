FROM node:lts-alpine AS builder
WORKDIR /build
COPY package*.json ./
RUN npm i
COPY . .
RUN npm run build

FROM node:lts-alpine
WORKDIR /app
RUN npm install -g http-server
COPY --from=builder /build/dist /app
EXPOSE 8080
CMD ["http-server", "."]
