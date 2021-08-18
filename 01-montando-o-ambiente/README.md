	- vm para o ambiente
		- instalar o vmware workstation pro
            - download em: [workstation-pro-evaluation.html](https://www.vmware.com/products/workstation-pro/workstation-pro-evaluation.html)
        
		- baixar a iso do linux mint 
            - download em [linuxmint-20.2-cinnamon-64bit.iso](http://mirror.ufscar.br/mint-cd/stable/20.2/linuxmint-20.2-cinnamon-64bit.iso)

    - instalando o linux mint
        - instruções da criação da vm
            - no vmware apert crtl + n
            - selecione a iso do linux mint
            - sistema operacional selecione linux
                - em escolha o sistema pode ser ubuntu x64
            - nome do sistema WORKSPACE
            - espaço 20GB
        - instruções da instalação do linux
            - selecione no boot a primeira opção 
            - use o icone install mint no desktop
            - avance todas as etapas mantendo o default
            - o usuario pode ser aula
            - a senha 1234 ( nota essa senha sera pedida varias vezes por isto de tornar ela facil para as aulas )
    

## instalando aplicações

- git no bash ( terminal ) rodar: 
    > $ apt update
    
    > $ apt install git

    - create github account

      - https://github.com/signup?ref_cta=Sign+up&ref_loc=header+logged+out&ref_page=%2F&source=header-home

    - autenticar git
      - 

    > $ ssh-keygen -t ed25519 -C "your_email@example.com"

    > $ sudo apt-get update

    > $ sudo apt-get install xclip

    - copiar a chave ssh para o clipboard do linux

    > $ xclip -selection clipboard < ~/.ssh/id_ed25519.pub 

    - adicionar a chave ssh nomeada aula no link [https://github.com/settings/keys](https://github.com/settings/keys)
 
    
- jetbrains toolbox no bash ( terminal ) rodar: 
    > $ wget https://download.jetbrains.com/toolbox/jetbrains-toolbox-1.21.9547.tar.gz

    > $ tar -zxvf jetbrains-toolbox-1.21.9547.tar.gz

    > $ ./jetbrains-toolbox-1.21.9547/jetbrains-toolbox 

    > baixar pelo toolbox o goland

    -- FIXME terminar doc de import do projeto
    > importar o projeto
      
      > enable go modules
  
- docker-compose
  
  > $ sudo apt install docker-compose
  > $ sudo usermod -aG docker $USER

## fazer fork do projeto go-traing

- no projeto [go-training](https://github.com/templarfelix/go-training) clicke em fork e selecione sua conta e avançe até o projeto ser copiado para sua conta pessoal
  
## clonar repositorio go-training

- no bash ( terminal ) rodar: 

    > $ mkdir workspace

    > $ cd workspace 

    > $ git clone git@github.com:[SEU_USER_GITHUB]/go-training.git

## atualizar projeto go-training com bases nos fontes do templarfelix

- no bash ( terminal ) na pasta workspace/go-training rodar: 

    > $ git remote add templarfelix git@github.com:templarfelix/go-training.git

    > $ git fetch templarfelix

    > $ git merge templarfelix/main

    > $
  > 