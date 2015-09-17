package main

import "encoding/xml"
import "fmt"
import "strings"
type Resource struct{
			//SurfaceForm string `xml:"surfaceForm,attr"`
			URI string `xml:"URI,attr"1`
	}

func main(){
	
	entities := make(map[string]int)
	type wrapper struct{
			Resource []Resource `xml:"Resources>Resource"`
	}
	var temp wrapper

	err := xml.Unmarshal( []byte(xmlTest), &temp)
	if err != nil {
		fmt.Println("Entities xml Unmarshall Error:", err)
	}
	//Need to remove duplicates and get Categories
	for _,URI :=range temp.Resource{
		_, prs := entities[URI.URI]
		if !prs{
			split:=strings.Split(URI.URI,`/`)
			
			entities[split[len(split)-1]]+=1
		}
	}
	
}


const xmlTest = `<?xml version="1.0" encoding="utf-8"?>
<Annotation>
<Resources>
<Resource URI="http://dbpedia.org/resource/The_Timberland_Company" support="91" types="DBpedia:Agent,Schema:Organization,DBpedia:Organisation,DBpedia:Company" surfaceForm="Timberland" offset="14" similarityScore="0.9999999998713065" percentageOfSecondRank="1.286755919703084E-10"/>
<Resource URI="http://dbpedia.org/resource/The_Timberland_Company" support="91" types="DBpedia:Agent,Schema:Organization,DBpedia:Organisation,DBpedia:Company" surfaceForm="Timberland" offset="350" similarityScore="0.9999999998713065" percentageOfSecondRank="1.286755919703084E-10"/>
<Resource URI="http://dbpedia.org/resource/Generation_Y" support="254" types="" surfaceForm="millennials" offset="490" similarityScore="1.0" percentageOfSecondRank="0.0"/>
<Resource URI="http://dbpedia.org/resource/The_Timberland_Company" support="91" types="DBpedia:Agent,Schema:Organization,DBpedia:Organisation,DBpedia:Company" surfaceForm="Timberland" offset="507" similarityScore="0.9999999998713065" percentageOfSecondRank="1.286755919703084E-10"/>
<Resource URI="http://dbpedia.org/resource/The_Timberland_Company" support="91" types="DBpedia:Agent,Schema:Organization,DBpedia:Organisation,DBpedia:Company" surfaceForm="Timberland" offset="888" similarityScore="0.9999999998713065" percentageOfSecondRank="1.286755919703084E-10"/>
<Resource URI="http://dbpedia.org/resource/Lookbook" support="10" types="" surfaceForm="lookbooks" offset="1225" similarityScore="1.0" percentageOfSecondRank="0.0"/>
<Resource URI="http://dbpedia.org/resource/The_Timberland_Company" support="91" types="DBpedia:Agent,Schema:Organization,DBpedia:Organisation,DBpedia:Company" surfaceForm="Timberland" offset="1362" similarityScore="0.9999999998713065" percentageOfSecondRank="1.286755919703084E-10"/>
<Resource URI="http://dbpedia.org/resource/VF_Corporation" support="69" types="DBpedia:Agent,Schema:Organization,DBpedia:Organisation,DBpedia:Company" surfaceForm="VF Corp." offset="1619" similarityScore="1.0" percentageOfSecondRank="0.0"/>
<Resource URI="http://dbpedia.org/resource/The_Timberland_Company" support="91" types="DBpedia:Agent,Schema:Organization,DBpedia:Organisation,DBpedia:Company" surfaceForm="Timberland" offset="1674" similarityScore="0.9999999998713065" percentageOfSecondRank="1.286755919703084E-10"/>
<Resource URI="http://dbpedia.org/resource/United_Kingdom" support="170647" types="Schema:Place,DBpedia:Place,DBpedia:PopulatedPlace,Schema:Country,DBpedia:Country" surfaceForm="British" offset="1743" similarityScore="0.9999897378132996" percentageOfSecondRank="6.265051588227626E-6"/>
<Resource URI="http://dbpedia.org/resource/Digital_video" support="488" types="" surfaceForm="digital videos" offset="1910" similarityScore="1.0" percentageOfSecondRank="0.0"/>
<Resource URI="http://dbpedia.org/resource/The_Timberland_Company" support="91" types="DBpedia:Agent,Schema:Organization,DBpedia:Organisation,DBpedia:Company" surfaceForm="Timberland" offset="2308" similarityScore="0.9999999998713065" percentageOfSecondRank="1.286755919703084E-10"/>
<Resource URI="http://dbpedia.org/resource/Teen_Vogue" support="196" types="Schema:CreativeWork,DBpedia:Work,DBpedia:WrittenWork,DBpedia:PeriodicalLiterature,DBpedia:Magazine" surfaceForm="Teen Vogue" offset="2491" similarityScore="1.0" percentageOfSecondRank="0.0"/>
<Resource URI="http://dbpedia.org/resource/Nylon_(magazine)" support="176" types="Schema:CreativeWork,DBpedia:Work,DBpedia:WrittenWork,DBpedia:PeriodicalLiterature,DBpedia:Magazine" surfaceForm="Nylon" offset="2503" similarityScore="0.9999669204003921" percentageOfSecondRank="3.308069393776717E-5"/>
<Resource URI="http://dbpedia.org/resource/People_(magazine)" support="2431" types="Schema:CreativeWork,DBpedia:Work,DBpedia:WrittenWork,DBpedia:PeriodicalLiterature,DBpedia:Magazine" surfaceForm="People StyleWatch" offset="2510" similarityScore="1.0" percentageOfSecondRank="0.0"/>
<Resource URI="http://dbpedia.org/resource/InStyle" support="197" types="Schema:CreativeWork,DBpedia:Work,DBpedia:WrittenWork,DBpedia:PeriodicalLiterature,DBpedia:Magazine" surfaceForm="InStyle" offset="2529" similarityScore="1.0" percentageOfSecondRank="3.3375378394918295E-42"/>
<Resource URI="http://dbpedia.org/resource/Marie_Claire" support="549" types="Schema:CreativeWork,DBpedia:Work,DBpedia:WrittenWork,DBpedia:PeriodicalLiterature,DBpedia:Magazine" surfaceForm="Marie Claire" offset="2541" similarityScore="1.0" percentageOfSecondRank="1.0208950968881831E-50"/>
<Resource URI="http://dbpedia.org/resource/VF_Corporation" support="69" types="DBpedia:Agent,Schema:Organization,DBpedia:Organisation,DBpedia:Company" surfaceForm="VF Corp." offset="2828" similarityScore="1.0" percentageOfSecondRank="0.0"/>
<Resource URI="http://dbpedia.org/resource/Public_relations" support="3077" types="" surfaceForm="public relations" offset="3275" similarityScore="0.9999999999333795" percentageOfSecondRank="6.663020510564207E-11"/>
</Resources>
</Annotation>
`