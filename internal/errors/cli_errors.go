package errors

import "fmt"

var ErrInvalidCommandSyntax = fmt.Errorf("Некорректный синтаксис, проверьте введенную команду.")
var ErrInvalidCommandValue = fmt.Errorf("Неверный тип данных операнда(ов).")
var ErrRootDoesntExist = fmt.Errorf("В данный момент дерево пусто.")