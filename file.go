package main

import(
	"encoding/json"
	"os"
)

func savedata(file string, s *Store) error {
	data, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(file,data,0644)
}

func loaddata(file string) (*Store, error){
	data, err := os.ReadFile(file)
	if err != nil{
		if os.IsNotExist(err){
			return DefaultStore(), nil
		}
		return nil, err
	}

	var s Store
	if err := json.Unmarshal(data, &s); err != nil{
		return nil, err
	}
	return &s, nil
}