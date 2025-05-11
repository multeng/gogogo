package module4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInputAsFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return 0, fmt.Errorf("ошибка чтения: %v", err)
	}
	input = strings.TrimSpace(input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка %v, введите число", err)
	}

	return math.Round(value*10) / 10, nil
}

func getBodyState(imt float64) string {
	switch {
	case imt < 18.5:
		return "Недостаточный вес"
	case imt < 25:
		return "Нормальный вес"
	case imt < 30:
		return "Избыточный вес"
	default:
		return "Ожирение"
	}
}

func Task2() {
	fmt.Println("Введите ваш вес в кг")
	weight, err := readInputAsFloat()

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Введите ваш рост в см")

	height, err := readInputAsFloat()

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	imt := weight / math.Pow(height/100, 2)
	bodyState := getBodyState(imt)
	fmt.Printf("Ваш ИМТ: %.2f\n", imt)
	fmt.Printf("Категория: %s\n", bodyState)
}
