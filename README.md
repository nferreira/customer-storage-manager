# Customer Storage Manager

O customer storage mananger faz parte do FastData, ele será responsável por toda comunicação entre os serviços com PMID...

O customer storage manager está abaixo dos pacotes: 

```bash
	- adapter
	- adapter-fiber 
	- app
```

Tudo que iremos fazer é usar o customer-storage-manager para o desenvolvimento de nossa api, e abaixo está os passos que poderíamos fazer para entender e executar o nosso customer-storage-manager.

O customer-storage-manager poderá ser executado localmente usando docker, docker-compose ou no k8s.

Quando tiver rodando locamente você poderá subir o mongo em sua máquina de diversas formas, porém deixamos bonitinho ai no docker-compose.yaml a configuração bonitinha para que possam usar sem muita complicação, podendo executar quando tiver debugando, simulando, testando etc.. 

#### Subindo o Mongo

Para subir somente o mongo basta rodar o comando abaixo:
```bash
$ docker-compose up -d mongo
$ docker-compose ps
```
Após subir o mongo, ele está configurado com user, password tudo com dados default, o nosso customer-storage-manager já está configurado também como default e irá fazer a comunicação nativa quando subir o customer-storage-manager.

Para visualizar o mongo você poderá usar qualquer aplicação de UI como MongoDB compass ou usar o bash para acessar o mongo.

Para acessar o mongo vai bash basta fazer o seguinte:
```bash
$ docker exec -it mongodb bash
# mongo -uroot -pabc123

> use customerDB
> db.customers.find()
```
Agora é brincar com os comandos do mongo e ir para cima!!

#### Rodando Localmente

Faça um clone do projeto para sua máquina.
```bash
$ git clone https://github.com/nferreira/customer-storage-manager
```

Não podemos esquecer do nosso ~/.netrc para autenticação com gitlab.
Vou deixar aqui o formato do arquivo como ele é:
```bash
$ cat ~/.netrc
machine gitlab.engdb.com.br
login <your user>
password <your password>
```

Este arquivo é que irá permitir baixarmos corretamente nossos pacotes do gitblab para que nosso customer-storage-manager execute corretamente.

#### Executando nosso customer
Para rodar localmente temos algumas 
Para executa-lo basta rodar customer-storage-manager podemos usar o make.

```bash
$ make run
       _______ __                 
  ____ / ____(_) /_  ___  _____   HOST   0.0.0.0  OS    LINUX
_____ / /_  / / __ \/ _ \/ ___/   PORT   8080     CORES 4
  __ / __/ / / /_/ /  __/ /       TLS    FALSE    MEM   15.6G
    /_/   /_/_.___/\___/_/1.12.2  ROUTES 7        PPID  8627
```

Irá aparecer diversas variáveis de ambiene, elas podem ser setadas em seu ambiente para que possa fazer seus testes e configurações conforme a necessidade.
Elas estão sendo setadas por um valor default. Até o momento temos variaveis do FIBER que é o framework que estamos usando e do MONGO.

```bash
Looking for ENV: [FIBER_CONCURRENCY], but could not find the value, therefore returning default value [262144]
Looking for ENV: [FIBER_READ_TIMEOUT], but could not find the value, therefore returning default value [3m0s]
Looking for ENV: [FIBER_WRITER_TIMEOUT], but could not find the value, therefore returning default value [3m0s]
Looking for ENV: [FIBER_READ_BUFFER], but could not find the value, therefore returning default value [4096]
Looking for ENV: [FIBER_WRITE_BUFFER], but could not find the value, therefore returning default value [4096]
Looking for ENV: [FIBER_USE_COMPRESSION], but could not find the value, therefore returning default value [false]
Looking for ENV: [FIND_BY_SOCIAL_SECURITY_NUMBER_TIMEOUT], but could not find the value, therefore returning default value [5s]
Looking for ENV: [FIND_BY_ID_TIMEOUT], but could not find the value, therefore returning default value [5s]
Looking for ENV: [CREATE_CUSTOMER_TIMEOUT], but could not find the value, therefore returning default value [5s]
Looking for ENV: [UPDATE_CUSTOMER_TIMEOUT], but could not find the value, therefore returning default value [5s]
Looking for ENV: [DELETE_CUSTOMER_BY_ID_TIMEOUT], but could not find the value, therefore returning default value [5s]
Looking for ENV: [FIBER_HTTP_PORT], but could not find the value, therefore returning default value [8080]
Looking for ENV: [MONGODB_SCHEMA], but could not find the value, therefore returning default value [mongodb]
Looking for ENV: [MONGODB_URI], but could not find the value, therefore returning default value [localhost:27017]
Looking for ENV: [MONGODB_USERNAME], but could not find the value, therefore returning default value [root]
Looking for ENV: [MONGODB_PASSWORD], but could not find the value, therefore returning default value [abc123]
Looking for ENV: [MONGODB_DATABASE], but could not find the value, therefore returning default value [customerDB]
Looking for ENV: [MONGODB_COLLECTION], but could not find the value, therefore returning default value [customers]
Looking for ENV: [MONGODB_PING_TIMEOUT], but could not find the value, therefore returning default value [2s]
Looking for ENV: [MONGODB_READ_TIMEOUT], but could not find the value, therefore returning default value [5s]
Looking for ENV: [MONGODB_WRITE_TIMEOUT], but could not find the value, therefore returning default value [5s]
Looking for ENV: [MONGODB_OPTIONS], but could not find the value, ... [authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false]
Looking for ENV: [MONGODB_CONNECTION_TIMEOUT_IN_SECONDS], but could not find the value, therefore returning default value [10ns]
```

