package main

func commandMapb(config *Config) error {

	locationAreas, err := getLocationAreas(config.Prev)
	if err != nil {
		return err
	}

	config.Next = locationAreas.Next
	config.Prev = locationAreas.Previous

	printLocationAreas(locationAreas)

	return nil
}