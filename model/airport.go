package model

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

type Airport struct {
	Name      string `json:"name"`
	City      string `json:"city"`
	Country   string `json:"country"`
	IATA3     string `json:"iata"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

const (
	airportFileName string = "airports"
)

func GetAirportsByRoute(originCity, originCountry, destinationCity, destinationCountry string) ([]string, []string, error) {
	csvFile, err := os.Open("./data/" + airportFileName + ".csv")
	if err != nil {
		return nil, nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var origins []string
	var destinations []string

	isHeader := true
	for {
		line, err := reader.Read()
		if isHeader {
			isHeader = false
			continue
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return origins, destinations, err
		}

		if line[1] == originCity && line[2] == originCountry {
			origins = append(origins, line[3])
		} else if line[1] == destinationCity && line[2] == destinationCountry {
			destinations = append(destinations, line[3])
		}
	}

	return origins, destinations, nil
}

func GetAirportByIATA(iata string) (*Airport, error) {
	csvFile, err := os.Open("./data/" + airportFileName + ".csv")
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))

	isHeader := true
	for {
		line, err := reader.Read()
		if isHeader {
			isHeader = false
			continue
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if line[3] != iata {
			continue
		}

		return &Airport{
			Name:      line[0],
			City:      line[1],
			Country:   line[2],
			IATA3:     line[3],
			Latitude:  line[4],
			Longitude: line[5],
		}, nil
	}

	return nil, nil
}

/*
Name											City			Country				IATA 3	Latitute		Longitude
Goroka Airport									Goroka			Papua New Guinea	GKA		-6.081689835	145.3919983
Madang Airport									Madang			Papua New Guinea	MAG		-5.207079887	145.7890015
Mount Hagen Kagamuga Airport					Mount Hagen		Papua New Guinea	HGU		-5.826789856	144.2960052
Nadzab Airport									Nadzab			Papua New Guinea	LAE		-6.569803		146.725977
Port Moresby Jacksons International Airport		Port Moresby	Papua New Guinea	POM		-9.443380356	147.2200012
Wewak International Airport						Wewak			Papua New Guinea	WWK		-3.583830118	143.6690063
Narsarsuaq Airport								Narssarssuaq	Greenland			UAK		61.16049957		-45.42599869
Godthaab / Nuuk Airport							Godthaab		Greenland			GOH		64.19090271		-51.67810059
Kangerlussuaq Airport							Sondrestrom		Greenland			SFJ		67.0122219		-50.71160316
Thule Air Base									Thule			Greenland			THU		76.53119659		-68.70320129
Akureyri Airport								Akureyri		Iceland				AEY		65.66000366		-18.0727005
Egilsstaðir Airport								Egilsstadir		Iceland				EGS		65.28330231		-14.40139961
Keflavik International Airport					Keflavik		Iceland				KEF		63.98500061		-22.60560036
Patreksfjörður Airport							Patreksfjordur	Iceland				PFJ		65.555801		-23.965
Reykjavik Airport								Reykjavik		Iceland				RKV		64.12999725		-21.94059944
Kugaaruk Airport								Pelly Bay		Canada				YBB		68.534401		-89.808098
Baie Comeau Airport								Baie Comeau		Canada				YBC		49.13249969		-68.20439911
CFB Bagotville									Bagotville		Canada				YBG		48.33060074		-70.99639893
Baker Lake Airport								Baker Lake		Canada				YBK		64.29889679		-96.07779694
Campbell River Airport							Campbell River	Canada				YBL		49.95080185		-125.2710037
Brandon Municipal Airport						Brandon			Canada				YBR		49.91			-99.951897
Cambridge Bay Airport							Cambridge Bay	Canada				YCB		69.10810089		-105.1380005
Nanaimo Airport									Nanaimo			Canada				YCD		49.05497022		-123.8698626
Castlegar/West Kootenay Regional Airport		Castlegar		Canada				YCG		49.29639816		-117.6320038
Miramichi Airport								Chatham			Canada				YCH		47.007801		-65.449203
Charlo Airport									Charlo			Canada				YCL		47.990799		-66.330299
*/
