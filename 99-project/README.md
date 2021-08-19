-- motivo
    vamos criar um webservice para cadastro de clientes
    vamos criar um webservice para cadastro de contas
    vamos criar um webservice para controle de transações

-- endpoints
    -- bancodofernando.com/client
        POST - criar
        GET - buscar
        * DELETE - remover --> não pode
        PATCH - atualizar
    -- bancodofernando.com/account
        POST - criar
        GET - buscar
        PATCH - atualizar
    -- bancodofernando.com/transaction
        POST - criar
        GET - buscar

-- regras de negocio
    
    - usuario
        - nome do cliente não pode ser vazio
        - numero do documento do cliente tem que ser valido
    
    - conta
        - numero da conta não pode ser vazio
    
    - transação
        - somente aceitar contas validas
        - transações não podem exceder o limite
            - se exceder o limit retornar erro HTTP.STATUS 400
                - cliente sem limite

-- DER --> modelo de entidades
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
        - value
        - date
        - destinyAccount

-- MER -> modelo de banco de dados
    -- user
        - ID
        - name
        - documentNumber
        - created_at
        - updated_at
    
   -- user_account
        - user_ID
        - account_ID
    
   -- account
        - ID
        - accountType
        - accountNumber
        - amount
        - limit
        - created_at
        - updated_at

   -- transaction
      - ID
      - transactionType
      - value
      - date
      - destinyAccount -> account
      - created_at
      - updated_at

   -- account_transaction
      - account_ID
      - transaction_ID