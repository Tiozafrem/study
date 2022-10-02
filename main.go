package main

import (
	"fmt"
	"time"

	"github.com/tiozafrem/study/first"
)

func main() {
	var m, k, p, q float32

	fmt.Print("\n", `
	Средняя заработная плата Z в январе составила M руб., а стоимость S 
	потребительской корзины – K руб. Найти ежемесячную разность Z-S вплоть до декабря 
	в предположении роста Z на p % каждый месяц, а стоимости S – на  q % каждый месяц.`, "\n")

	fmt.Print("\n\tВведите через пробел: m k p q: ")
	fmt.Scanf("%f %f %f %f", &m, &k, &p, &q)

	monthValues := first.SubstractMK(m, k, p, q)

	fmt.Printf("	%-12s|%-10s|\n", "Месяц", "Z-S")
	for monthValue := range monthValues {
		fmt.Printf("	%-12s|%-10.2f|\n", time.Month(monthValue.Month).String(), monthValue.Value)
	}

	fmt.Print("\n", `
	В феврале стоимость потребительской корзины увеличится на q %, 
	а в каждый следующий месяц ее процентное увеличение будет уменьшаться в 2 раза 
	по сравнению с предыдущим месяцем.`, "\n")
	monthValues = first.SubstractAdvancedMK(m, k, p, q)

	fmt.Printf("	%-12s|%-10s|\n", "Месяц", "Z-S")
	for monthValue := range monthValues {
		fmt.Printf("	%-12s|%-10.2f|\n", time.Month(monthValue.Month).String(), monthValue.Value)
	}

	fmt.Print("\n", `
	Стоимость S потребительской корзины в январе составила К руб. Согласно прогнозу, 
	в феврале она увеличится на р %, в марте—на (р+0,1) %,  в апреле—на (р+0,2) %, ..., в декабре—на 
	(р+1) % (по отношению к предыдущему месяцу). Найти среднее значение S за год.`, "\n")
	averageValue := first.AverageValueYear(k, p)
	fmt.Printf("\n\tСредняя значение S за год: %f\n", averageValue)
}
