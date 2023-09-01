package second

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/tiozafrem/study/second/model"
	"github.com/tiozafrem/study/second/repository"
	"github.com/tiozafrem/study/second/repository/postgres"
	"github.com/tiozafrem/study/second/services"
)

var Task = `
Известны данные о результатах лыжного забега: фамилии участников, возраст, время старта и время финиша. 
По возрасту выделены 3 возрастные категории, заданные диапазонами. 
Найти чемпиона по каждой возрастной категории. 
Выполнить сортировку списка.
`

var repo *repository.Repository
var ser *services.Services
var delete *bool

func Main() {
	task := flag.Bool("task", false, "Показать задание")
	age := flag.Bool("ages", false, "Управление возрастом")
	peoples := flag.Bool("peoples", false, "Управление людьми")
	delete = flag.Bool("delete", false, "устанавливает режим удаления")
	tournament := flag.Bool("tournament", false, "Управление результатом")
	flag.Parse()
	config := postgres.Config{
		Host:     "localhost",
		Username: "skiers",
		Password: "skiers",
		DBName:   "skiers",
		Port:     "5432",
		SSLMode:  "disable",
	}

	db, err := postgres.NewPostgresDB(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repo = repository.NewRepository(db)
	ser = services.NewServices(repo)
	if *task {
		fmt.Println(Task)
	} else if *age {
		manageAge()
	} else if *peoples {
		managePeople()
	} else if *tournament {
		manageTournament()
	}

}

func manageAge() {
	for {
		fmt.Printf("Список всех возрастов\n")

		ages, err := ser.Age.GetAll(context.Background())
		if err != nil {
			panic(err)
		}

		for i, age := range ages {
			fmt.Printf("%d - %s начальный возраст - %d\n",
				i+1, age.Name, age.AgeStart)
		}

		var id int
		for {
			fmt.Printf("Введите номер возраста для ввода нового значения или пустое значение для выхода: ")
			var input string
			_, err = fmt.Scanln(&input)
			if err != nil {
				return
			}

			_, err = fmt.Sscanf(input, "%d", &id)
			if err != nil {
				fmt.Println("Введено не верное значение")
				continue
			}

			if id < 1 || id > len(ages) {
				fmt.Println("Введено не существующий индекс элемента")
				continue
			}

			id--
			break
		}

		var age int
		for {
			fmt.Printf("Введите новое значение: ")
			var input string
			fmt.Scanln(&input)

			_, err = fmt.Sscanf(input, "%d", &age)
			if err != nil {
				fmt.Println("Введено не верное значение")
				continue
			}

			if age < 1 {
				fmt.Println("Возраст не может быть меньше 1")
				continue
			}
			if age > 122 {
				fmt.Println("Возраст не может быть больше 122")
				continue
			}
			break
		}

		ages[id].AgeStart = age
		err = ser.Age.Update(context.Background(), ages[id])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Изменено успешно\n")
	}
}

func managePeople() {
	for {
		fmt.Printf("Список всех людей\n")

		peoples, err := ser.Person.GetAll(context.Background())
		if err != nil {
			panic(err)
		}

		for i, person := range peoples {

			fmt.Printf("%d - Фамилия %s возраст %d время старта %s время финиша %s  результат %0.2f\n",
				i+1, person.Surname, person.Age, person.TimeStart.Format("15:04:05"),
				person.TimeEnd.Format("15:04:05"), peoples[i].Interval.Minutes())
		}

		var id int
		for {
			fmt.Printf("Введите номер человека для ввода новых значений/удаления или пустое значение для выхода\n")
			if !*delete {
				fmt.Printf("Не должно быть челоевка, с одинаковым возрастом и фамилией одновременно\n")
				fmt.Printf("Для добавления нового человека введите число 0 ")
			}
			fmt.Print(":")
			var input string
			_, err = fmt.Scanln(&input)
			if err != nil {
				return
			}

			_, err = fmt.Sscanf(input, "%d", &id)
			if err != nil {
				fmt.Println("Введено не верное значение")
				continue
			}

			if id < 0 || id > len(peoples) {
				fmt.Println("Введено не существующий индекс элемента")
				continue
			}

			id--
			break
		}

		var people model.Person
		if id >= 0 {
			people = peoples[id]
		}

		if *delete {
			if id < 0 {
				fmt.Fprintf(os.Stderr, "Ошибка: %s\n", "индекс должен быть больше 0")
				continue
			}
			err = ser.Person.Delete(context.Background(), people)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
				os.Exit(1)
			}
			continue
		}

		fmt.Printf("Введите новое значение. Для сохранения предыдущего значения нажмите Enter\n")
		for {
			var input string
			fmt.Printf("Фамилия [%s]: ", people.Surname)
			fmt.Scanln(&input)

			_, err = fmt.Sscanf(input, "%s", &people.Surname)

			if len(input) == 0 && people.Id != 0 {
				break
			}
			if err != nil {
				fmt.Println("Не правильный формат ввода")
				continue
			}

			if len(people.Surname) <= 1 {
				fmt.Println("Фамилия не может быть меньше 2 символов")
				continue
			}
			break
		}

		for {
			var input string
			fmt.Printf("Возраст [%d]: ", people.Age)
			fmt.Scanln(&input)

			_, err = fmt.Sscanf(input, "%d", &people.Age)
			if len(input) == 0 && people.Id != 0 {
				break
			}
			if err != nil {
				fmt.Println("Не правильный формат ввода")
				continue
			}

			if people.Age < 1 {
				fmt.Println("Возраст не может быть меньше 1")
				continue
			}

			if people.Age > 122 {
				fmt.Println("Возраст не может быть больше 122")
				continue
			}
			break
		}

		for {
			var input string
			fmt.Printf("Время старта [%s]: ", people.TimeStart.Format("15:04:05"))
			fmt.Scanln(&input)
			if len(input) == 0 && people.Id != 0 {
				println(input, id)
				break
			}
			timeStart, err := time.Parse("15:04:05", input)
			if err != nil {
				fmt.Println("Введено не верное значение")
				continue
			}
			people.TimeStart = timeStart
			break
		}

		for {
			var input string
			fmt.Printf("Время финиша [%s]: ", people.TimeEnd.Format("15:04:05"))
			fmt.Scanln(&input)

			if len(input) == 0 && people.Id != 0 {
				break
			}
			timeEnd, err := time.Parse("15:04:05", input)
			if err != nil {
				fmt.Println("Введено не верное значение")
				continue
			}
			people.TimeEnd = timeEnd
			break
		}

		method := ser.Person.Create
		if people.Id != 0 {
			method = ser.Person.Update
		}

		err = method(context.Background(), people)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Изменено успешно")
	}
}

func manageTournament() {
	for {
		fmt.Printf("Список всех возрастов\n")

		ages, err := ser.Age.GetAll(context.Background())
		if err != nil {
			panic(err)
		}

		for i, age := range ages {
			fmt.Printf("%d - %s начальный возраст - %d\n",
				i+1, age.Name, age.AgeStart)
		}

		var id int
		for {
			fmt.Printf("Введите номер возраста для вывода турнирного списка: ")
			var input string
			_, err = fmt.Scanln(&input)
			if err != nil {
				return
			}

			_, err = fmt.Sscanf(input, "%d", &id)
			if err != nil {
				fmt.Println("Введено не верное значение")
				continue
			}

			if id < 1 || id > len(ages) {
				fmt.Println("Введено не существующий индекс элемента")
				continue
			}
			id--
			break
		}
		fmt.Printf("\n\n")
		persons, err := ser.Person.GetChampionsByAge(context.Background(), ages[id])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
			os.Exit(1)
		}

		for i := 0; i < len(persons); i++ {
			fmt.Printf("Чемпион - %s %d со временем %.2f\n", persons[i].Surname, persons[i].Age, persons[i].Interval.Minutes())
		}
	}
}
