# Biblioteca de Caché LRU en Go

Esta es una biblioteca simple de caché LRU (Least Recently Used - Menos Recientemente Usado) escrita en Go. Proporciona
una forma conveniente de almacenar pares clave-valor con características opcionales como establecer una capacidad máxima
y un tiempo de expiración para cada elemento.

## Uso

### 1. Importación

```bash
go get github.com/kevinlugodev/golrucache
```

### 2. Crear caché
```go
cache := lrucache.NewLRUCache(&uint(2))
```

### 3. Establecer un Par Clave-Valor
```go
cache.Set("clave", "valor")
```

### 4. Obtener un Valor
```go
valor, encontrado := cache.Get("clave")
if encontrado {
    fmt.Println("Valor:", valor)
} else {
    fmt.Println("Clave no encontrada")
}
```

### 5. Limíar caché
```go
cache.Clear()
```

## Contacto

Si tienes alguna pregunta o sugerencia, no dudes en contactarme:

- **Instagram:** [@kevinlugo.dev](https://www.instagram.com/kevinlugo.dev?theme=dark)
- **LinkedIn:** [Kevin Harrinson Lugo Díaz](https://www.linkedin.com/in/kevinlugodiaz)