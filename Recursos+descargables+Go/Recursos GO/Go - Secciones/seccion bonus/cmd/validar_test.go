package cmd

import (
    "testing"
)

func TestValidarUsuario(t *testing.T) {
    // Caso válido
    err := ValidarUsuario("Ana", 25)
    if err != nil {
        t.Errorf("Esperaba nil, obtuvo error: %v", err)
    }

    // Caso nombre vacío
    err = ValidarUsuario("", 25)
    if err == nil {
        t.Error("Esperaba error por nombre vacío, obtuvo nil")
    }

    // Caso edad <= 0
    err = ValidarUsuario("Pedro", 0)
    if err == nil {
        t.Error("Esperaba error por edad <= 0, obtuvo nil")
    }
}
