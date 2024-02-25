package hostsfile

import (
	"fmt"
	"io"
	"os"
	"time"
)

const HOST_FILE_BACKUP_PREFIX = "hosts_bk"

func CreateBackup() {

	now := time.Now()
	formatted := now.Format("20060102_150405") // YYYYMMDD_HHMMSS
	fileName := WINDOWS_HOST_DIRECTORY + "\\" + HOST_FILE_BACKUP_PREFIX + "_" + formatted
	windows_hosts_path := WINDOWS_HOST_DIRECTORY + "\\hosts"
	// Read
	orig, err := os.Open(windows_hosts_path)
	if err != nil {
		fmt.Println("Error opening:", err)
		return
	}
	defer orig.Close()

	// Crear el archivo de destino
	dest, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating copy:", err)
		return
	}
	defer dest.Close()

	_, err = io.Copy(dest, orig)
	if err != nil {
		fmt.Println("Error copying:", err)
		return
	}

	fmt.Println("Backup created")

}
