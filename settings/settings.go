package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sim/models"
)

type Settings struct {
	Shot         models.Shot
	Round        models.Round
	OverallScore models.Score
	Field        models.Field
	Target       models.Target
}

func SaveSettings(s *Settings) {
	byt, _ := json.MarshalIndent(s, "", " ")
	_ = ioutil.WriteFile("./state.json", byt, 0644)
}

func ReadSettings(s *Settings) {
	byt, err := ioutil.ReadFile("./state.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(byt, s)
}
