# Estrutura de pastas do GO

após instalação do GO, seram geradas 3 diretórios que são utilizados pela linguagem para armazenar binarios, arquivosde compilação e código fonte de arquivos .go

- go/bin -> ficam todos os binários utilizados pelo go
- go/pkg -> ficam arquivos utilizados para pré compilação do go
- go/src -> ficam os os arquivos .go para o go conseguir interpretar o fonte

# Anotações 01_fundamentos

- `go env` lista todas as variaveis de ambiante que a golang usa no SO
    `GOPATH` indica onde os binarios, libs, fonte do GO ficam

- `go mod` é para gerenciamento de dependencias do GO, isso permite o desenvolvimento de apps GO fora do diretório go/src
- modulos permitem desenvolver apps Go fora do diretorio de instalacao do go
- `go mod init nome-modulo` inicia um novo modulo para o Go gerenciar os pacotes/dependencias, isso faz o Go nao olhar mais para a paste go/src e sim para a pasta corrente do projeto
    - rodando esse comando, é gerado um arquivo go.mod, que vai ter o nome do modulo, a versao do go e todas as dependencias da aplicacao
    - `go mod tidy` é usado para otimizar o go.mod, baixando as deps que a app utiliza mas nao estao no go.mod e também faz a remoção de pacotes que nao sao mais utilizados

- arquivos .go por padrão devem ter uma declaração de `package` logo no inicio, esse package deve ter o nome do diretório que o arquivo esta, com exeção do package main, pois esse é o package principal que contem a funcao main que é o start da app

- todos os aquivos .go dentro de um diretório devem ter o mesmo nome de package, ou seja, o nome do diretório

- tudo que está dentro do mesmo diretório é visivel em todos os arquivos .go, variaveis, funcoes publicas, structs ...

- GO é uma linguagem fortemente tipada

- GO permite a criação de novos tipos de dados customizados

- _ (underline) é um blank identifier, ele diz para o go ignorar alguma variavel, retorno de funcao...

- e GO é possivel que uma funçã retorne mais de um valor

- em GO não existe herança, apenas composição entre structs e polimorfismo com interfaces

- em GO metodos de interfaces são implementados implicitamente pela linguaguem, ou seja, se uma struct tiver um metodo com a mesma assinatura de um método de uma interface, automaticamente essa struct estara implementando a interface.

- variavel aponta para endereço de memoria (ponteiro) que tem um valor; pegando diretamente esse endereço da memoria (ponteiro) e mudar seu valor (o dado que o ponteiro aponta), quando a variavel acessar esse valor ele vai estar diferente; variavel -> ponteiro -> valor ou ponteiro -> valor

- nao devemos usar ponteiros quando queremos apenas passar uma cópia dos dados
- devemos usar ponteiros quando queremos tornar os dados mutáveis em qualquer ponto do código, alterando valor direto a memoria

- structs, variaveis, metodos e funcoes iniciadas com letra maiuscula são publicas e visiveis por quem importar o pacote
- structs, variaveis, metodos e funcoes iniciadas com letra minuscula são privadas e visiveis apenas internamente no pacote

- `go get nome-pacote` diz para o Go instalar um novo pacote
- go.sum é gerado depois do go get, e ele garante a versao dos pacotes externos instalados, para eles nao sere atualizados sempre

- em Go não existe if ternário e else-if

- Go permite escolher qual SO queremos compilar nosso app

- `go build main.go` compila pra gerar o binário da app

- `GOOS=sistema-operacional (linux, windows ...) go build main.go` compila e gera o binário da app para um SO especifico

- `GOOS=sistema-operacional (linux, windows, mac ...) GOARCH=arquitetura-so-processador go build main.go` compila e gera o binário da app para um SO especifico e um tipo de arquitetura de processador especifico

- `go build` em uma app com modulo, na hora do build o go vai escanear toda a app para achar a funcao/package main para gerar o binario e por padrao o binario vai ser gerado com o nome do modulo

- `go build -o nome-binario` gera o binario com um nome especifico atraves do parametro -o

