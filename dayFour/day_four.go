package dayFour

import (
	"github.com/jgsheppa/advent-of-code-2022/common"
	"strconv"
	"strings"
)

type Section struct {
	start int
	end   int
}

func CalculateSectionsContainingOtherSections(filename string) (int, error) {
	testData, err := common.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	overLappingSections := 0

	for _, pair := range testData {
		splitSections := strings.Split(pair, ",")

		sectionsAsNumbers := make(map[int]Section)

		for index, section := range splitSections {
			split := strings.Split(section, "-")
			start, err := strconv.Atoi(split[0])
			if err != nil {
				return 0, err
			}
			finish, err := strconv.Atoi(split[1])
			if err != nil {
				return 0, err
			}

			sectionsAsNumbers[index] = Section{
				start,
				finish,
			}

		}

		switch true {
		case sectionsAsNumbers[0].start >= sectionsAsNumbers[1].start && sectionsAsNumbers[0].end <= sectionsAsNumbers[1].end:
			overLappingSections++
			break
		case sectionsAsNumbers[1].start >= sectionsAsNumbers[0].start && sectionsAsNumbers[1].end <= sectionsAsNumbers[0].end:
			overLappingSections++
			break
		}
	}
	return overLappingSections, nil
}

func CalculateOverlappingSections(filename string) (int, error) {
	testData, err := common.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	overLappingSections := 0

	for _, pair := range testData {
		splitSections := strings.Split(pair, ",")

		sectionsAsNumbers := make(map[int]Section)

		for index, section := range splitSections {
			split := strings.Split(section, "-")
			start, err := strconv.Atoi(split[0])
			if err != nil {
				return 0, err
			}
			finish, err := strconv.Atoi(split[1])
			if err != nil {
				return 0, err
			}

			sectionsAsNumbers[index] = Section{
				start,
				finish,
			}

		}

		switch true {
		case sectionsAsNumbers[0].start >= sectionsAsNumbers[1].start && sectionsAsNumbers[0].start <= sectionsAsNumbers[1].end:
			overLappingSections++
			break
		case sectionsAsNumbers[0].end <= sectionsAsNumbers[1].end && sectionsAsNumbers[0].end >= sectionsAsNumbers[1].start:
			overLappingSections++
			break
		case sectionsAsNumbers[0].start >= sectionsAsNumbers[1].start && sectionsAsNumbers[0].end <= sectionsAsNumbers[1].end:
			overLappingSections++
			break
		case sectionsAsNumbers[1].start >= sectionsAsNumbers[0].start && sectionsAsNumbers[1].end <= sectionsAsNumbers[0].end:
			overLappingSections++
			break
		}
	}
	return overLappingSections, nil
}
