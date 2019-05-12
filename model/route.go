package model

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

type Route struct {
	AirlineID   string
	Origin      string
	Destination string
}

const (
	routeFileName string = "routes"
)

func GetAllRoutes() ([]Route, error) {
	csvFile, err := os.Open("./data/" + routeFileName + ".csv")
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var routes []Route

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
			return routes, err
		}

		routes = append(routes, Route{
			AirlineID:   line[0],
			Origin:      line[1],
			Destination: line[2],
		})
	}

	return routes, nil
}

/*
Airline Id	Origin	Destination
AC			ABJ		BRU
AC			ABJ		OUA
AC			ADD		JED
AC			AMS		CPH
AC			ANU		YYZ
AC			ATL		DEN
AC			ATL		YYZ
AC			AUA		YYZ
AC			AUH		YYZ
AC			AZS		YUL
AC			AZS		YYZ
AC			BAH		DOH
AC			BCN		YYZ
AC			BDA		YYZ
CZ			AVA		CAN
CZ			BAV		CGO
CZ			BAV		CSX
CZ			BAV		CTU
CZ			BAV		SHE
CZ			BAV		SJW
CZ			BAV		URC
CZ			BAV		WUH
CZ			BFJ		CAN
CZ			BFJ		SZX
CZ			BHY		CAN
CZ			BHY		CGO
CZ			BHY		CKG
CZ			BHY		CSX
CZ			BHY		KMG
TK			ADD		IST
TK			ADD		JUB
TK			ADD		SSG
TK			ADE		IST
TK			ADF		ESB
TK			ADF		IST
TK			AER		IST
TK			AGP		IST
TK			AJI		ESB
TK			AJI		IST
TK			AKL		BKK
TK			ALA		IST
TK			ALG		IST
TK			AMM		IST
TK			AMS		IST
TK			AMS		SAW
TK			AOE		BRU
TK			AQJ		IST
UA			ALB		ORD
UA			ALS		DEN
UA			ALS		FMN
UA			AMA		DEN
UA			AMA		IAH
UA			AMS		EWR
UA			AMS		IAD
UA			AMS		IAH
UA			AMS		ORD
UA			ANC		DEN
UA			ANC		ORD
UA			ANC		SEA
UA			ANU		EWR
UA			AOO		IAD
UA			AOO		JST
WN			ATL		DTW
WN			ATL		FLL
WN			ATL		HOU
WN			ATL		IND
WN			ATL		JAX
WN			ATL		LAS
WN			ATL		LAX
WN			ATL		LGA
WN			ATL		MBJ
WN			ATL		MCI
WN			ATL		MCO
WN			ATL		MDW
WN			ATL		MKE
WN			ATL		MSP
WN			ATL		MSY
WN			ATL		NAS
WN			ATL		OKC
WN			ATL		ORF
WN			ATL		PBI
WN			ATL		PHL
WS			JFK		YUL
WS			JFK		YYC
WS			JFK		YYZ
WS			KIN		YYZ
WS			KOA		YVR
WS			LAS		YEG
WS			LAS		YQR
WS			LAS		YUL
WS			LAS		YVR
WS			LAS		YWG
WS			LAS		YXE
WS			LAS		YYC
WS			LAS		YYJ
WS			LAS		YYZ
WS			LAX		YEG
WS			LAX		YVR
WS			LAX		YYC
WS			LAX		YYZ
WS			LGA		ATL
*/
