package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/madimaa/aoc2020/lib"
)

type passportStruct struct {
	byr, iyr, eyr, cid int
	hgt, pid, hcl, ecl string
}

func main() {
	lib.Start()
	fmt.Println("Part 1")

	input := lib.OpenFile("04.txt")
	passports := make([]*passportStruct, 0)
	passport := &passportStruct{}
	validPassports := 0
	for _, row := range input {
		if len(row) == 0 {
			passports = append(passports, passport)
			if validate(passport) {
				validPassports++
			}
			passport = &passportStruct{}
		}

		content := strings.Split(row, " ")
		parsePassport(passport, content)
	}

	passports = append(passports, passport)
	if validate(passport) {
		validPassports++
	}

	fmt.Println("Result: ", validPassports)
	lib.Elapsed()

	lib.Start()
	fmt.Println("Part 2")

	validPassports = 0
	for _, passport := range passports {
		if stricterValidate(passport) {
			validPassports++
		}
	}

	fmt.Println("Result: ", validPassports)

	lib.Elapsed()
	os.Exit(0)
}

func parsePassport(passport *passportStruct, fields []string) {
	for _, field := range fields {
		content := strings.Split(field, ":")
		switch content[0] {
		case "byr":
			val, _ := strconv.Atoi(content[1])
			passport.byr = val
		case "iyr":
			val, _ := strconv.Atoi(content[1])
			passport.iyr = val
		case "eyr":
			val, _ := strconv.Atoi(content[1])
			passport.eyr = val
		case "hgt":
			passport.hgt = content[1]
		case "pid":
			passport.pid = content[1]
		case "cid":
			val, _ := strconv.Atoi(content[1])
			passport.cid = val
		case "hcl":
			passport.hcl = content[1]
		case "ecl":
			passport.ecl = content[1]
		}
	}
}

func validate(passport *passportStruct) bool {
	return passport.byr != 0 && passport.iyr != 0 && passport.eyr != 0 && passport.hgt != "" && passport.pid != "" && passport.hcl != "" && passport.ecl != ""
}

func stricterValidate(passport *passportStruct) bool {
	byr, iyr, eyr, hgt, pid, hcl, ecl := true, true, true, true, true, true, false

	if passport.byr < 1920 || passport.byr > 2002 {
		byr = false
	}

	if passport.iyr < 2010 || passport.iyr > 2020 {
		iyr = false
	}

	if passport.eyr < 2020 || passport.eyr > 2030 {
		eyr = false
	}

	if strings.Contains(passport.hgt, "cm") {
		heightStr := strings.TrimSuffix(passport.hgt, "cm")
		height, err := strconv.Atoi(heightStr)
		if err != nil {
			hgt = false
		} else if height < 150 || height > 193 {
			hgt = false
		}
	} else if strings.Contains(passport.hgt, "in") {
		heightStr := strings.TrimSuffix(passport.hgt, "in")
		height, err := strconv.Atoi(heightStr)
		if err != nil {
			hgt = false
		} else if height < 59 || height > 76 {
			hgt = false
		}
	} else {
		hgt = false
	}

	if match, _ := regexp.MatchString("^#[0-9a-f]{6}", passport.hcl); !match {
		hcl = false
	}

	switch passport.ecl {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		ecl = true
	}

	if match, _ := regexp.MatchString("^[0-9]{9}$", passport.pid); !match {
		pid = false
	}

	return byr && iyr && eyr && hgt && pid && hcl && ecl
}
