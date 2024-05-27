# Weather App by CEP

## Descrição

Este é um sistema desenvolvido em Go que recebe um CEP, identifica a cidade e retorna o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin). O sistema é implantado no Google Cloud Run.

## Como Executar

### Requisitos

- Docker
- Docker Compose
- Conta no Google Cloud

### Passos para Executar Localmente

1. **Clone o repositório:**

   ```sh
   git clone https://github.com/AnaGFranco/goexpert-temperature-system-by-cep.git
   cd goexpert-temperature-system-by-cep
   ```

2.**Construir as Imagens**: Na raiz do projeto, execute:

   ```
   docker-compose build
   ```

3. **Executar o Docker Compose**: Na raiz do projeto (onde está o arquivo `docker-compose.yml`), execute:

   ```
   docker-compose up
   ```


4. **Acesse a aplicação:**

   A aplicação estará disponível em `http://localhost:8080`.

### Exemplos de Requisição

- **Para obter o clima de um CEP específico:**

   ```sh
   curl -X GET http://localhost:8080/weather/29902555
   ```

### Testando o Serviço no Google Cloud Run

URL do Google Cloud Run. Use essa URL para fazer requisições:

```sh
curl -X GET https://goexpert-temperature-system-by-cep-xhekudbara-uc.a.run.app/weather/29902555
```
