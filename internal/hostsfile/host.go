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

	// Mapa para almacenar las direcciones IP y sus dominios asociados

	// scan line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Ignorar empty and comment lines
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}
		// Dividir la línea en campos separados por espacios
		fields := strings.Fields(line)
		// La primera palabra es la dirección IP, las siguientes son los dominios asociados
		ip := fields[0]
		hosts := fields[1:]
		// Agregar los dominios al mapa
		IpHosts[ip] = append(IpHosts[ip], hosts...)
	}

	// Comprobar errores en el escaneo del archivo
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al escanear el archivo hosts:", err)
		return
	}

	// Imprimir el mapa de direcciones IP y dominios asociados
	for ip, hosts := range IpHosts {
		fmt.Printf("IP: %s\n", ip)
		fmt.Println("Hosts:")
		for _, host := range hosts {
			fmt.Printf("- %s\n", host)
		}
		fmt.Println()
	}
}

func GetHostEntries() map[string][]string {
	copiedMap := make(map[string][]string)
	for key, value := range IpHosts {
		copiedMap[key] = value
	}
	return copiedMap
}
