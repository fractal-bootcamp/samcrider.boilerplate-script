package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// helper function to execute commands
func boundCommand(name string, arg ...string) *exec.Cmd {
	command := exec.Command(name, arg...)

	// bind command to terminal
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command
}

func main() {
	// remove the dist folder if it exists
	err := os.RemoveAll("dist")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the package.json file
	data, err := os.ReadFile("package.json")
	if err != nil {
		fmt.Println("Error reading package.json:", err)
		os.Exit(1)
	}

	// Parse the JSON data
	var pkg map[string]interface{}
	err = json.Unmarshal(data, &pkg)
	if err != nil {
		fmt.Println("Error parsing package.json:", err)
		os.Exit(1)
	}

	// Print the current version
	currentVersion, ok := pkg["version"].(string)
	if !ok {
		fmt.Println("Error parsing version:", err)
		os.Exit(1)
	}
	fmt.Println("Current version:", currentVersion)

	// Prompt for the new version
	fmt.Print("Enter new version: ")
	var newVersion string
	fmt.Scanln(&newVersion)

	// Update the version
	pkg["version"] = newVersion

	// Marshal the updated struct back to JSON
	updatedData, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		fmt.Println("Error creating updated JSON:", err)
		os.Exit(1)
	}

	// Write the updated JSON back to package.json
	err = os.WriteFile("package.json", updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing updated package.json:", err)
		os.Exit(1)
	}

	fmt.Println("Version updated successfully to", newVersion)

	// git add .
	cmd := boundCommand("git", "add", ".")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// git commit -m "Update version to vX.X.X"
	cmd = boundCommand("git", "commit", "-m", "Update version to v"+newVersion)
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// git push origin main
	cmd = boundCommand("git", "push", "origin", "main")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// git tag vX.X.X
	cmd = boundCommand("git", "tag", "-a", "v"+newVersion, "-m", "Release v"+newVersion)
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// run goreleaser release
	cmd = boundCommand("goreleaser", "release")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// npm publish
	cmd = boundCommand("npm", "publish")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("SUCCESS! Update Complete.")

}
