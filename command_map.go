package main

func commandMap(config *Config) error {

	locationAreas, err := getLocationAreas(config.Next)
	if err != nil {
		return err
	}

	config.Next = locationAreas.Next
	config.Prev = locationAreas.Previous

	printLocationAreas(locationAreas)

	return nil
}