# Anotações 02_pacotes_importantes

- `defer` é um statement em go que faz uma linha de instrucao dentro de uma funcao ser executada por ultimo, o defer vai segurar a execucao da instrucao até que ela seja a ultima coisa a ser executada

- `json.Marshal(value)` geralmente é utilizado quando precisamos armazenar o retorno em json em uma variavel para usar posteriormente

- `json.Unmarshal(value)` é o inverso do Marshal
    - para ter conversao correta do json com struct, o json deve ter os mesmos campos que a struct tem para a funcao conseguir fazer a conversao correta exceto se tiver sendo utilizado tags na struct

- `json.NewEncoder(out).Encode(value)` geralmente é utilizado quando queremos converter o valor para jjson e já enviar para algum lugar (console, arquivo...), sem a necessidade de armazenar em uma variavel

- `json.NewDecoder(out).Decode(value)` fuciona inversamente ao Encoder, convertendo um json para algum valor e entregando para alguem (console, arquivo...)

- Tags `json:"numero"` são como anotacoes em Go, o que diz para as bibliotecas como fazerem o bid das informacoes para as structs
    - tags também são usadas para validacao de dados em uma struct
    - para ignorar campos usando as tags no parse para json, usar `json:"-"`

- `Multiplexer` -> atachador de rotas global no Go quando um server é iniciado
- ServeMux é utilizado para ter mais controle no servidor, no registro de rotas

- Sempre que trabalhar com templates em html, usar o pacote `html/template` pois ele vai implementações seguras para evitar alguns tipos de ataque em htmls

- Sempre que trabalhar com templates em texto, usar o pacote `text/template`

# Anotações 03_contexts

- pacote de context, servem para controlar o tempo de uma operacao podendo cancelar a operacao caso exceda o tempo

- contextos são utilizados para cancelar operacoes

- é possivel armazenar informacoes dentro de um contexto, mas é controverso

- contextos podem ser usandos em chamadas http, consultas a banco de dados ...

- por convenção da linguagem o contexto sempre dever ser o primeiro parametro das funcoes

- não é recomendado passar dados por contexto

# Anotações 04_banco_de_dados

- para abrir conexao com banco: sql.Open(diverName, connString) -> necessario informar o driver (mysql, sqlite...) e a string de conexao

- Go só expoe as interfaces para interação com DB, as implementacoes não fazem parte do Go, por isso é necessario instalar os pacotes com as implementacoes especificas

- GORM é utilizado para fazer o mapeamento objeto relacional, mapeando as structs para tabelas da base

- soft delete significa que os registros de uma base nunca serao excluidos realmente, eles terao um campo indicando a data e hora em que o registro foi "deletado", porem o dado continua na base

- BelongsTo: um dado que pertence a outro dado, por exemplo, um produto que pertence a uma categoria ou varios produtos que pertencem a uma categoria

- HasOne: é uma relacionamento de 1:1

- HasMany: é um relacionamento de 1:n

- ManyToMany: é um relacionamento de n:n (muitos para muitos)

- Lock otimista: versiona qualquer tipo de alteração em um dado na base de dados
    - versão igual, commit é feito;
    - versão diferente, aconteceu alguma atualização no dado o que fará o processo ser reiniciado, pois o dado foi alterado.
    - usado em um ambiente com muitas transacoes, mas sem concorrencia

- Lock pessimista: locka a tabela, linha na base de dados durante as atualizações para que ninguem atualize o dado.
    - usado em um ambiente com muita concorrencia. Para garantir que nenhum processo sobreescreva a atualizacao do outro.
    - na query FOR UPDATE indica que a linha será lockada até que a operacao termine

# Anotações 05_packaging

- Funcoes, Struct que iniciarem com letras Maiusculas são exportadas

- Funcoes, Struct que iniciarem com letras Minusculas são visiveis apenas dentro do proprio diretório

- Go trabalha com módulos;

