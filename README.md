
# 3dcart Api Sales Scraper

A series of scripts to QUICKLY pull data from 3dcart/Shift4shop, and save as a json file.




## Installation 

All you need is [Cobra](https://github.com/spf13/cobra)
## How to Use

Find out which integer is the status that you need.

https://apirest.3dcart.com/v1/order-status/index.html#order-status

Note: Order Status look up support will be included in future version. 

## Usage

```golang

go run sales_script go. 

```

## Usage

```bash
$ go run sales_script.go 
Enter Order Status to scrape for(as a intenger): 
10
orderstatus:  10
Is this correct? (y/n)
y
orderstatus_id:  10
Confirmed, searching for orders with orderstatus:  10
Scraping for orders...
Script is done. JSON file is saved in the same directory as the script.

```
