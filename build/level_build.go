package build

import (
	"fmt"
	"strings"
)

func CheckAccess() {
	fmt.Println("Поднесите скуд:")
	var id string
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("Ошибка ID:", err.Error())
		return
	}
	floor := 0
	fmt.Println("Выберите этаж от 1 до 4:")
	_, err = fmt.Scan(&floor)
	if err != nil {
		fmt.Println("Ошибка этажа:", err.Error())
		return
	}
	switch floor {
	case 1, 2, 3, 4:
		fmt.Println("OK")
		break
	default:
		fmt.Println("Неверный этаж!")
	}
	if strings.HasPrefix(id, "1") {
		if floor > 1 {
			fmt.Println("Нет доступа!")
			return
		}
	} else if strings.HasPrefix(id, "2") {
		if floor > 2 {
			fmt.Println("Нет доступа!")
			return
		}
	} else if strings.HasPrefix(id, "3") {
		if floor > 3 {
			fmt.Println("Нет доступа!")
			return
		}
	} else if strings.HasPrefix(id, "4") {
		if floor > 4 {
			fmt.Println("Нет доступа!")
			return
		}
	} else {
		fmt.Println("Неверный ID!")
		return
	}

	fmt.Println("Добро пожаловать!")
}
