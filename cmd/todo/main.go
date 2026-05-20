package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fiston7-code/todo-cli-cobra/models"
	"github.com/fiston7-code/todo-cli-cobra/storage"
	"github.com/spf13/cobra"
)

var myStore = models.NewStore()
var taskDesc string

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
		desc := taskDesc

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

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Supprime une tâche grâce à son ID",
	Args:  cobra.ExactArgs(1), // Oblige l'utilisateur à donner exactement 1 argument (l'ID)
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Convertir l'argument (string) en entier (int)
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("❌ L'ID doit être un nombre valide.")
			return
		}

		// 2. Appeler ta méthode DeleteTask du store
		err = myStore.DeleteTask(id)
		if err != nil {
			fmt.Printf("❌ Erreur : %v\n", err)
			return
		}

		// 3. Sauvegarder les modifications sur le disque
		err = storage.Save(myStore.Tasks)
		if err != nil {
			fmt.Printf("Erreur lors de la sauvegarde : %v\n", err)
			return
		}

		fmt.Printf("🗑️ Tâche %d supprimée avec succès !\n", id)
	},
}

func init() {
	// 1. On attache la commande list
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)

	// On lie le flag à notre variable taskDesc
	addCmd.Flags().StringVarP(&taskDesc, "desc", "d", "", "La description de la tâche")

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
