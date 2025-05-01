package s21_common

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)
// DBReader is a struct
type DBReader interface {
	Read(data []byte) (Recipes, error)
}
// Wrong comment
type Item struct {
	ItemName string `json:"ingredient_name"  xml:"itemname"`
	Count    string `json:"ingredient_count" xml:"itemcount"`
	Unit     string `json:"ingredient_unit"  xml:"itemunit"`
}

type Cake struct {
	CakeName    string `json:"name" xml:"name"`
	Time        string `json:"time" xml:"stovetime"`
	Ingredients []Item `json:"ingredients" xml:"ingredients>item"`
}

type Recipes struct {
	XMLName xml.Name `json:"-" xml:"recipes"`
	Cakes   []Cake   `json:"cake" xml:"cake"`
}

type JSONRecipes struct {
	Data Recipes
}

type XMLRecipes struct {
	Data Recipes
}

func (obj *JSONRecipes) Read(data []byte) (Recipes, error) {
	err := json.Unmarshal(data, &obj.Data)
	if err != nil {
		return Recipes{}, err
	}
	return obj.Data, nil
}

func (obj *XMLRecipes) Read(data []byte) (Recipes, error) {
	err := xml.Unmarshal(data, &obj.Data)
	if err != nil {
		return Recipes{}, err
	}
	return obj.Data, nil
}

func GetFile(filename string) ([]byte, string, error) {
	fileExtention := strings.ToLower(filepath.Ext(filename))
	data, err := os.ReadFile(filename)
	if err != nil {
		return data, fileExtention, err
	}

	return data, fileExtention, nil
}

func ParseFile(filename string) error {
	data, fileExtention, err := GetFile(filename)
	if err != nil {
		return err
	}

	info, err := GetRecipes(data, fileExtention)
	if err != nil {
		return err
	}

	err = PrintReverse(&info, fileExtention)
	if err != nil {
		return err
	}

	return nil
}

func GetRecipes(data []byte, fileExtention string) (Recipes, error) {
	var dbreader DBReader
	if fileExtention == ".json" {
		dbreader = &JSONRecipes{}

	} else if fileExtention == ".xml" {
		dbreader = &XMLRecipes{}

	} else {
		return Recipes{}, fmt.Errorf("wrong extention is used. App supports only .json and .xml")
	}

	return dbreader.Read(data)
}

func PrintReverse(info *Recipes, fileExtention string) error {
	if fileExtention == ".json" {
		xml_packet, err := xml.MarshalIndent(info, "", "    ")
		if err != nil {
			return err
		}
		fmt.Println(string(xml_packet))
	} else if fileExtention == ".xml" {
		json_packet, err := json.MarshalIndent(info, "", "    ")
		if err != nil {
			return err
		}
		fmt.Println(string(json_packet))
	}
	return nil
}
