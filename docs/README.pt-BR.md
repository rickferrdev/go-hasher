# Bcrypt Hasher Wrapper (Go-Hasher)

Este projeto é um wrapper simples para a biblioteca `golang.org/x/crypto/bcrypt`. Ele foi desenvolvido exclusivamente para fins de estudos pessoais, focando na abstração de lógica de hashing e implementação de interfaces em Go.

## Objetivo

O pacote fornece uma interface padronizada para geração e comparação de hashes, facilitando a substituição de implementações ou o mock de comportamentos em testes unitários.

## Funcionalidades

* Gerar hashes bcrypt a partir de slices de bytes.
* Validar senhas contra hashes existentes.
* Tratamento de erros customizado.
* Validação automática do custo (cost) do algoritmo.

## Interface

O pacote expõe a seguinte interface:

```go
type Hasher interface {
    Generate(password []byte) ([]byte, error)
    Compare(hash, password []byte) error
}

```

## Exemplo de Uso

### Inicialização

Para criar uma nova instância com o custo padrão:

```go
h := hasher.New(hasher.DefaultCost)

```

### Gerando um Hash

```go
password := []byte("minha_senha_segura")
hash, err := h.Generate(password)
if err != nil {
    // tratar erro
}

```

### Comparando uma Senha

```go
err := h.Compare(hash, []byte("minha_senha_segura"))
if err != nil {
    // se o erro for "hasher: invalid password", a senha está incorreta
}

```

## Constantes de Custo

O wrapper expõe as constantes de custo do bcrypt para facilitar a configuração:

* `MinCost`: 4
* `MaxCost`: 31
* `DefaultCost`: 10

## Limitações

Conforme a implementação do bcrypt em Go, o comprimento máximo da senha é de 72 bytes. Caso uma senha maior seja enviada ao método `Generate`, o erro `bcrypt.ErrPasswordTooLong` será retornado.
