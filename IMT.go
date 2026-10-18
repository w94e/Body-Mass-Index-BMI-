package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("《 Калькулятор индекса массы тела (ИМТ) 》")
	userHeight, userWeight := getUserInput()
	IMT := calculateIMTCategory(userHeight, userWeight)
	category := getIMTCategory(IMT)
	outputIMTCategory(IMT, category)
}

func outputIMTCategory(IMT float64, category string) {
	result := fmt.Sprintf("✦› Ваш индекс массы тела (ИМТ): %.1f\n✦› Категория: %s\n", IMT, category)
	fmt.Print(result)
}

func calculateIMTCategory(userHeight float64, userWeight float64) float64 {
	IMT := userWeight / math.Pow(userHeight/100, 2)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight, userWeight float64
	fmt.Print("✦› Введите ваш рост (в сантиметрах): ")
	fmt.Scan(&userHeight)
	fmt.Print("✦› Введите ваш вес (в килограммах): ")
	fmt.Scan(&userWeight)
	if userHeight <= 0 || userWeight <= 0 {
		fmt.Println("✦› ❌ Ошибка: введены некорректные данные! Пожалуйста, введите положительные значения для роста и веса.")
		return getUserInput()
	}
	return userHeight, userWeight
}

func getIMTCategory(IMT float64) string {
	switch {
	case IMT < 18.5:
		return "Недостаточный вес"
	case IMT >= 18.5 && IMT < 24.9:
		return "Нормальный вес"
	case IMT >= 25 && IMT < 29.9:
		return "Избыточный вес"
	default:
		return "Ожирение"
	}
}