Caso queira compilar basta fazer:
```bash
$ make build
```
Ele irá gerar um binário dentro do diretório cmd ou seja cmd/customer-storage-manager
para executá-lo basta fazer isto aqui.
```bash
$ cmd/customer-storage-manager
        _______ __                 
  ____ / ____(_) /_  ___  _____   HOST   0.0.0.0  OS    LINUX
_____ / /_  / / __ \/ _ \/ ___/   PORT   8080     CORES 4
  __ / __/ / / /_/ /  __/ /       TLS    FALSE    MEM   15.6G
    /_/   /_/_.___/\___/_/1.12.2  ROUTES 7        PPID  3371

```

#### Rodando em docker localmente
Para rodar o customers-storage-manager utilizando o docker basta fazer build.
```bash

$ docker build --no-cache -f Dockerfile -t gcr.io/tim-pmid-dev/customer-storage-manager:latest .

```
Para testar basta rodar docker run.
```bash
$ docker run -p 8080:8080 --rm --name customer-storage  gcr.io/tim-pmid-dev/customer-storage-manager:latest
```

### Rodando em docker-compose

Para rodar nosso docker compose bonitinho é importante que tenha gerado a imagem.
```bash

$ docker build --no-cache -f Dockerfile -t gcr.io/tim-pmid-dev/customer-storage-manager:latest .

```

Com a imagem gerada, podemos agora subir nossos serviços completinho usando docker-compose. 
Temos configurado o customer-storage-manager e o mongo.
```bash

$ docker-compose up -d --build
$ docker-compose ps

```

#### Subindo para Register Google cloud

Agora vamoms configurar nosso ambiente do Google Cloud, para interargirmos com os serviços da cloud.
Antes de subirmos nosso docker
Para que este comando funcione você precisa está logado no google cloud com sua conta, gcloud init caso não tenha feito ainda.

Rode o comando abaixo para certificar que está tudo certinho:
```bash

$ gcloud auth configure-docker 

```

```bash

$ docker push gcr.io/tim-pmid-dev/customer-storage-manager

```

#### Subindo ara k8s 

Todo nosso projeto está em gcr.io/tim-pmid-dev no Register e no k8s estamos na região southamerica-east1 zone southamerica-east1-a projeto tim-pmid-dev.
Caso já tenha tudo configurado e bonitinho em sua máquina pode pular esta etapa.

Para entrar no projeto você precisa logar e selecionar o projeto poderá usar:
```bash

$ gcloud init 

```

Em seguida pode setar suas credencials:
```bash
$ gcloud container clusters get-credentials pmid-dev --region southamerica-east1 --project tim-pmid-dev

```

Agora iremos listar os pods para ver se está tudo certo
```bash

$ kubectl get pods

```
Saída:
```bash
NAME                                        READY   STATUS    RESTARTS   AGE
customer-storage-manager-6f489c8d6b-8dbkq   1/1     Running   0          8m55s
kong-kong-595d7b7599-42thm                  2/2     Running   4          19h
kong-kong-595d7b7599-4tq5d                  2/2     Running   4          8h
kong-kong-595d7b7599-z5nrw                  2/2     Running   4          46h
mongodb-bd5d8499b-lv5v6                     1/1     Running   0          2d4h
test-tools-77444cf5dd-rqcth                 1/1     Running   0          46h

```

Agora vamos subir no customer-storage-manager, no diretório /deployment/k8s tem um arquivo que criamos para que possa fazer o deployment para o k8s.
```bash

$ kubectl apply -f deployments/k8s/deployment.yaml

```
Agora para checar basta fazer:

```bash

$ kubectl get pods 

```
saída:
```bash
NAME                                        READY   STATUS    RESTARTS   AGE
customer-storage-manager-6f489c8d6b-8dbkq   1/1     Running   0          8m55s
```

Agora que subimos, para não precisarmos export uma porta pública ou tentar acessar via rede na VPC vamos fazer um forward para testar nosso serviço.
```bash

$ kubectl port-forward customer-storage-manager-6f489c8d6b-8dbkq 8080

```

### Vamos testar nossos endpoints e CRUD usando Mongo

Agora vamos simular nosso CRUD se está tudo respondendo conforme o esperado.

