# go-tdd-lab

Laboratorio de exploración de TDD en Go

### Tecnologias usadas

- Go-kit
- Go modules
- gomock


### Funcionalidad

***"¿Cuanto falta para el almuerzo?"***

```
Servicio HTTP que expone un método GET,
que retorna el tiempo faltante hasta la siguiente hora de almuerzo.

Requerimientos
 -  Se debe exponer un metodo http GET que retorne un JSON con la respuesta:

    {
        // cantidad de minutos faltantes para el almuerzo                    
        "MinutosFaltantes" : "#" 
    }

 - La hora de almuerzo DEBE consultarse en un repositorio
 
```