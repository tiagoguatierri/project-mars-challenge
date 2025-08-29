# Mars Rover Challenge

## Descrição do Problema

Este projeto implementa uma solução para o desafio "Explorando Marte", onde um conjunto de sondas (rovers) é enviado pela NASA à Marte e deve pousar num planalto retangular. As sondas percorrem a área mapeada seguindo instruções de controle para se movimentar e capturar imagens.

### Regras Básicas

- As posições no planalto são representadas por coordenadas x-y
- A direção da sonda é indicada por uma letra que representa um ponto cardeal (N, S, E, W)
- O planalto é definido pelas coordenadas do canto inferior esquerdo (0,0) e do canto superior direito
- As sondas recebem sequências de instruções:
  - **L**: girar 90 graus à esquerda
  - **R**: girar 90 graus à direita
  - **M**: mover um ponto à frente na direção atual
- Uma sonda executa todas as suas instruções antes que a próxima sonda seja processada

## Como Executar

### Pré-requisitos

- Go 1.24.5 ou superior

### Instalação e Execução

1. Clone o repositório:
```bash
git clone <repository-url>
cd github.com/tiagoguatierri/project-mars-challenge
```

2. Execute o programa:
```bash
go run cmd/app/main.go
```

3. Verifique o resultado no arquivo `output.txt`

### Formato de Entrada

O programa lê as instruções do arquivo `input.txt` no seguinte formato:

```
5 5
1 2 N Goku
LMLMLMLMM
3 3 E Niraw
MMRMMRMRRM
```

Onde:
- Primeira linha: dimensões do planalto (maxX maxY)
- Linhas subsequentes em pares:
  - Posição inicial e direção da sonda (x y direção nome)
  - Sequência de comandos (L, R, M)

### Formato de Saída

O resultado é salvo no arquivo `output.txt`:

```
1 3 N Goku
5 1 E Niraw
```

## Arquitetura do Projeto

### Estrutura de Diretórios

O projeto segue o **Layout Padrão do Go** (Go Standard Project Layout), que é uma convenção amplamente aceita pela comunidade Go:

```
github.com/tiagoguatierri/project-mars-challenge/
├── cmd/                  # Aplicações executáveis
│   └── app/              # Aplicação principal
│       └── main.go       # Ponto de entrada da aplicação
├── internal/             # Código interno do projeto (privado)
│   ├── domain/           # Entidades de domínio
│   │   ├── plateau.go    # Representação do planalto
│   │   ├── rover.go      # Representação da sonda
│   │   ├── direction.go  # Enumeração de direções
│   │   └── coord.go      # Estrutura de coordenadas
│   └── command/          # Padrão Command
│       ├── command.go    # Interface e dispatcher
│       ├── move.go       # Comando de movimento
│       ├── turn_left.go  # Comando de rotação à esquerda
│       └── turn_right.go # Comando de rotação à direita
├── input.txt             # Arquivo de entrada
├── output.txt            # Arquivo de saída
├── go.mod                # Dependências do Go
└── README.md             # Documentação do projeto
```

### Design Patterns Implementados

#### 1. Command Pattern

O projeto utiliza o **Command Pattern** para encapsular as instruções de movimento das sondas. Cada comando (L, R, M) é representado por uma classe separada que implementa a interface `Command`.

**Benefícios:**
- Separação de responsabilidades
- Facilita a adição de novos comandos
- Permite desacoplamento entre o dispatcher e os comandos específicos

**Implementação:**
```go
type Command interface {
    Execute(rover *domain.Rover) error
}
```

#### 2. Factory Pattern

Utilizamos o **Factory Pattern** para criar instâncias de `Plateau` e `Rover`, encapsulando a lógica de parsing das instruções.

**Benefícios:**
- Centraliza a lógica de criação de objetos
- Facilita a validação de dados de entrada
- Melhora a testabilidade do código

#### 3. Domain-Driven Design (DDD)

O projeto segue princípios de DDD com:
- **Entidades de Domínio**: `Plateau`, `Rover`, `Direction`, `Coord`
- **Validação de Regras de Negócio**: Verificação de limites e colisões
- **Separação de Responsabilidades**: Cada entidade tem suas responsabilidades bem definidas

### Princípios SOLID Aplicados

