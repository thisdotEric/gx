package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func getBranchName(input string) (string, error) {

	re := regexp.MustCompile(`\[(.*?)\]`)
	match := re.FindStringSubmatch(input)

	if len(match) > 1 {
		return match[1], nil
	} else {
		return "", errors.New("Git branch name not found.")
	}

}

func main() {
	// Read input from pipe
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		// Process the first line from the pipe
		line := scanner.Text()

		match, err := getBranchName(line)

		// Not a commit operation
		if err != nil {
			os.Exit(1)
		}

		branchName := strings.Split(match, " ")[0]

		devBranchName := branchName + "-dev"

		// Check if dev branch already exists
		devBranchChecker := exec.Command("git", "show-ref", "refs/heads/"+devBranchName)
		output, _ := devBranchChecker.Output()

		var gitCheckoutCmd *exec.Cmd

		// dev branch already exists
		if string(output) != "" {
			gitCheckoutCmd = exec.Command("git", "checkout", devBranchName)
		} else {
			gitCheckoutCmd = exec.Command("git", "checkout", "-b", devBranchName)
		}

		// Run the command
		checkoutErr := gitCheckoutCmd.Run()
		if checkoutErr != nil {
			fmt.Println("Failed to execute git checkout command", err)
		}

		// Merge the base branch to the dev branch
		gitMergeCommand := exec.Command("git", "merge", branchName)

		// Run the command
		mergeErr := gitMergeCommand.Run()
		if mergeErr != nil {
			fmt.Println("Failed to merge commits. Possible merge conflict error.")
		}

		// Checkout back to the base branch
		gitCheckoutBackCmd := exec.Command("git", "checkout", branchName)

		// Run the command
		checkoutBackErr := gitCheckoutBackCmd.Run()
		if checkoutBackErr != nil {
			fmt.Println("Failed to checkout back to base branch")
		}

		// First line of the git commit output
		fmt.Println(line)
		// Print the rest of the actual output of the 'git commit' command
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
		os.Exit(1)
	}

}
