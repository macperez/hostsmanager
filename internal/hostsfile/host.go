package hostsfile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// const WINDOWS_HOST_DIRECTORY = "C:\\Windows\\System32\\drivers\\etc\\"
const WINDOWS_HOST_DIRECTORY = "C:\\Users\\macastrope\\testhosts\\"

type HostManager struct {
	IpHosts   map[string][]string
	populated bool
}

func (man *HostManager) populate() {
	windows_hosts_path := WINDOWS_HOST_DIRECTORY + "\\hosts"
	fmt.Printf("Reading hosts from %s\n", windows_hosts_path)
	file, err := os.Open(windows_hosts_path)
	if err != nil {
		fmt.Println("Error opening hosts:", err)
		return
	}
	defer file.Close()
	// scan line by line
	scanner := bufio.NewScanner(file)

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
		man.IpHosts[ip] = append(man.IpHosts[ip], hosts...)
		man.populated = true
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning hosts:", err)
		return
	}
}

func (man *HostManager) AddEntry(ip string, host string) {
	man.populate()
	hosts := man.IpHosts[ip]
	hosts = append(hosts, host)
	man.IpHosts[ip] = hosts

}

var manager HostManager

func init() {
	manager = HostManager{
		IpHosts: make(map[string][]string),
	}
}

func Show() {
	if !manager.populated {
		manager.populate()
	}
	for ip, hosts := range manager.IpHosts {
		fmt.Printf("%s --> ", ip)
		for _, host := range hosts {
			fmt.Printf(" | %s", host)
		}
		fmt.Println()
	}
}

func GetHosts(ip string) []string {
	if !manager.populated {
		manager.populate()
	}
	hosts, present := manager.IpHosts[ip]
	if !present {
		return nil
	}
	return hosts
}

func AddEntry(ip string, host string) {
	manager.AddEntry(ip, host)
}

func AddHostsEntries(ip string, hosts []string) {
	// First create backup
	CreateBackup()

	for _, host := range hosts {
		manager.AddEntry(ip, host)
	}
	// save changes in new hosts file
	//SaveChanges()
}