#### 1. Single Responsibility Principle (SRP)
- Cada classe tem uma única responsabilidade
- `Plateau`: gerencia o estado do planalto e validações de posição
- `Rover`: representa o estado de uma sonda
- `Command`: encapsula uma ação específica

#### 2. Open/Closed Principle (OCP)
- O sistema é aberto para extensão (novos comandos) e fechado para modificação
- Novos comandos podem ser adicionados sem modificar o dispatcher existente

#### 3. Liskov Substitution Principle (LSP)
- Todos os comandos implementam a mesma interface `Command`
- Podem ser substituídos sem afetar o comportamento do sistema

#### 4. Interface Segregation Principle (ISP)
- Interfaces pequenas e específicas
- `Command` interface contém apenas o método necessário

#### 5. Dependency Inversion Principle (DIP)
- O dispatcher depende de abstrações (interface `Command`)
- Não depende de implementações concretas

### Validações Implementadas

#### 1. Validação de Limites
- Verifica se o movimento manteria a sonda dentro dos limites do planalto
- Impede movimentos que levariam a sonda para fora do planalto

#### 2. Conflito de Posição
- Verifica se a posição de destino já está ocupada por outra sonda
- Impede colisões entre sondas

#### 3. Validação de Entrada
- Verifica se as coordenadas são números inteiros válidos
- Valida se as direções são válidas (N, S, E, W)
- Verifica se todos os argumentos necessários estão presentes

## Debugging no VSCode

1. **Definir Breakpoints**: Clique na margem esquerda do editor para definir breakpoints
2. **Iniciar Debugging**: Pressione F5 ou use o menu "Run and Debug"
3. **Step-by-Step**: Use F10 (Step Over), F11 (Step Into), F12 (Step Out)
4. **Inspeção de Variáveis**: Use o painel "Variables" para examinar valores
5. **Call Stack**: Acompanhe a pilha de chamadas no painel "Call Stack"

### Exemplo de Debugging

Para debugar um problema de movimento:
1. Defina breakpoints em `internal/command/move.go:Execute()`
2. Defina breakpoints em `internal/domain/plateau.go:Move()`
3. Execute o programa em modo debug
4. Inspecione os valores de `delta`, `nX`, `nY` durante a execução
5. Verifique se as validações de limite estão funcionando corretamente

## Pipeline de Continuous Integration

### GitHub Actions

Crie um arquivo `.github/workflows/ci.yml`:

```yaml
name: CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.24.5'
    
    - name: Test
      run: go test -v ./...
    
    - name: Build
      run: go build -v ./...
    
    - name: Run Mars Rover
      run: |
        go run cmd/app/main.go
        cat output.txt
```

### Como Acompanhar os Resultados

1. Acesse a aba "Actions" no GitHub
2. Clique no workflow "CI"
3. Verifique os logs de execução
4. Os testes são executados automaticamente a cada push/PR

## Testes

Para executar os testes:

```bash
go test -v ./...
```

Para executar testes com cobertura:

```bash
go test -v -cover ./...
```

## Decisões de Projeto

### 1. Linguagem de Programação: Go

**Justificativa:**
- Simplicidade e legibilidade
- Forte tipagem estática
- Excelente suporte para concorrência (futuras extensões)
- Ferramentas nativas para testes e formatação

### 2. Arquitetura Modular (Layout Go Standard)

**Justificativa:**
- Segue o layout padrão do Go com `cmd/` e `internal/`
- Facilita manutenção e extensão
- Melhora testabilidade
- Segue princípios de Clean Architecture
- `cmd/` contém aplicações executáveis
- `internal/` contém código privado do projeto

### 3. Tratamento de Erros

**Justificativa:**
- Uso de tipos de erro customizados
- Propagação adequada de erros
- Logs informativos para debugging

### 4. Formato de Entrada/Saída

**Justificativa:**
- Arquivo de texto simples para entrada
- Formato legível e fácil de testar
- Facilita integração com outros sistemas

## Extensões Futuras

1. **Interface Web**: API REST para controle das sondas
2. **Concorrência**: Múltiplas sondas se movendo simultaneamente
3. **Persistência**: Banco de dados para histórico de movimentos
4. **Visualização**: Interface gráfica para acompanhar movimentos
5. **Novos Comandos**: Comandos avançados como "S" (stop) ou "P" (photo)

## Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.
