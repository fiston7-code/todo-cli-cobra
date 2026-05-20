package main

import (
	"fmt"
	"os"

	"github.com/fiston7-code/todo-cli-cobra/models"
	"github.com/fiston7-code/todo-cli-cobra/storage"
	"github.com/spf13/cobra"
)

var myStore = models.NewStore()

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A powerful CLI tool built with Cobra",
	Long: `this is a proffesionnal todo list using cli.
   
   This application shows how to create professional command-line
.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Affiche toutes les tâches en attente",
	Run: func(cmd *cobra.Command, args []string) {
		myStore.Show()
	},
}

var addCmd = &cobra.Command{
	Use:   "add [titre]",
	Short: "Ajoute une nouvelle tâche",
	Args:  cobra.MinimumNArgs(1), // Oblige l'utilisateur à donner au moins un titre
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		desc := "Pas de description" // En attendant de gérer les flags plus tard

		// 1. On crée la tâche dans le store en mémoire
		newTask := myStore.CreateTask(title, desc)

		// 2. On sauvegarde la nouvelle liste sur le disque avec ton package storage
		err := storage.Save(myStore.Tasks)
		if err != nil {
			fmt.Printf("Erreur lors de la sauvegarde : %v\n", err)
			return
		}

		fmt.Printf("✅ Tâche \"%s\" ajoutée avec succès (ID: %d) !\n", newTask.Title, newTask.ID)
	},
}

func init() {
	// 1. On attache la commande list
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)

	// 2. On charge les tâches via ton package storage
	existingTasks, err := storage.Load()
	if err != nil {
		fmt.Printf("Erreur lors du chargement des tâches : %v\n", err)
		os.Exit(1)
	}

	// 3. On remplit ton store avec ce qu'on a lu sur le disque
	myStore.SetTasks(existingTasks)
}

func main() {
	// On lance le moteur de Cobra et on récupère l'erreur s'il y en a une
	if err := rootCmd.Execute(); err != nil {
		// Si Cobra détecte une erreur (ex: mauvaise commande), on l'affiche et on quitte
		fmt.Println(err)
		os.Exit(1)
	}

}
