// Archivos mal formateados (campos faltantes).
// Valores inesperados (texto donde debería ir un número).
// Codificación inconsistente (caracteres especiales).
// Archivos incompletos o líneas vacías.

// Validando longitud de registros:
for {
    record, err := reader.Read()
    if err != nil {
        break
    }

    if len(record) < 3 {
        log.Println("Registro incompleto:", record)
        continue
    }

    // Procesar normalmente...
}

// validar conversion de tipos

edad, err := strconv.Atoi(record[1])
if err != nil {
    log.Println("Edad inválida:", record[1])
    continue
}

// ignorar lineas vacias

if len(record) == 0 {
    continue
}


// validar valores con reglas de negocio


if edad <= 0 {
    log.Println("Edad no válida:", edad)
    continue
}

log.Println
log.Printf


