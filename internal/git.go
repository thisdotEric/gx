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
	branchChecker := exec.Command("git", "show-ref", "refs/heads/"+branch)
	output, _ := branchChecker.Output()

	return string(output) != ""
}

func getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "symbolic-ref", "--short", "HEAD")

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get current branch: %w", err)
	}

	branch := strings.TrimSpace(string(output))
	return branch, nil
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

func getDefaultCheckoutBranch(args []string) string {
	var checkoutBranch = "dev"

	if len(args) > 1 {
		checkoutBranch = args[1:][0]
	}

	return checkoutBranch
}

func processGitCommands(baseBranchName, targetCheckoutBranch string) error {

	var gitCheckoutCmd *exec.Cmd

	if branchExists(targetCheckoutBranch) {
		gitCheckoutCmd = createGitCheckoutBranchCmd(targetCheckoutBranch, false)
	} else {
		gitCheckoutCmd = createGitCheckoutBranchCmd(targetCheckoutBranch, true)
	}

	err := gitCheckoutCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute git checkout command, make sure to commit all changes first")
	}

	gitMergeCommand := createGitMergeBranchCmd(baseBranchName)

	err = gitMergeCommand.Run()
	if err != nil {
		return fmt.Errorf("failed to merge commits, possible merge conflict error")
	}

	// Checkout back to the base branch
	gitCheckoutBackCmd := createGitCheckoutBranchCmd(baseBranchName, false)

	err = gitCheckoutBackCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to checkout back to base branch")
	}

	return nil
}

func HandlePipeInput(args []string) error {

	checkoutBranch := getDefaultCheckoutBranch(args)

	// Read input from pipe
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		// Process the first line from the pipe
		line := scanner.Text()
		match, err := getBranchName(line)

		// Not a commit operation
		if err != nil {
			return err
		}

		branchName := strings.Split(match, " ")[0]
		targetCheckoutBranch := branchName + "-" + checkoutBranch

		err = processGitCommands(branchName, targetCheckoutBranch)

		if err != nil {
			return err
		}

		fmt.Println(line)
		// Print the rest of the actual output of the 'git commit' command
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading standard input: %w", err)
	}

	return nil

}

func HandleCLIInput(args []string) error {

	checkoutBranch := getDefaultCheckoutBranch(args)
	// Get the current branch
	baseBranchName, err := getCurrentBranch()

	if err != nil {
		return nil
	}

	targetBranchName := baseBranchName + "-" + checkoutBranch

	err = processGitCommands(baseBranchName, targetBranchName)

	if err != nil {
		return nil
	}

	fmt.Println("changes merged to", targetBranchName)
	return nil
}
