package docops

import (
	"os"
	"fmt"
	"errors"
	"strconv"
)

// read from the account balance file
// added error as a type so i could use the errors package to generate custom error messages
// to use a function within a package in a different package, it must begin with an uppercase character
func RetrieveFloatFromFile(file string) (float64, error) {
	data, err := os.ReadFile(file)

	// error handling
	if err != nil {
		return 700, errors.New("The intended file could not be found.")
	}

	valueTxt := string(data) // convert the data that's converted to byte using writetofile to string format
	value, err := strconv.ParseFloat(valueTxt, 64) // convert the string data back to float64
	// error handling for parsefloat as it could also generate an error
	if err != nil {
		return 700, errors.New("Could not parse the stored value to float.")
	}

	// add nil to the return to show there's no error
	return value, nil
}

// persist the account balance to a file
func PersistFloatToDoc(value float64, file string) {
	valueTxt := fmt. Sprint(value)
	os.WriteFile(file, []byte(valueTxt), 644)
}