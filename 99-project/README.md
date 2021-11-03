- motivo
    
    - vamos criar um webservice para cadastro de clientes
    - vamos criar um webservice para cadastro de contas
    - vamos criar um webservice para controle de transações

- endpoints
    - bancodofernando.com/user
        - POST - criar
        - GET - buscar
        - DELETE - remover --> não pode
        - PATCH - atualizar
    - bancodofernando.com/account
        - POST - criar
        - GET - buscar
        - PATCH - atualizar
    - bancodofernando.com/transaction
        - POST - criar
        - GET - buscar

- regras de negocio
    
    - usuario
        - nome do cliente não pode ser vazio
        - numero do documento do cliente tem que ser valido
    
    - conta
        - numero da conta não pode ser vazio
    
    - transação
      - criar transação
          - somente aceitar contas validas
          - transações não podem exceder o limite
              - se exceder o limit retornar erro HTTP.STATUS 400
                  - cliente sem limite
          - remover ou creditar saldo na conta
      - consultar transações
        - id conta
        - init date >= data
        - end date <= data
        - init date = 10/10/2010 end date = 11/10/2010
          - where transactions.id = id and date >= init date and date <= end date
        - init date = 10/10/2010
          - where transactions.id = id and date >= init date
        - end date = 11/10/2010
          - where transactions.id = id and date <= end date

- DER --> modelo de entidades
    - user
        - name
        - documentNumber
        - accounts[]
    - account 
        - accountType
        - accountNumber
        - amount
        - limit
        - transactions[]
    - transactions
        - transactionType
        - account
        - value
        - date
        *- destinyAccount

- MER -> modelo de banco de dados
    - users
        - ID
        - name
        - documentNumber
        - created_at
        - updated_at

```
CREATE TABLE users (
    id UUID NOT NULL,
    name VARCHAR(255),
    documentNumber TEXT,
    created_at date,
    updated_at date
);
```
    - accounts
         - ID
         - user_ID
         - accountType
         - accountNumber
         - amount
         - limit
         - created_at
         - updated_at

    - transaction
       - ID
       - transactionType
       - value
       - date
       - destinyAccount -> account
       - created_at
       - updated_at

    - account_transaction
       - account_ID
       - transaction_ID



## ARQUITETURA BASEADA EM:

https://echo.labstack.com/cookbook/crud/
https://stackoverflow.com/questions/13319165/go-fmt-on-a-whole-source-tree


// criar
curl -X POST \
-H 'Content-Type: application/json' \
-d '{"name":"Teste 1", "document_number": "212.712.270-45"}' \
localhost:9090/v1/users

// listar todos
curl -X GET \
-H 'Content-Type: application/json' \
localhost:9090/v1/users
