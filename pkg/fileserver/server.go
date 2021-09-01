package fileserver

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func openbrowser(url string) {
	fmt.Printf("Opening at %s...\n", url)
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func ServeFiles(outputDir string, filename string, port int) {
	fmt.Printf("server path: %s\n", outputDir)
	http.Handle("/", http.FileServer(http.Dir(outputDir)))
	previewUrl := fmt.Sprintf("http://localhost:%d/%s", port, filename)
	openbrowser(previewUrl)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
