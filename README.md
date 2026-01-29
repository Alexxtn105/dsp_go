## Тестирование
### Запуск всех тестов
```bash
go test ./pkg/filters/...
```

### Запуск только БИХ-тестов
```bash
go test ./pkg/filters/... -run "FIR"
```

### Запуск только БИХ-тестов
```bash
go test ./pkg/filters/... -run "IIR"
```

### Запуск только тестов фильтра Герцеля
```bash
go test ./pkg/filters/... -run "Goertzel"
```

### Запуск с подробным выводом
```bash
go test -v ./pkg/filters/...
```

### Запуск с покрытием кода
```bash
go test -cover ./pkg/filters/...
```

### Запуск бенчмарков
```bash
go test -bench=. ./pkg/filters/...
```

### Генерация godoc
```bash
godoc -http=:6060
```

### Просмотр тестового покрытия
```bash
go test -coverprofile=coverage.out ./pkg/filters/
```

```bash
go tool cover -html=coverage.out
```