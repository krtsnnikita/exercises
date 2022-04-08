/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/

// Программа для расчёта стоимости поездки
package main

import (
        "fmt"
        "os"
)

// Главная функция 
func main() {

        const gasolineCost = 48 // Стоимость бензина за 1 литр
        var distance int        // Расстояние поездки
        var consumption int     // Расход топлива

        fmt.Print("Введите расстояние от 50 до 10000 км:")
        fmt.Scanf("%d\n", &distance)

        // Проверка корректности введённого значения
        if distance < 50 || distance > 10000 {
                fmt.Println("Вы ввели недопустимое значение!")
                os.Exit(0)
        }

        fmt.Print("Введите расход в литрах:")
        fmt.Scanf("%d\n", &consumption)

        // Проверка корректности введённого значения
        if consumption < 5 || consumption > 25 {
                fmt.Println("Вы ввели недопустимое значение!")
                os.Exit(0)
        }

        // Рсчёт общего расхода топлива для поездки
        fuelConsumpton := (float64(distance) / float64(100)) * float64(consumption)
        fmt.Printf("Общий расход топлива за поездку:%.f литров\n", fuelConsumpton)

        // Расчёт итоговой стоимости поездки
        totalCost := fuelConsumpton * gasolineCost
        fmt.Printf("Стоимость позедки:%.f рублей\n", totalCost)
}