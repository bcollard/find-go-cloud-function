// git repo: https://github.com/bcollard/find-go-cloud-function
// Runtime: GO 1.13 (beta)
// Date: 05/05/2020

package find

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

// Find is the function entrypoint
func Find(w http.ResponseWriter, r *http.Request) {
	wd, _ := os.Getwd()
	fmt.Fprintf(w, "$> pwd\n%s\n\n", wd)

	bytes, _ := exec.Command("find", "/", "-type", "f", "-name", "hi.txt").CombinedOutput()
	fmt.Fprintf(w, "$> find / -name hi.txt -type f\n%s\n\n", bytes)

	pid := os.Getpid()
	bytes, _ = exec.Command("ls", "-l", "/proc/"+strconv.Itoa(pid)+"/").CombinedOutput()
	fmt.Fprintf(w, "$> ls -l /proc/PID/\n%s\n\n", bytes)

	bytes, _ = exec.Command("cat", "/proc/"+strconv.Itoa(pid)+"/stat").CombinedOutput()
	fmt.Fprintf(w, "$> cat /proc/PID/stat\n%s\n\n", bytes)

	bytes, _ = exec.Command("cat", "/proc/"+strconv.Itoa(pid)+"/cmdline").CombinedOutput()
	fmt.Fprintf(w, "$> cat /proc/PID/cmdline\n%s\n\n", bytes)

	bytes, _ = exec.Command("ls", "-l").CombinedOutput()
	fmt.Fprintf(w, "$> ls -l\n%s\n\n", bytes)

	bytes, _ = exec.Command("ls", "-l", "serverless_function_source_code").CombinedOutput()
	fmt.Fprintf(w, "$> ls -l serverless_function_source_code\n%s\n\n", bytes)

	env := os.Environ()
	fmt.Fprintf(w, "$> env\n%v\n\n", env)

	bytes, _ = exec.Command("ls", "-l", "/workspace").CombinedOutput()
	fmt.Fprintf(w, "$> ls -l /workspace\n%s\n\n", bytes)

	bytes, _ = exec.Command("ls", "-l", "/").CombinedOutput()
	fmt.Fprintf(w, "$> ls -l /\n%s\n\n", bytes)

	bytes, _ = exec.Command("ls", "-l", "serverless_function_source_code/cmd").CombinedOutput()
	fmt.Fprintf(w, "$> ls -l serverless_function_source_code/cmd\n%s\n\n", bytes)

}
