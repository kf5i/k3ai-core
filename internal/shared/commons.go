package shared

import (
	"fmt"
	"github.com/enescakir/emoji"
	"os"
	"os/exec"
	"strings"
)

// IncludeSlash append the / where needed
func IncludeSlash(path string, typeSeparator string) string {
	if strings.HasSuffix(path, typeSeparator) {
		return path
	}
	return path + typeSeparator
}

//IncludeOsSeparator include os path separator
func IncludeOsSeparator(path string) string {
	return IncludeSlash(path, string(os.PathSeparator))
}

// NormalizePath applies the "/" in the right position
func NormalizePath(file string, args ...string) string {
	result := ""
	for _, subPath := range args {
		result += IncludeOsSeparator(subPath)
	}
	return result + file
}

// NormalizeURL applies the "/" in the right position
func NormalizeURL(args ...string) string {
	result := ""
	for _, subPath := range args {
		result += IncludeSlash(subPath, "/")
	}
	return result
}

// GetDefaultIfEmpty get a value if empty
func GetDefaultIfEmpty(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

// CommandExists check if file exist
func CommandExists(cmd string, osFlavor string, test bool) bool {

	if osFlavor == "windows" {
		// Create an *exec.Cmd
		cmd := exec.Command("bash", "which", cmd)

		// Combine stdout and stderr
		printCommand(cmd)
		output, err := cmd.CombinedOutput()
		printError(err)
		printOutput(output)
		if len(output) > 0 {
			test = true
		} else {
			test = false
		}
	} else {
		var err error
		_, err = exec.LookPath(cmd)
		if err != nil {
			printError(err)
			test = false
		} else {
			test = true
		}
	}
	return test
}

func printCommand(cmd *exec.Cmd) {
	strings.Join(cmd.Args, " ")
}

func printError(err error) {
	if err != nil {
		//os.Stderr.WriteString(err.Error())
		fmt.Printf("Whoops seems we are missing something here..let me fix it for you... %v\n", emoji.Collision)
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf(string(outs))
	}
}

//CheckKubectl check the presence of kubectl binary and in case return true if exist
func CheckKubectl(osFlavor string, filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
