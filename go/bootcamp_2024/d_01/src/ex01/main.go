package main

import (
	lib "day01/s21_common"
	"flag"
	"fmt"
	"os"
)

const (
	Added = iota
	Removed
	Changed
)

func main() {
	wordPtrNew := flag.String("new", "", "Enter a filename")
	wordPtrOld := flag.String("old", "", "Enter a filename")
	flag.Parse()

	old, err := ParseToStruct(*wordPtrOld)
	ErrorHandler(err)
	new, err := ParseToStruct(*wordPtrNew)
	ErrorHandler(err)

	var oldCakeNames []lib.Cake
	var newCakeNames []lib.Cake

	oldCakeNames = append(oldCakeNames, old.Cakes...)
	newCakeNames = append(newCakeNames, new.Cakes...)

	var commonCakesfromOld []lib.Cake
	var commonCakesfromNew []lib.Cake
	for i := 0; i < len(oldCakeNames); i++ {
		for j := i; j < len(newCakeNames); j++ {
			if oldCakeNames[i].CakeName == newCakeNames[j].CakeName {
				commonCakesfromOld = append(commonCakesfromOld, oldCakeNames[i])
				commonCakesfromNew = append(commonCakesfromNew, newCakeNames[j])
				break
			}
		}
	}

	for _, newCake := range newCakeNames {
		var is_new_exist bool
		for _, commonCake := range commonCakesfromOld {
			if newCake.CakeName == commonCake.CakeName {
				is_new_exist = true
				break
			}
		}
		if !is_new_exist {
			ChechCake(newCake.CakeName, Added)
		}
	}

	for _, oldCake := range oldCakeNames {
		var is_old_exist bool
		for _, commonName := range commonCakesfromOld {
			if oldCake.CakeName == commonName.CakeName {
				is_old_exist = true
				break
			}
		}
		if !is_old_exist {
			ChechCake(oldCake.CakeName, Removed)
		}
	}

	for i := 0; i < len(commonCakesfromOld); i++ {
		timeOld := commonCakesfromOld[i].Time
		timeNew := commonCakesfromNew[i].Time
		if timeOld != timeNew {
			CheckTime(commonCakesfromOld[i], commonCakesfromNew[i], Changed)
		}
		var ItemsOld []lib.Item
		var ItemsNew []lib.Item
		var commonItemsNew []lib.Item
		var commonItemsOld []lib.Item
		ItemsOld = append(ItemsOld, commonCakesfromOld[i].Ingredients...)
		ItemsNew = append(ItemsNew, commonCakesfromNew[i].Ingredients...)

		itemMapCommon := make(map[string]bool)
		for _, oldItem := range ItemsOld {
			for _, newItem := range ItemsNew {
				if oldItem.ItemName == newItem.ItemName {
					itemMapCommon[oldItem.ItemName] = true
					break
				}
			}
		}

		for _, item := range ItemsNew {
			if itemMapCommon[item.ItemName] {
				commonItemsNew = append(commonItemsNew, item)
			} else {
				CheckItemName(commonCakesfromOld[i].CakeName, item.ItemName, Added)
			}
		}

		for _, item := range ItemsOld {
			if itemMapCommon[item.ItemName] {
				commonItemsOld = append(commonItemsOld, item)
			} else {
				CheckItemName(commonCakesfromOld[i].CakeName, item.ItemName, Removed)
			}
		}

		for _, itemOld := range commonItemsOld {
			for _, itemNew := range commonItemsNew {
				if itemOld.ItemName == itemNew.ItemName {
					if itemOld.Unit != itemNew.Unit {
						if itemNew.Unit == "" {
							CheckItemUnit(commonCakesfromOld[i].CakeName, itemOld, itemNew, Removed)
						} else if itemOld.Unit == "" {
							CheckItemUnit(commonCakesfromNew[i].CakeName, itemNew, itemNew, Added)
						} else {
							CheckItemUnit(commonCakesfromOld[i].CakeName, itemOld, itemNew, Changed)
						}
					} else if itemOld.Count != itemNew.Count {
						CheckItemCount(commonCakesfromOld[i].CakeName, itemOld, itemNew, Changed)

					}
				}
			}
		}
	}

}
func CheckItemUnit(cake string, old lib.Item, new lib.Item, mode int) {
	if mode == Changed {
		fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", old.ItemName, cake, new.Unit, old.Unit)
	} else if mode == Removed {
		fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", old.Unit, old.ItemName, cake)
	} else if mode == Added {
		fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake  \"%s\"\n", new.Unit, new.ItemName, cake)
	}
}

func CheckItemCount(cake string, old lib.Item, new lib.Item, mode int) {
	if mode == Changed {
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", old.ItemName, cake, new.Count, old.Count)
	}
}

func CheckItemName(cakeName string, itemName string, mode int) {
	var action string
	if mode == Added {
		action = "ADDED"
	} else if mode == Removed {
		action = "REMOVED"
	}
	fmt.Printf("%s ingredient \"%s\" for cake  \"%s\"\n", action, itemName, cakeName)

}

func ParseToStruct(filename string) (lib.Recipes, error) {
	data, fileExtention, err := lib.GetFile(filename)
	if err != nil {
		return lib.Recipes{}, err
	}

	info, err := lib.GetRecipes(data, fileExtention)
	if err != nil {
		return lib.Recipes{}, err
	}

	return info, nil
}

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ChechCake(cakeName string, mode int) {
	name := cakeName
	action := ""
	if mode == Added {
		action = "ADDED"
	}
	if mode == Removed {
		action = "REMOVED"
	}
	fmt.Printf("%s cake \"%s\"\n", action, name)
}

func CheckTime(old lib.Cake, new lib.Cake, mode int) {
	fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", old.CakeName, new.Time, old.Time)
}