- por boa prática é bom iniciar projetos em Go como modulos, pois o Go sempre espera que o projetos sejam criados dentro do GOTPATH, criando a aplicação como modulo podemos estar desenvolvendo fora do GOPATH

- go mod trabalha de forma descentralizada, buscando os modulos onde eles estiverem hospedados

- ao iniciar o modulo, por convesao devemos dar o nome da url do repositório onde o projeto esta, deve ser unico para evitar conflito
`go mod init url-local-do-modulo`

- `go mod tidy` -> vai avaliar o código e as importações de pacotes para baixar os pacotes que ainda não estão na dependencia, caso um pacote deixa de ser utilizado esse comando vai remover essa dependencia desnecessaria

- `go mod tidy -e` vai fazer o go ignorar os pacotes que ele não achar

- go.mod vai ser um gerenciador de dependencias das aplicacoes em Go

- go.sum é um arquivo de lock, para garantir a versão das dependencias, para o go mod tidy nao ficar atualizando as versoes

- go workspaces -> são workspaces locais para isolar as dependencias
    `go work init nome-libs/modulos` comando para criar workspaces locais no go

- cmd/ fica os pacotes com a funcao main

# Anotações 06_testing

- Go já possui um suite de testes embutida na linguagem

- por convenção os arquivos de teste devem ter o sufixo `_test`

- `go test .` faz o Go rodar todos os testes dentro do diretorio

- `go test -coverprofile=coverage.out` faz o go rodar os testes verificando a cobertura dos testes no código

- `go tool cover -html=coverage.out` exporta para html o ponto do código onde nao esta tendo cobertura de código 

- go possui uma ferramenta de benchmarking para saber a performance das funcoes

- `go test -bench .` faz o go rodar um teste de benchmark de uma função

- `go test -bench . -run=^#` faz o go rodar apenas os testes de benchmark
BenchmarkCalculateTax-n-cores-cpu | 1000000000(n operacoes) | 0.2540 ns/op (operacoes por nanosegundos)

- `go test -bench . -run=^# -benchmem` rodando branchmark validando a memoria

- fuzzing é um tipo de teste em Go, que vai testar uma função com variações dos parametros que a função recebe para ver se em algum momento um dos parametros passados quebra a função

- `go test -fuzz . -run=^#` usado para rodar testes de mutação no Go

- `go test -run=caminho-da-fuzz-gerado` comando para executar um teste fuzz especifico

- `go test -fuzz . -fuzztime 5s -run=^#` rodando os testes fuzz com um tempo especifico

- `testify` é um pacote bem famoso que facilita os asserts de testes em go

# Anotações 07_APIs

Principais diretórios convencionados estruturar apps em Go

- `cmd/` -> aqui ficam os arquivos Go que geram o executavel da aplicação, ou seja, possuem a função main. Geralmente, dentro desse diretório existe um outro diretório com o nome da aplicação em questão que fica o arquivo com a função main.

- `internal/` -> aqui ficam os arquivos Go, referentes a aplicação que esta sendo desenvolvida. São arquivos que só serão usados internamente na aplicação que esta sendo desenvolvida.

- `pkg/` -> aqui ficam os arquivos Go, que são considerados libs e que são genericos o suficiente para serem reutilizados por outras aplicações ou pela mesma app.

- `configs/` -> aqui ficam os arquivos de configuração de inicialização da app (arquivos Go ou outro tipo de template de configuração)

- `test/` -> aqui ficam arquivos adicionais que são utilizados pelos arquivos de teste (podem ser arquivos Go ou outros tipos de arquivos)

- `api/` -> aqui ficam os arquivos de especificações, documentações da api. Como por exemplo, swagger.

- roteadores são responsáveis por registrar e agrupar rotas, middlewares ...

- middlaware funciona como um intermediário recebendo uma req (ou outra coisa), realizando algum processamento em cima dessa req e chama um handler (ou outro middleware) para a req continuar

- https://github.com/swaggo/swag para gerar documentação de api

- `swag init -g caminho-arquivo-go-base` comando para gerar a doc swagger inicial