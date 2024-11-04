# Zipcode Temperature System API

## Descrição

A **Zipcode Temperature System API** fornece uma interface para consultar a temperatura atual de uma cidade brasileira
com base no CEP. A API usa a [API WeatherAPI](https://www.weatherapi.com/) para obter dados de temperatura
e [ViaCEP](https://viacep.com.br/) para buscar dados de localização.

## Funcionalidades

- Buscar cidade com base no CEP.
- Consultar a temperatura atual da cidade.

## Estrutura de Diretórios

- `internal/server`: Contém o `main.go`, onde a aplicação é inicializada.
- `internal/service`: Contém a lógica de negócio para acessar as APIs ViaCEP e WeatherAPI.
- `docs`: Contém a documentação Swagger gerada automaticamente.

## Pré-requisitos

- **Docker**.
- Conta e chave de API na [WeatherAPI](https://www.weatherapi.com/).

## Configuração

1. Crie um arquivo `.env` na raiz do projeto e adicione sua chave de API para o WeatherAPI:

    ```dotenv
    WEATHER_API_KEY=your_weather_api_key
    ```

## Como Usar

1. **Construa e execute a aplicação usando Docker Compose**:

    ```bash
    docker-compose up --build
    ```

2. A aplicação estará disponível em `http://localhost:8080`.

3. **Endpoints Disponíveis**:
    - **GET** `/temperature/{cep}`: Retorna a temperatura atual de uma cidade baseada no CEP.
    - Exemplos de uso:
        - `http://localhost:8000/temperature/01001000`

## Documentação da API

Após iniciar a aplicação, você pode acessar a documentação Swagger em:

http://localhost:8000/swagger/index.html

## Testes

Para executar os testes, utilize o comando abaixo:

```bash
docker-compose run tests 
```

Isso executará os testes de cobertura e exibirá os resultados no terminal.


## Aplicação disponível no Google Cloud:
https://zipcode-temperature-system-ggv2ooifja-uc.a.run.app/swagger/index.html#/default
