# Checklist para Desenvolvimento do Sistema de CEP e Clima

## 1. Planejamento e Configuração do Ambiente
- [x] Criar o servidor HTTP básico em Go.
- [x] Configurar Dockerfile para build e execução no Google Cloud Run.
- [X] Configurar `docker-compose` para facilitar testes locais.

---

## 2. Desenvolvimento da Aplicação
### 2.1. Endpoint para CEP
- [x] Criar endpoint HTTP básico para receber o CEP.
- [X] Adicionar validação para o CEP (8 dígitos e apenas números).
- [ ] Tratar cenários de erro:
    - [X] Retornar HTTP 422 para CEP inválido.
    - [ ] Retornar HTTP 404 caso o CEP não seja encontrado na API ViaCEP.

### 2.2. Consulta à API ViaCEP
- [ ] Implementar função para consultar a API ViaCEP.
- [ ] Extrair o nome da cidade e estado a partir da resposta da API.
- [ ] Tratar erros da API ViaCEP.

### 2.3. Consulta à API WeatherAPI
- [ ] Implementar função para consultar a API WeatherAPI com base na localização obtida.
- [ ] Obter temperatura em graus Celsius.
- [ ] Tratar erros da API WeatherAPI.

### 2.4. Conversões de Temperatura
- [ ] Implementar função para converter de Celsius para Fahrenheit.
- [ ] Implementar função para converter de Celsius para Kelvin.

### 2.5. Resposta da API
- [ ] Formatar resposta JSON para sucesso:
    - Exemplo: { "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.65 }
- [ ] Implementar retorno de erros apropriados:
    - HTTP 422: {"message": "invalid zipcode"}
    - HTTP 404: {"message": "can not find zipcode"}

---

## 3. Testes Automatizados
- [ ] Implementar testes unitários para as funções:
    - Validação do CEP.
    - Consulta à API ViaCEP.
    - Consulta à API WeatherAPI.
    - Conversões de temperatura.
- [ ] Implementar testes de integração para o endpoint HTTP.

---

## 4. Deploy no Google Cloud Run
- [x] Configurar Dockerfile para produção.
- [ ] Criar conta no Google Cloud e configurar o projeto.
- [ ] Realizar o deploy no Google Cloud Run (Free Tier).
- [ ] Testar o endpoint no ambiente de produção.

---

## 5. Documentação
- [ ] Criar um `README.md` com instruções para rodar o projeto:
    - Localmente.
    - Com Docker.
    - Endpoint de produção no Google Cloud Run.
- [ ] Documentar as dependências e como obter as chaves de API.

---

## 6. Entrega
- [ ] Compartilhar o repositório com o código-fonte.
- [ ] Informar o link do endpoint publicado no Google Cloud Run.
- [ ] Adicionar exemplos de requisições no `README.md`.
