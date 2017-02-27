package main

import (
	"encoding/json"
	"fmt"
)

type DataPoint []uint64

type GraphiteData struct {
	Target     string      `json:"target"`
	Dataponits []DataPoint `json:"datapoints"`
}

func main() {
	s := `
		[
		    {
				"target": "summarize(log.dGVuZ2luZS1wcm94eWhvc3Q=.YWNjZXNzLnJlcXVlc3Q=, \"1min\", \"sum\")",
				"datapoints": [
					[
					    null,
						1474008900
					],
					[
					    null,
						1474008960
					]
				]
			}
		]
	`

	var graphiteDatas []GraphiteData
	json.Unmarshal([]byte(s), &graphiteDatas)
	fmt.Println("%v", graphiteDatas)

}