Vamos iniciar fazendo nosso post.
```bash
$ curl -i -XPOST localhost:8080/api/v1/customers \
-H "Content-Type:application/json" \
-d '{"social_security_number":"234.567.891-34","name":"jeffotoni"}'
HTTP/1.1 201 Created
{
   "id":"e7c91765-004f-4cc4-95fb-88956c4e11ff",
   "name":"jeffotoni",
   "social_security_number":"234.567.891-34"
}
```

Agora vamos buscar pelo ID
```bash
$ curl -i -XGET \
localhost:8080/api/v1/customers/e7c91765-004f-4cc4-95fb-88956c4e11ff
HTTP/1.1 200 OK
{
   "id":"e7c91765-004f-4cc4-95fb-88956c4e11ff",
   "name":"jeffotoni",
   "social_security_number":"234.567.891-34"
}
```

Agora vamos buscar pelo social_security_number
```bash
$ curl -i -XGET \
 "localhost:8080/api/v1/customers?social_security_number=234.567.891-34"
HTTP/1.1 200 OK
{
   "id":"e7c91765-004f-4cc4-95fb-88956c4e11ff",
   "name":"jeffotoni",
   "social_security_number":"234.567.891-34"
}
```

Vamos agora atualizar os dados
```bash
$ curl -i -XPUT localhost:8080/api/v1/customers/e7c91765-004f-4cc4-95fb-88956c4e11ff \
-H "Content-Type:application/json" \
-d '{"social_security_number":"123.432.111-99","name":"Nadilson Ferreira"}'
HTTP/1.1 201 Created
{
   "name":"Nadilson Ferreira",
   "social_security_number":"123.432.111-99"
}
```

Vamos visualizar para ver se deu tudo certo nosso PUT.
```bash
$ curl -i -XGET \
localhost:8080/api/v1/customers/e7c91765-004f-4cc4-95fb-88956c4e11ff
HTTP/1.1 200 OK
{
   "id":"e7c91765-004f-4cc4-95fb-88956c4e11ff",
   "name":"Nadilson Ferreira",
   "social_security_number":"123.432.111-99"
}
```

Agora vamos remover
```bash
curl -i -XDELETE \
localhost:8080/api/v1/customers/e0c91765-004f-4cc4-95fb-88956c4e11ff

HTTP/1.1 200 OK
```

### Teste automatizado

Os testes automatizados poderão ser feitos de forma nativa, ou usando o gomock, assert ou monkey.
Iremos documentar aqui alguns passos de como está sendo feito nosso testes.

Estamos utilizando o Makefile, para otimizar alguns comandos então para executar os testes poderá usar:
```bash
$ make test
########## Executando Tests
ok    github.com/nferreira/customer-storage-manager/internal/pkg/repository/mongo (cached)
```

Quando quiser visualizar todos os detalhes do test basta fazer:
```bash
$ make test-v
--- PASS: TestMongoDbCustomerRepository (0.00s)
    --- PASS: TestMongoDbCustomerRepository/Delete_scenarios (0.00s)
        --- PASS: TestMongoDbCustomerRepository/Delete_scenarios/Connection_failed_scenarios (0.00s)
            --- PASS: TestMongoDbCustomerRepository/Delete_scenarios/Connection_failed_scenarios/Client_creation_failed (0.00s)
            --- PASS: TestMongoDbCustomerRepository/Delete_scenarios/Connection_failed_scenarios/Client_connect_failed (0.00s)
        --- PASS: TestMongoDbCustomerRepository/Delete_scenarios/Actual_deletion (0.00s)
            --- PASS: TestMongoDbCustomerRepository/Delete_scenarios/Actual_deletion/Delete_with_invalid_Mongo_ObjectID_format (0.00s)
            --- PASS: TestMongoDbCustomerRepository/Delete_scenarios/Actual_deletion/Delete_with_with_non_existent_id (0.00s)
            --- PASS: TestMongoDbCustomerRepository/Delete_scenarios/Actual_deletion/Delete_with_with_existent_id (0.00s)
PASS
ok    github.com/nferreira/customer-storage-manager/internal/pkg/repository/mongo (cached)
```

Caso tenha interesse em fazer chamadas específicas para os métodos dos testes você poderão fazer da seguinte forma.
```bash
$  go test github.com/nferreira/customer-storage-manager/internal/pkg/repository/mongo -v -run ^TestUpdateCustomerRepositoryService_Execute$
=== RUN   TestUpdateCustomerRepositoryService_Execute
=== RUN   TestUpdateCustomerRepositoryService_Execute/test_update_execute_1
=== RUN   TestUpdateCustomerRepositoryService_Execute/test_update_execute_2
--- PASS: TestUpdateCustomerRepositoryService_Execute (0.00s)
    --- PASS: TestUpdateCustomerRepositoryService_Execute/test_update_execute_1 (0.00s)
    --- PASS: TestUpdateCustomerRepositoryService_Execute/test_update_execute_2 (0.00s)
PASS
ok    github.com/nferreira/customer-storage-manager/internal/pkg/repository/mongo (cached)

```
