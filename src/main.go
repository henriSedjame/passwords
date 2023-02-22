package main

import (
	"flag"
	"fmt"
	"github.com/henriSedjame/passwords/src/messages"
	"github.com/henriSedjame/passwords/src/models"
	"github.com/henriSedjame/passwords/src/views"
	"os"
)

const fileName = ".passwords.json"

func main() {

	add := flag.Bool(models.Add, false, models.AddDesc)
	update := flag.Bool(models.Update, false, models.UpdateDesc)
	show := flag.Bool(models.Show, false, models.ShowDesc)
	del := flag.Bool(models.Delete, false, models.DeleteDesc)
	list := flag.Bool(models.List, false, models.DeleteDesc)
	label := flag.String(models.Label, "", models.LabelDesc)
	value := flag.String(models.Value, "", models.ValueDesc)

	flag.Parse()

	passwords := &models.Passwords{}

	err := passwords.Load(fileName)
	if os.IsNotExist(err) {
		f, err := os.Create(fileName)
		defer func() {
			_ = f.Close()
		}()
		handleError(err)
	}

	switch {
	case *add:
		if *label != "" && *value != "" {
			applyFns(
				func() error {
					return passwords.Add(*label, *value)
				},
				func() error {
					err := passwords.Store(fileName)
					if err == nil {
						views.ViewInfo(fmt.Sprintf(messages.AddSuccess, *label))
					}
					return err
				},
			)
		} else {
			views.ViewErrorAndExample(messages.AddLabelAndValueMissing, messages.AddExample)
		}
	case *update:
		if *label != "" && *value != "" {
			applyFns(
				func() error {
					return passwords.Update(*label, *value)
				},
				func() error {
					err := passwords.Store(fileName)
					if err == nil {
						views.ViewInfo(fmt.Sprintf(messages.UpdateSuccess, *label))
					}
					return err
				},
			)
		} else {
			views.ViewErrorAndExample(messages.UpdateLabelAndValueMissing, messages.UpdateExample)
		}
	case *show:
		if *label != "" {
			applyFns(
				func() error {
					pass, err := passwords.Find(*label)
					if err == nil {
						views.ViewAsTable(models.Password{
							Label: *label,
							Value: pass,
						})
					}
					return err
				})

		} else {
			views.ViewErrorAndExample(messages.ShowLabelMissing, messages.ShowExample)
		}
	case *del:
		if *label != "" {
			applyFns(
				func() error {
					return passwords.Delete(*label)
				},
				func() error {
					err := passwords.Store(fileName)
					if err == nil {
						views.ViewInfo(fmt.Sprintf(messages.DeleteSuccess, *label))
					}
					return err
				},
			)
		} else {
			views.ViewErrorAndExample(messages.DeleteLabelMissing, messages.DeleteExample)
		}
	case *list:
		views.ViewAsTable(*passwords...)
	}

}

func handleError(err error) {
	if err != nil {
		views.ViewError(err.Error())
	}
}

func applyFns(fns ...func() error) {
	for _, fn := range fns {
		if err := fn(); err != nil {
			handleError(err)
		}
	}
}
