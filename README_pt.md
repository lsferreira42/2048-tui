
# 2048 TUI

Uma implementação baseada em terminal do popular jogo 2048, escrita em Go. Esta versão apresenta uma interface colorida e efeitos sonoros para uma experiência de jogo aprimorada.

## Funcionalidades

- Jogue 2048 diretamente do seu terminal.
- Interface colorida para terminal usando o framework Bubble Tea.
- Efeitos sonoros para movimentos e game over (quando compilado com suporte de áudio).
- Compatibilidade entre plataformas (macOS, Linux, Windows).
- Layout responsivo que se ajusta ao tamanho do terminal.
- Interface de usuário simples e intuitiva para uma experiência agradável.
- Escrita em Go para alta performance e portabilidade.

## Pré-requisitos

Para compilar e executar este jogo, você precisa ter o Go instalado em seu sistema. O projeto usa módulos Go para gerenciamento de dependências.

## Instalação

Para começar com o `2048-tui`, clone o repositório e instale as dependências necessárias:

```bash
git clone https://github.com/lsferreira42/2048-tui
cd 2048-tui
```

Em seguida, execute o script fornecido para instalar as dependências:

```bash
./deps.sh
```

Este script instalará o [go-bindata](https://github.com/go-bindata/go-bindata) para gerenciar os recursos do seu projeto.

## Compilando

Para compilar o jogo com suporte a áudio, execute:

```bash
make build
```

Isso criará um binário no diretório `build`.

## Executando o Jogo

Após a compilação, você pode executar o jogo usando:

```bash
./build/2048
```

Alternativamente, para executar o jogo diretamente sem compilar:

```bash
make run
```

Ou você pode usar:

```bash
go run -tags=audio .
```

## Jogando o Jogo

- Use as setas do teclado para mover os blocos.
- Combine blocos do mesmo número para criar números maiores.
- Pressione 'r' para reiniciar o jogo.
- Pressione 'q' para sair.
- Tente alcançar o bloco 2048!

## Distribuição

Para criar binários para várias plataformas:

```bash
make dist
```

Isso criará binários para macOS (x86_64 e ARM), Linux (x86, ARM e x64) e Windows (64-bit) no diretório `dist`.

## Estrutura do Projeto

- `2048.go`: Lógica principal do jogo e interface.
- `audio.go`: Funcionalidade de reprodução de áudio.
- `deps.sh`: Script para instalar dependências adicionais.
- `Makefile`: Comandos para compilar e executar.
- `move_sound.mp3`: Efeito sonoro para movimentos.
- `ending.mp3`: Efeito sonoro para game over.

## Dependências

- **Go**: A versão mais recente é recomendada. Go é necessário para compilar e executar o projeto. Você pode instalá-lo seguindo as instruções no [site oficial](https://golang.org/doc/install).
- **go-bindata**: Esta ferramenta é usada para embutir dados em aplicativos Go, facilitando o gerenciamento de recursos estáticos. Você pode instalá-la executando o seguinte comando:

  ```bash
  go install -a -v github.com/go-bindata/go-bindata/...@latest
  ```

- **GNU Make**: Usado para compilar o projeto com o `Makefile`. Certifique-se de que o GNU Make está instalado no seu sistema. Na maioria dos sistemas Linux, você pode instalá-lo com:

  ```bash
  sudo apt-get install make
  ```

  No macOS, você pode usar:

  ```bash
  brew install make
  ```

- **Bubble Tea**: Framework de interface de usuário para terminal usado para construir a interface colorida.
- **Lip Gloss**: Definições de estilo para a interface do terminal.

## Licença

Este projeto é de código aberto e está disponível sob a [Licença MIT](LICENSE).

## Contribuindo

Sinta-se à vontade para enviar problemas, sugestões ou pull requests para ajudar a melhorar o projeto. Contribuições são sempre bem-vindas!

## Contato

Mantenedor: [Leandro Ferreira](https://github.com/lsferreira42)

URL do Repositório: [2048 TUI no GitHub](https://github.com/lsferreira42/2048-tui)

## Agradecimentos

- Inspirado pelo jogo original 2048 de Gabriele Cirulli.
- Agradecimentos à equipe Charm pelas excelentes bibliotecas de interface de terminal.
