package main

import (
	"fmt"
	"math/rand"
)

type Patient struct {
	PatientID      int64  `json:"patient_id`
	PatientName    string `json:"patient_name"`
	PatientDisease string `json: patient_disease"`
	PatientAge     int64  `json:"patient_age"`
}

var patients []Patient

func addPatient(patient Patient) {
	patients = append(patients, patient)
	fmt.Println(patients)
}

func main() {

	for {
		fmt.Println("Add Patient")
		var add string
		fmt.Scanln(&add)

		if add != "yes" {
			fmt.Println(patients)
			break
		} else {

			fmt.Println("Patient name :")
			var _name string
			fmt.Scanln(&_name)
			fmt.Println("Patient Disease :")
			var _disease string
			fmt.Scanln(&_disease)
			fmt.Println("Patient Age :")
			var _age int64
			fmt.Scanln(&_age)

			var patient Patient
			patient.PatientID = rand.Int63n(1000)

			patient.PatientName = _name
			patient.PatientDisease = _disease
			patient.PatientAge = _age

			addPatient(patient)

		}
	}

}
