package hostsfile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const WINDOWS_HOST_PATH = "C:\\Windows\\System32\\drivers\\etc\\hosts"

var IpHosts map[string][]string

func init() {
	IpHosts = make(map[string][]string)
	populate()
}

func populate() {
	file, err := os.Open(WINDOWS_HOST_PATH)
	if err != nil {
		fmt.Println("Error al abrir el archivo hosts:", err)
		return
	}
	defer file.Close()
	// scan line by line
	scanner := bufio.NewScanner(file)
	// Crear un nuevo buffer de bytes sin el BOM

	for scanner.Scan() {
		line := scanner.Text()

		// delete  BOM character
		line = strings.TrimPrefix(line, "\ufeff")
		line = strings.TrimSpace(line)
		// Ignore empty and comment lines
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}
		fields := strings.Fields(line)
		ip := fields[0]
		hosts := fields[1:]
		IpHosts[ip] = append(IpHosts[ip], hosts...)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al escanear el archivo hosts:", err)
		return
	}

}

func GetHostEntries() map[string][]string {
	copiedMap := make(map[string][]string)
	for key, value := range IpHosts {
		copiedMap[key] = value
	}
	return copiedMap
}

func Show() {
	for ip, hosts := range IpHosts {
		fmt.Printf("%s --> ", ip)
		for _, host := range hosts {
			fmt.Printf(" | %s", host)
		}
		fmt.Println()
	}
}
