package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
	"todo-app/providers/store"
	"todo-app/tasks"

	"github.com/eiannone/keyboard"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Windows
	} else {
		cmd = exec.Command("clear") // Linux/macOS
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func clearMenu() {
	// Move cursor to the start of the menu area
	fmt.Print("\033[10;1H") // Move cursor to line 10, column 1 (adjust as needed)

	// Clear the menu area by printing empty lines
	for i := 0; i < 30; i++ { // Adjust the number of lines as needed
		fmt.Println("a\033[K") // Clear the current line
	}

	// Move cursor back to the start of the menu area
	fmt.Print("\033[10;1H")
}

func menuNavigation(tasksList []tasks.Task) int {
	options := []string{
		"Add Task",
		"Complete Task",
		"Undo Task",
		"Delete Task",
		"Delete Completed Tasks",
		"Work on Task",
		"Exit",
	}
	selected := 6

	keyboard.Open()
	defer keyboard.Close()

	for {
		clearScreen()

		if len(tasksList) == 0 {
			fmt.Println("\nNo tasks found.")
		}

		fmt.Println("\nTask List:")
		for _, t := range tasksList {
			duration := time.Since(t.CreatedAt).Hours()
			status := "[ ]"
			if t.Completed {
				status = "[✔]"
			}
			fmt.Printf("%s %d: \033[1m%s\033[0m", status, t.ID, t.Title)
			if t.Description != "" {
				fmt.Printf(" - %s", t.Description)
			}
			fmt.Printf(" [Created: %s] [Updated: %s] [Hours since creation: %.2f]\n",
				t.CreatedAt.Format("2006-01-02 15:04:05"), t.UpdatedAt.Format("2006-01-02 15:04:05"), duration)
		}
		fmt.Println()

		fmt.Println("Use UP/DOWN arrows to navigate, ENTER to select:")
		for i, option := range options {
			if i == selected {
				fmt.Printf(" > %s\n", option)
			} else {
				fmt.Printf("   %s\n", option)
			}
		}

		_, Key, _ := keyboard.GetKey()
		if Key == keyboard.KeyArrowUp {
			if selected > 0 {
				selected--
			}
		} else if Key == keyboard.KeyArrowDown {
			if selected < len(options)-1 {
				selected++
			}
		} else if Key == keyboard.KeyEnter {
			return selected + 1
		}
	}
}

func getUserInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	homePath := os.Getenv("HOME")
	var sqlPath string
	if homePath != "" {
		sqlPath = homePath + "/todo.db"
	} else {
		sqlPath = "todo.db"
	}

	sqliteStore, err := store.NewSQLiteStore(sqlPath)
	if err != nil {
		fmt.Println("Error initializing SQLite store:", err)
		return
	}

	for {
		tasksList, err := sqliteStore.ListTasks()
		if err != nil {
			fmt.Println("Error listing tasks:", err)
			return
		}
		choice := menuNavigation(tasksList)
		switch choice {
		case 1: // Add Task
			title := getUserInput("Enter Title: ")
			description := getUserInput("Enter Description: ")
			sqliteStore.CreateTask(title, description)
		case 2: // Complete Task
			idStr := getUserInput("Enter Task ID to Complete: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid Task ID.")
				continue
			}
			sqliteStore.CompleteTask(id)
		case 3: // Undo Task
			idStr := getUserInput("Enter Task ID to Undo: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid Task ID.")
				continue
			}
			sqliteStore.UndoTask(id)
		case 4: // Delete Task
			idStr := getUserInput("Enter Task ID to Delete: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid Task ID.")
				continue
			}
			sqliteStore.DeleteTask(id)
		case 5: // Delete Completed Tasks
			sqliteStore.DeleteCompletedTasks()
		case 6: // Work on Task
			// Implement work on task functionality here
		case 7: // Exit
			fmt.Println("Exiting...")
			return
		}
	}
}
