package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addDeployment = &cobra.Command{
	Use:   "add [title]",
	Short: "add a note to a deployment",
	Args:  cobra.ExactArgs(2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) == 0 {
			deployments := getdeployments()
			var suggestions []string
			for _, deployment := range deployments {
				if strings.HasPrefix(deployment, toComplete) {
					suggestions = append(suggestions, deployment)
				}
			}
			return suggestions, cobra.ShellCompDirectiveNoFileComp
		}
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		addCustomData(args[0], args[1])
		fmt.Println("Note added successfully")
	},
}

func addCustomData(title string, note string) error {

	data := make(map[string]string)
	file, err := os.OpenFile(jsonFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("can't open nor create file")
		return err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("unable to read the file")
		return err
	}

	if len(byteValue) > 0 {
		err = json.Unmarshal(byteValue, &data)
		if err != nil {
			fmt.Println("unable to unmarshal JSON, probably due to an empty or invalid file")
			return err
		}
	}

	data[title] = note

	file.Truncate(0)
	file.Seek(0, 0)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("unable to write to the JSON file")
		return err
	}

	return nil
}
