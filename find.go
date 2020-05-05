package find

import (
	"fmt"
	"net/http"
	"os/exec"
)

func Find(w http.ResponseWriter, r *http.Request) {
	bytes, _ := exec.Command("find", "/", "-type", "f", "-name", "hi.txt").CombinedOutput()
	fmt.Fprintf(w, "%s", bytes)
	// Output: /workspace/serverless_function_source_code/hi.txt
	// Runtime: GO 1.13 (beta)
	// Date: 05/05/2020
}
