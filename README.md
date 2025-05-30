# aws-go-lambda

Estrutura simples de uma AWS Lambda escrita em Go, com provisionamento de recursos via Terraform.

## Funcionalidades

- Lambda em Go pronta para deploy
- Provisionamento de infraestrutura com Terraform
- Makefile para build do binário

## Como usar

1. **Build do binário**

    ```sh
    make build
    ```

2. **Provisionar recursos na AWS**

    ```sh
    terraform init
    terraform apply
    ```

3. **Deploy da Lambda**

    Siga as instruções do Makefile e do Terraform para deploy.

## Requisitos

- Go instalado
- Terraform instalado
- AWS CLI configurado

## Estrutura

- `main.go`: Código da Lambda
- `Makefile`: Comandos para build e deploy
- `terraform/`: Scripts de infraestrutura

