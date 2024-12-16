package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aeonva1ues/avltree/internal/algorithms"
	"github.com/aeonva1ues/avltree/internal/entity"
	"github.com/aeonva1ues/avltree/internal/errors"
)

func main() {
	var executed bool
	avl := entity.NewEmptyTreeNode()
	fmt.Println(
		"Здравствуйте! Данная программа реализует структуру данных AVL дерево и предоставляет операции для работы с ней.\n",
		"Доступные операции:\n",
		"PRINT [1 - InOrder | 2 - PostOrder | 3 - PreOrder] - вывести дерево\n",
		"ROOT - вывести корень дерева\n",
		"ADD [ЦЕЛ. ЧИСЛО] - добавить значение в дерево\n",
		"DELETE [ЦЕЛ. ЧИСЛО] - удалить значение из дерева\n",
		"EXIT - выйти из программы",
	)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		executed = false
		command := sc.Text()
		if command == "EXIT" {
			break
		}
		operands := strings.Split(command, " ")
		if len(operands) == 2 {
			if operands[0] == "PRINT" {
				if operands[1] == "1" {
					algorithms.InOrder(avl)
					executed = true
				} else if operands[1] == "2" {
					algorithms.PostOrder(avl)
					executed = true
				} else if operands[1] == "3" {
					algorithms.PreOrder(avl)
					executed = true
				}
			} else if operands[0] == "ADD" {
				newNumber, err := strconv.Atoi(operands[1])
				if err != nil {
					fmt.Println(errors.ErrInvalidCommandValue)
				} else {
					avl = algorithms.AvlInsert(avl, newNumber, 0)
					executed = true
				}
			} else if operands[0] == "DELETE" {
				newNumber, err := strconv.Atoi(operands[1])
				if err != nil {
					fmt.Println(errors.ErrInvalidCommandValue)
				} else {
					avl = algorithms.Delete(avl, newNumber)
					executed = true
				}
			}
		} else if len(operands) == 1 {
			if command == "ROOT" {
				fmt.Println(avl.Value)
				executed = true
			}
		}
		if !executed {
			fmt.Println(errors.ErrInvalidCommandSyntax)
		} else {
			fmt.Println("[+] Выполнено успешно")
		}
	}
}