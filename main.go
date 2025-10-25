package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"github.com/soft2help/php-wrapper/internal/php"
)

const (
	PHPWrapperWebsite  = "https://github.com/soft2help/php-wrapper"
	downloadURL   = "https://windows.php.net/downloads/releases/php-8.3.27-Win32-vs16-x64.zip"
	installFolder = ".s2hWS"
)

func main() {

	var phpExe string
	var phpFolder string
	var phpExeAbs string

	var err error

	host := flag.String("S", "localhost:8080", "<addr>:<port> - Run with built-in web server.")
	docRoot := flag.String("t", "public_html", "<docroot> - Specify document root <docroot> for built-in web server.")
	routerFileOnPublicHtml := flag.String("r", "index.php", "<router> - Specify files that will redirect.")
	flag.Parse()
	s := []string{*docRoot, "\\", *routerFileOnPublicHtml}
	router := strings.Join(s, "")
	fmt.Println(*routerFileOnPublicHtml)
	fmt.Println(router)


	phpExe, err = php.FindPHPExec(installFolder)
	if err != nil {

		phpExe, err = php.FastInstall(downloadURL, installFolder)

		if err != nil {
			fmt.Printf("Unable to install PHP: %v", err)
			php.ExitPHPWrapper()
		}
	}

	phpExeAbs, _ = filepath.Abs(phpExe)
	phpFolder, _ = filepath.Split(phpExeAbs)
	sep := fmt.Sprintf("%c", filepath.Separator)
	phpFolder = strings.TrimSuffix(phpFolder, sep)
	docRootAbs, _ := filepath.Abs(*docRoot)

	if _, err := os.Stat(*docRoot); os.IsNotExist(err) {
		os.Mkdir(*docRoot, 0755)
	}

	fmt.Println("")
	fmt.Println("** Information **")
	fmt.Println("Copy your PHP files in the Document root folder.")
	fmt.Println("")
	fmt.Printf("%-16s | %s\n", "PHP located in", phpExeAbs)
	fmt.Printf("%-16s | %s\n", "PHP Folder in", phpFolder)
	fmt.Printf("%-16s | %s\n", "Document root", docRootAbs)
	fmt.Printf("%-16s | %s\n", "Router file", router)
	fmt.Printf("%-16s | %s\n", "Website URL", fmt.Sprintf("%s%s", "http://", *host))
	fmt.Println("")

	cmdHideFolder := exec.Command("attrib", "+s", "+h", phpFolder)
	cmdErrHideFolder := cmdHideFolder.Run()
	if cmdErrHideFolder != nil {
		log.Fatal(cmdErrHideFolder)
	}

	args := []string{"-S", *host, "-t", *docRoot}
	if !(len(router) == 0) {
		args = append(args, router)
	}

	cmd := exec.Command(phpExe, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmdErr := cmd.Run()
	if cmdErr != nil {
		log.Fatal(err)
	}

}
