package internal

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
		return "", errors.New("git branch name not found")
	}

}

func branchExists(branch string) bool {
	// Check if dev branch already exists
	devBranchChecker := exec.Command("git", "show-ref", "refs/heads/"+branch)
	output, _ := devBranchChecker.Output()

	return string(output) != ""
}

func createGitCheckoutBranchCmd(targetBranch string, createNewBranch bool) *exec.Cmd {
	if createNewBranch {
		return exec.Command("git", "checkout", "-b", targetBranch)
	}

	return exec.Command("git", "checkout", targetBranch)
}

func createGitMergeBranchCmd(sourceBranch string) *exec.Cmd {
	return exec.Command("git", "merge", sourceBranch)
}

func HandlePipeInput(targetBranch string) {
	var checkoutBranch = "dev"

	if targetBranch != "" {
		checkoutBranch = targetBranch
	}

	// Read input from pipe
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		// Process the first line from the pipe
		line := scanner.Text()

		match, err := getBranchName(line)

		// Not a commit operation
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		branchName := strings.Split(match, " ")[0]

		targetCheckoutBranch := branchName + "-" + checkoutBranch

		var gitCheckoutCmd *exec.Cmd

		if branchExists(targetCheckoutBranch) {
			gitCheckoutCmd = createGitCheckoutBranchCmd(targetCheckoutBranch, false)
		} else {
			gitCheckoutCmd = createGitCheckoutBranchCmd(targetCheckoutBranch, true)
		}

		// Run the command
		checkoutErr := gitCheckoutCmd.Run()
		if checkoutErr != nil {
			fmt.Println("Failed to execute git checkout command", err)
			os.Exit(1)
		}

		// Merge the base branch to the dev branch
		gitMergeCommand := createGitMergeBranchCmd(branchName)

		// Run the command
		mergeErr := gitMergeCommand.Run()
		if mergeErr != nil {
			fmt.Println("Failed to merge commits. Possible merge conflict error.")
			os.Exit(1)
		}

		// Checkout back to the base branch
		gitCheckoutBackCmd := createGitCheckoutBranchCmd(branchName, false)

		// Run the command
		checkoutBackErr := gitCheckoutBackCmd.Run()
		if checkoutBackErr != nil {
			fmt.Println("Failed to checkout back to base branch")
			os.Exit(1)
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

func HandleCLIInput(args []string) {

	var checkoutBranch = "dev"

	if len(args) > 1 {
		checkoutBranch = args[1:][0]
		fmt.Println(checkoutBranch)
	}

	HandlePipeInput("someone")

}
