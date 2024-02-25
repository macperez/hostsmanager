package hostsfile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const WINDOWS_HOST_DIRECTORY = "C:\\Windows\\System32\\drivers\\etc\\"

var IpHosts map[string][]string

func init() {
	IpHosts = make(map[string][]string)

}

func populate() {
	windows_hosts_path := WINDOWS_HOST_DIRECTORY + "\\hosts"
	file, err := os.Open(windows_hosts_path)
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
		//val, ok := IpHosts[ip]
		IpHosts[ip] = append(IpHosts[ip], hosts...)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning hosts:", err)
		return
	}

}

func Show() {
	populate()
	for ip, hosts := range IpHosts {
		fmt.Printf("%s --> ", ip)
		for _, host := range hosts {
			fmt.Printf(" | %s", host)
		}
		fmt.Println()
	}
}

func AddEntry(ip string, hostsList []string) {
	IpHosts[ip] = hostsList
}

func GetHosts(ip string) []string {
	populate()
	hosts, present := IpHosts[ip]
	if !present {
		return nil
	}
	return hosts
}
