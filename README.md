# DSP_Go

Библиотека инструментов для цифровой обработки сигналов. 
Попытка написать что-то собственное.

## Тестирование
### Запуск всех тестов
```bash
go test ./pkg/...
```
### Запуск всех с подробным выводом
```bash
go test -v ./pkg/...
```

### Запуск только КИХ-тестов
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

### Запуск только генераторов
```bash
go test ./pkg/generators/...
```

### Запуск только генераторов
```bash
go test ./pkg/generators/... -run "TestInfo"
```


### Запуск с покрытием кода
```bash
go test -cover ./pkg/...
```

### Запуск бенчмарков
```bash
go test -bench=./pkg/...
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