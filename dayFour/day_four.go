package dayFour

import (
	"github.com/jgsheppa/advent-of-code-2022/common"
	"strconv"
	"strings"
)

// Section contains the start and end
// of an elf's cleaning section
type Section struct {
	// start is the beginning of the elf's cleaning section
	start int
	// end is the end of the elf's cleaning section
	end int
}

// CalculateSectionsContainingOtherSections finds section pairs in which one
// section is contained in the second and returns the sum of such pairs
func CalculateSectionsContainingOtherSections(filename string) (int, error) {
	testData, err := common.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	overLappingSections := 0

	for _, pair := range testData {
		splitSections := strings.Split(pair, ",")

		sectionsAsNumbers, err := createMapOfSectionNumbers(splitSections)
		if err != nil {
			return 0, nil
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

func createMapOfSectionNumbers(stringSections []string) (map[int]Section, error) {
	sectionsAsNumbers := make(map[int]Section)

	for index, section := range stringSections {
		split := strings.Split(section, "-")
		start, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		finish, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}

		sectionsAsNumbers[index] = Section{
			start,
			finish,
		}
	}
	return sectionsAsNumbers, nil
}

// CalculateOverlappingSections finds section pairs in which any
// parts of the sections overlap and returns the sum of these pairs
func CalculateOverlappingSections(filename string) (int, error) {
	testData, err := common.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	overLappingSections := 0

	for _, pair := range testData {
		splitSections := strings.Split(pair, ",")

		sectionsAsNumbers, err := createMapOfSectionNumbers(splitSections)
		if err != nil {
			return 0, nil
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
