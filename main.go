package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/aquasecurity/table"
	"github.com/spf13/cobra"
)

type CustomData map[string]string

const jsonFilePath = "/tmp/custom_values.json"

func getdeployments() []string {
	cmd := exec.Command("kubectl", "get", "deployments")
	var result []string

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error setting the stdout pipe:", err)
		return nil
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting get deployments:", err)
		return nil
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Scan()
	// this is just to skip the header

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		result = append(result, fields[0])
	}
	cmd.Wait()
	return result
}
func loadCustomData() (CustomData, error) {
	data := CustomData{}
	file, err := os.OpenFile(jsonFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		println("can't open nor create file")
		return data, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		println("unable to read the file")

		return data, err
	}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		println("unable to unmarchal json probably cause probably cause the file is empty")

	}
	return data, err
}
func show() {
	// cmd := exec.Command("kubectl ktop", "--flag")
	cmd := exec.Command("kubectl", "get", "deployments")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error setting the stdout pipe:", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting ktop:", err)
		return
	}

	scanner := bufio.NewScanner(stdout)
	customData, err := loadCustomData()
	if err != nil {
		fmt.Println("Error loading JSON:", err, "probably cause the file is empty")

	}
	scanner.Scan()
	// line := scanner.Text()
	// fmt.Printf("%s   %s\n", line, "NOTE")
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("NAME", "READY", "STATUS", "RESTARTS", "AGE", "NOTE")

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		title := fields[0]

		if extraValue, exists := customData[title]; exists {
			// fmt.Printf("%s   %s\n", line, extraValue) // this will add the note column
			table.AddRow(fields[0], fields[1], fields[2], fields[3], fields[4], extraValue)
		} else {
			// fmt.Println(line) // original column
			table.AddRow(fields[0], fields[1], fields[2], fields[3], fields[4], extraValue)
		}
	}

	table.Render()
}

// completion command for auto-completion
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completion scripts",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Root().GenBashCompletion(os.Stdout)
	},
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "deploymentnote [command]",
		Short: "makes notes for deployments",
		Run: func(cmd *cobra.Command, args []string) {
			show()
		},
	}
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(addDeployment)
	rootCmd.AddCommand(delDeployment)
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	completionCmd.Hidden = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
