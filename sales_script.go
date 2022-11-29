package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {

	/// Check if user has set up environment variables, by reading .env file
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("Please set up environment variables by creating a .env file")
		os.Exit(1)
	}
	// If user has .env file, read it and set environment variables
	if _, err := os.Stat(".env"); err == nil {
		env, err := godotenv.Read()
		if err != nil {
			fmt.Println("Error reading .env file")
			os.Exit(1)
		}
		for k, v := range env {
			os.Setenv(k, v)
		}

		// Set .env variables to variables
		secureURL := os.Getenv("SecureURL")
		privateKey := os.Getenv("PrivateKey")
		token := os.Getenv("Token")

		// Ask user for input of orderstatus_id with cobra
		var orderstatus string
		var cmd = &cobra.Command{
			Use:   "sales_script",
			Short: "sales_script",
			Long:  `sales_script`,
			Run: func(cmd *cobra.Command, args []string) {

				fmt.Println("Enter Order Status to scrape for(as a intenger): ")

				fmt.Scanln(&orderstatus)

				fmt.Println("orderstatus: ", orderstatus)
				// Ask user to confirm orderstatus_id
				fmt.Println("Is this correct? (y/n)")
				var confirm string
				fmt.Scanln(&confirm)
				if confirm == "y" {
					// Save orderstatus_id to variable
					orderstatus_id := orderstatus
					// Print orderstatus_id
					fmt.Println("orderstatus_id: ", orderstatus_id)
					fmt.Println("Confirmed, searching for orders with orderstatus: ", orderstatus)
					// Show progress bar in the terminal
					fmt.Println("Scraping for orders...")

					url := "https://apirest.3dcart.com/3dCartWebAPI/v2/Orders?orderstatus=" + orderstatus

					method := "GET"

					client := &http.Client{}
					req, err := http.NewRequest(method, url, nil)

					if err != nil {
						fmt.Println(err)
						return
					}
					req.Header.Add("SecureURL", secureURL)
					req.Header.Add("PrivateKey", privateKey)
					req.Header.Add("Token", token)

					res, err := client.Do(req)
					if err != nil {
						fmt.Println(err)
						return
					}
					defer res.Body.Close()

					body, err := ioutil.ReadAll(res.Body)
					if err != nil {
						fmt.Println(err)
						return
					}

					// Save response to a JSON file with the status of the order, with a timestamp
					err = ioutil.WriteFile("orders_"+orderstatus_id+".json", body, 0644)

					if err != nil {
						fmt.Println(err)
						return
					}

					// Let user know script is done. And JSON file is saved.
					fmt.Println("Script is done. JSON file is saved in the same directory as the script.")

				} else {
					fmt.Println("Please re-run script")
					return
				}
			},
		}

		cmd.Execute()

	}
}
