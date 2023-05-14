// hollow rictangle

package main

import "fmt"

func main() {
	var rows, columns int

	fmt.Print("Rows: ")
	fmt.Scan(&rows)

	fmt.Print("Columns: ")
	fmt.Scan(&columns)

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if i == 0 || i == rows-1 || j == 0 || j == columns-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
