package catalog

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	//"gopkg.in/yaml.v2"
	"encoding/json"
)

var pwd, _ = os.Getwd()

var data = `
catalogItems:
  - product:
      name: Roses
      code: R12
    bundles:
      - quantity: 5
        price: 6.99
      - quantity: 10
        price: 12.99
`

//type Product struct {
//	Name string `yaml:"name"`
//	Code string `yaml:"code"`
//}
//
//type Bundle struct {
//	Quantity int64 `yaml:"quantity"`
//	Price float64 `yaml:"price"`
//}
//
//type CatalogItem struct {
//	Product Product `yaml:"product"`
//    Bundles []Bundle `yaml:"bundles"`
//}
//type Catalog struct {
//	CatalogItems [] CatalogItem `yaml:"catalogItems"`
//}

//func (c *Catalog) GetCatalogFromYamlFile(filePath string) *Catalog {
//
//	//yamlFile, err := ioutil.ReadFile(pwd + "/catalog.yaml")
//	fmt.Println("loading file from " + filePath)
//	yamlFile, err := ioutil.ReadFile(filePath)
//	if err != nil {
//		log.Printf("yamlFile.Get err   #%v ", err)
//	}
//	err = yaml.Unmarshal(yamlFile, c)
//	if err != nil {
//		log.Fatalf("Unmarshal: %v", err)
//	}
//
//	fmt.Println("Catalog=" + c.String())
//	return c
//}

func GetCatalogFromJsonFile(filePath string) []Catalog {
	catalogs := make([]Catalog, 0)
	pwd, err := os.Getwd()

	fmt.Println("loading file from " + pwd + filePath)
	jsonFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("jsonFile.Get err   #%v ", err)
	}
	err = json.Unmarshal(jsonFile, &catalogs)
	if err != nil {
		log.Fatalf("JSON Unmarshal Error: %v", err)
	}

	return catalogs
}

//func displayCatalogFromData() {
//	t := Catalog{}
//
//	err := yaml.Unmarshal([]byte(data), &t)
//	if err != nil {
//		log.Fatalf("error: %v", err)
//	}
//	fmt.Printf("--- t:\n%v\n\n", t)
//
//	d, err := yaml.Marshal(&t)
//	if err != nil {
//		log.Fatalf("error: %v", err)
//	}
//	fmt.Printf("--- t dump:\n%s\n\n", string(d))
//
//	m := make(map[interface{}]interface{})
//
//	err = yaml.Unmarshal([]byte(data), &m)
//	if err != nil {
//		log.Fatalf("error: %v", err)
//	}
//	fmt.Printf("--- m:\n%v\n\n", m)
//
//	d, err = yaml.Marshal(&m)
//	if err != nil {
//		log.Fatalf("error: %v", err)
//	}
//	fmt.Printf("--- m dump:\n%s\n\n", string(d))
//
//	fmt.Println("Product Name: " + t.CatalogItems[0].Product.Name)
//}
