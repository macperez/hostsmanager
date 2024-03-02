package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/macperez/hostsmanager/internal/hostsfile"
	"github.com/spf13/cobra"
)

func main() {
	// Ruta al archivo hosts en Windows
	//var insert bool
	var rootCmd = &cobra.Command{Use: "host"}

	var showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show entries",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			hostsfile.Show()
		},
	}

	var backupCmd = &cobra.Command{
		Use:   "backup",
		Short: "Create backup",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			hostsfile.CreateBackup()
		},
	}
	var ip, host string
	var getHostsByIPCmd = &cobra.Command{
		Use:   "hosts",
		Short: "Get hosts associated to a given IP",
		Run: func(cmd *cobra.Command, args []string) {
			ip, _ := cmd.Flags().GetString("IP")
			hostsFlag, _ := cmd.Flags().GetBool("hosts")
			ip = strings.Trim(ip, " ")
			if hostsFlag {
				hostList := readEntry()
				hostsfile.AddHostsEntries(ip, hostList)

			} else {
				hosts := hostsfile.GetHosts(ip)
				if hosts == nil {
					fmt.Println("This IP is not in hosts file")
				} else {
					for _, host := range hosts {
						fmt.Printf("%s\n", host)
					}
				}
			}

		},
	}
	var hostsInsert bool
	getHostsByIPCmd.Flags().StringVarP(&ip, "IP", "i", " ", "IP format")
	getHostsByIPCmd.Flags().StringVarP(&host, "host", "s", " ", "host")
	getHostsByIPCmd.Flags().BoolVarP(&hostsInsert, "hosts", "", false, "Loop to insert")
	getHostsByIPCmd.MarkFlagRequired("IP")

	/*
		var getMeasurementStationCmd = &cobra.Command{
			Use:   "get-measurements",
			Short: "Get measurements of given station",
			Run: func(cmd *cobra.Command, args []string) {
				from, _ := cmd.Flags().GetString("from")
				to, _ := cmd.Flags().GetString("to")
				to = strings.Trim(to, " ")
				from = strings.Trim(from, " ")
				station, _ := cmd.Flags().GetString("station")
				station = strings.Trim(station, " ")
				err := apirest.GetAemetMeasurements(station, from, to, insert)
				if err != nil {
					fmt.Println(err)
				}

			},
		}
		var station string
		getMeasurementStationCmd.Flags().StringVarP(&fromDate, "from", "f", " ", "Date for which to get the time (format: yyyy-mm-dd)")
		getMeasurementStationCmd.Flags().StringVarP(&toDate, "to", "t", " ", "Date for which to get the time (format: yyyy-mm-dd)")
		getMeasurementStationCmd.Flags().StringVarP(&station, "station", "s", " ", "Code for the station")
		getMeasurementStationCmd.MarkFlagRequired("from")
		getMeasurementStationCmd.MarkFlagRequired("to")
		getMeasurementStationCmd.MarkFlagRequired("station")

		rootCmd.PersistentFlags().BoolVar(&insert, "insert", false, "Insert into database")
	*/
	rootCmd.AddCommand(showCmd, backupCmd, getHostsByIPCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

}

func readEntry() []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Insert IP and host list ('!wq' for save changes)")
	hostList := make([]string, 0)

	for {
		scanner.Scan()

		entry := scanner.Text()

		if strings.EqualFold(entry, "!wq") {
			fmt.Println("Saving changes ... ")
			return hostList

		}

		if strings.EqualFold(entry, "!q") {
			fmt.Println("Exit without saving changes")
			return nil
		}

		// check errors before saving
		hostList = append(hostList, entry)
		// Imprimir la entrada del usuario
		fmt.Printf("\n%s\n", entry)
	}

	// Verificar si hubo errores durante la lectura
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading entry:", err)
	}
	return nil
}
