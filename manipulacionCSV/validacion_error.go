//* validando longitud de registros

for {
	reacord, err := reader.Read()
	if err != nil {
		break
	}

	if len(record) < 3 {
		log.Println("Registro incompleto:", record)
		continue
	}

	// procesar normalment ...

}

//** validar conversion de tipos
edad, err := strconv.Atoi(record[1])
if err != nil {
	log.Println("Edad inválida:", record[1])
	continue
}

// ignorar linead vacias

if len(record) == 0 {
	continue
}

// *validar valores con reglas de negocio

if edad <= 0 {
	log.Println("Edad no válida:", edad)
	continue
}

log.Println
log.Printf