# Projeto de Clima e CEP

Este projeto é uma aplicação em Go que recebe um CEP, encontra a cidade correspondente e retorna o clima atual, com a temperatura nas escalas Celsius, Fahrenheit e Kelvin.

## 0. Configuração
### **WeatherAPI**:
- Crie uma conta em [https://www.weatherapi.com/](https://www.weatherapi.com/) para obter uma chave de API gratuita.
- Após criar sua conta, obtenha sua chave no painel de controle da API.
- Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:
- WEATHER_API_KEY=your_weatherapi_key

## 1. Rodando o projeto localmente

```bash
go run main.go
```

#### para testar, utilizar endereço como localhost:8080

## 2. Rodando o Projeto com Docker

### 2.a. Rodar os testes
```bash
docker compose up tests
```

### 2.b. Rodar a aplicação
```bash
 docker compose up app
 ```

#### para testar, utilizar endereço como localhost:8080

## 3. Rodando o Projeto com Google Cloud Run
### Endereço para acessar a API:
```bash
https://lab-cloud-run-uztcvktcmq-uc.a.run.app/
```

## 4. Exemplos de requisições

### 4.a. CEP válido
```bash
<endereço>/80035050
```

### 4.b. CEP inválido
```bash
<endereço>/8003505a
```

### 4.c. CEP inexistente
```bash
<endereço>/12345678
```
