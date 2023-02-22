package views

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/henriSedjame/passwords/src/models"
	"github.com/olekukonko/tablewriter"
	"os"
)

func ViewError(err string) {
	msg := fmt.Sprintf("⛔ ⛔ ⛔️%s ", err)
	_, _ = color.New(color.BgRed).Add(color.FgBlack).Add(color.Italic).Println(msg)
}

func ViewErrorAndExample(err, example string) {
	ViewError(err)
	space := ""
	for i := 0; i < (len(example) - 5); i++ {
		space = space + " "
	}
	_, _ = color.New(color.BgYellow).Add(color.Bold).Add(color.Underline).Print("Example:")
	_, _ = color.New(color.BgYellow).Add(color.Bold).Println(space)
	_, _ = color.New(color.BgYellow).Add(color.Italic).Println("   " + example)
}

func ViewInfo(msg string) {
	_, _ = color.New(color.BgGreen).Add(color.Italic).Println(" " + msg + " ")
}

func ViewAsTable(passwords ...models.Password) {
	headers := []string{"Label", "Value"}

	var data [][]string

	for _, password := range passwords {
		data = append(data, []string{password.Label, password.Value})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.AppendBulk(data)
	table.Render()
}
