# Desafio: Explorando Marte

## Descrição do Problema

Um conjunto de sondas (rovers) foi enviado pela NASA à Marte e irá pousar num planalto retangular. Estas sondas devem percorrer a área mapeada, seguindo instruções de controle para se movimentar e capturar imagens.

### Regras Básicas

- As posições no planalto são representadas por coordenadas x-y
- A direção da sonda é indicada por uma letra que representa um ponto cardeal (N, S, E, W)
- O planalto é definido pelas coordenadas do canto inferior esquerdo (0,0) e do canto superior direito (coordenadas máximas dadas na entrada)
- As sondas recebem sequências de instruções:
  - **L**: girar 90 graus à esquerda
  - **R**: girar 90 graus à direita
  - **M**: mover um ponto à frente na direção atual
- Uma sonda deve executar todas as suas instruções antes que a próxima sonda seja processada

## Especificações Técnicas

### Validação de Limites

Se a sonda receber o comando "M" e isso a levaria para fora do planalto, você pode escolher como tratar:
- Impedir o movimento e ignorar o comando
- Lançar uma exceção/erro
- Registrar em log ou mensagem de erro

É obrigatório que sua solução trate esse cenário de alguma forma clara.

### Múltiplas Sondas & Possível Conflito de Posição

Vamos supor que não queremos que duas sondas ocupem a mesma posição simultaneamente. Caso a próxima sonda tente se mover para uma posição já ocupada, também decida como tratar: ignorar movimento, levantar exceção etc.

Não é necessário implementar paralelismo/concorrência de verdade (uma sonda se move por vez, sequencialmente), mas deve haver alguma checagem de posição para evitar "colisões".

### Formas de Entrada e Saída

Fica livre para você escolher, mas justifique no README ou na documentação do projeto por que você escolheu tal forma (CLI, Web API, arquivo de texto, etc.).

**Exemplo clássico de entrada:**
```
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
```

**Exemplo clássico de saída:**
```
1 3 N
5 1 E
```

## Requisitos de Implementação

### Código Orientado a Objetos

- **Abstração e Polimorfismo**: separe as responsabilidades em classes/módulos
- **Herança**: caso faça sentido (por exemplo, se houver uma classe abstrata para tipos de veículos espaciais)
- **SOLID**: é muito importante que possamos enxergar princípios como SRP (Single Responsibility), OCP (Open/Closed) etc.

### Design Patterns (GoF)

Você pode usar:
- **Command Pattern** para cada instrução (L, R, M)
- **Factory Pattern** para criar e gerenciar sondas/direções
- **Strategy Pattern** para orientar como cada sonda se move

Fique à vontade para escolher o que achar mais adequado, mas documente no README qual(is) pattern(s) você utilizou e por quê.

### Testes Unitários

Crie testes automatizados que validem pontos como:
- Movimentação básica e mudança de direção
- Respeito aos limites do planalto
- Conflito de posição (se implementado)
- Vários cenários de entrada (inclusive cenários de erro/exceções, se existirem)

Use o framework de teste da sua linguagem (JUnit, NUnit, PyTest, etc.). Teste suas abstrações (classes, métodos), não apenas um "teste geral" do main.

### Debugging e Troubleshooting

Ao menos em um trecho do seu README (ou em comentários no código), explique como você utilizou as ferramentas de debugging do VSCode (breakpoints, inspeção de variáveis etc.) para investigar um possível problema no código.

Não queremos um tutorial de VSCode, apenas uma descrição de como você configurou o launch.json (se aplicável) ou como define breakpoints e faz step-by-step. Queremos ver se você domina a ferramenta.

### Implementação de Pipeline de Continuous Integration

Crie um arquivo de configuração (por exemplo, `.github/workflows/ci.yml` ou `Jenkinsfile`) para garantir que seus testes sejam executados automaticamente.

Descreva no README como o pipeline foi configurado e como podemos verificar os resultados (logs, relatórios etc.).

## Critérios de Avaliação

- Qualidade do código (organização, clareza, aderência a boas práticas)
- Uso coerente de conceitos de OOP (herança, abstração, polimorfismo)
- Aplicação dos princípios SOLID
- Implementação de um ou mais Design Patterns e a justificativa de uso
- Testes unitários e clareza na forma como foram escritos (cobertura de cenários principais e possíveis falhas)
- Histórico de commits (commits pequenos e claros são preferíveis a um "pacotão" único)
- Documentação (um README explicando como rodar, como estão estruturadas as classes, como usar o debugger etc.)
- Pipeline de Continuous Integration (arquivo de configuração, execução dos testes em ambiente automatizado, explicação no README)

## Entregáveis

### GitHub
- Crie um repositório público
- Commits: Tente manter um fluxo de commits que mostre sua evolução

### README
O README deve conter:
- Descrição do problema de forma resumida
- Como configurar e rodar seu projeto
- Decisões de projeto
- Qual padrão de projeto (Design Pattern) foi adotado e por quê
- Como você fez o debugging no VSCode
- Como o pipeline de CI foi configurado e como acompanhar os resultados
- Se existir algo adicional que queira destacar, inclua no README ou como comentários no código

### Linguagem de Programação
Livre escolha.

## Entrega Final

1. Faça um vídeo com a gravação da sua tela enquanto está PROGRAMANDO o desafio, e suba no YouTube
2. Não tem problema se pausar a gravação durante o desafio ou enviar vários vídeos
3. Queremos entender a sua lógica de pensamento
4. Não queremos que mostre somente o resultado final, queremos ver como programou, seja usando AI ou não

## Observações Importantes

- **IMPORTANTE**: Na falta de algum dos itens acima, está desclassificado
- **IMPORTANTE 2**: Não vemos problema usar um Agente de Inteligência Artificial para resolver o desafio. Entretanto, caso utilize, o importante é você nos avisar qual Agente usou, e na etapa de Entrevista técnica, conseguir explicar e tirar um 10 nas explicações do código que fez em conjunto com o seu Agente

## Prazo
48h da aceitação do desafio

---

**Boa sorte! E esperamos que faça o seu melhor.**
