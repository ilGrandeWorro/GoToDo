package main

import ("fmt")
var toDoList = map[string]bool{}
var keyList = []string{}
func main() {
	

	fmt.Println("\nLista creata:")
	printList()

	addTask("Creare una to Do List")
	addTask("Crea gioco 'Indovina il numero'")
	addTask("Creare una app meteo")
	addTask("Creare una app Chat")
	addTask("Creare una to Blog personale")
	addTask("CANCELLARE questa task")
	addTask("Creare una calcolatrice")
	addTask("Cripta e decripta un file")
	addTask("Creare un music player")
	addTask("Creare un convertitore di valuta")
	addTask("Creare Space invaders")
	addTask("CANCELLARE questa task")
	addTask("MODIFICA QUESTA TASK")
	fmt.Println("\nDopo aver aggiunto gli elementi:")
	printList()

	deleteTask(11)
	deleteTask(5)
	fmt.Println("\nDopo aver cancellato degli elementi:")
	printList()

	updateTask(10,"AGGIORNATA")
	fmt.Println("\nDopo aver aggiornato la task:")
	printList()

	setAsDone(10)
	fmt.Println("\nDopo aver settato la task 10 come 'completa':")
	printList()

	setAsDone(0)
	fmt.Println("\nE anche la task di creare la toDo list!!!")
	printList()
}

func addTask(task string) {
	toDoList[task]=false
	keyList = append(keyList, task)
}

func deleteTask(index int){
	var key string = keyList[index]
	delete(toDoList, key)
	keyList = append(keyList[:index], keyList[index+1:]...)
}



func printList(){
		
		for _, key := range keyList{
			fmt.Printf("%v = %v\n", key, toDoList[key])
		}
}

func updateTask(index int, newTask string){
	deleteTask(index)
	addTask(newTask)
}

func setAsDone(index int){
	var key string = keyList[index]
	toDoList[key] = true
}