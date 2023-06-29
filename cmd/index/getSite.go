package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/YasserLoukniti/crowlerGo/pkg/protocols"
)

func getSites(buffer []byte, req protocols.GenericRequest, conn net.Conn, database protocols.Database) {

	fmt.Printf("Success => : %v\n", req)
	listReq := protocols.GetSiteRequest{}
	json.Unmarshal([]byte(buffer), &listReq)
	res := protocols.GetSiteResponse{}
	res.Command = "getSite"
	res.Sites = database.Sites
	res.Status = 200
	fmt.Printf("res => : %v\n", res)
	listres_json, err := json.Marshal(res)
	if err != nil {
		println("marshal failed:", err.Error())
		os.Exit(1)
	}

	conn.Write(listres_json)

}
