package dateparse

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

/*
1.январь
2.февраль
3.март
4.апрель
5.май
6.июнь
7.июль.
8.август
9.сентябрь
10.октябрь
11.ноябрь
12.декабрь
*/
/*
январь – january
февраль – february
март – march
апрель – april
май – may
июнь – june [ dʒu:n ]
июль – july [ dʒu’lai ]
август – august [ ˈɔːɡəst ]
сентябрь – september [ sep’tembə ]
октябрь – october [ ɔk’təubə ]
ноябрь – november [ nəu’vembə ]
декабрь – december [ di’sembə ]
*/
func ParseAny(d string) (time.Time, error) {
	date := strings.TrimSpace(strings.ToLower(d))
	typeDate := Broken
	//dd-mounth-yy    27-September-20
	SimpleDashAmericanYYR := regexp.MustCompile(`\b\d{2}\-[A-z]{3,9}\-\d\d\b`)
	if SimpleDashAmericanYYR.MatchString(date) {
		//fmt.Printf("[%s] is SimpleDashAmericanYY\n", date)
		typeDate = SimpleDashAmericanYY
	}
	//dd-mounth-yyyy    08-APR-2020
	SimpleDashAmericanYYYYR := regexp.MustCompile(`\b\d{2}\-[A-z]{3,9}\-\d\d\d\d\b`)
	if SimpleDashAmericanYYYYR.MatchString(date) {
		//	fmt.Printf("[%s] is SimpleDashAmericanYYYY\n", date)
		typeDate = SimpleDashAmericanYYYY
	}
	//dd.month.yy 09.Apr.20
	SimpleDotAmericanYYR := regexp.MustCompile(`\b\d{2}\.[A-z]{3,9}\.\d\d\b`)
	if SimpleDotAmericanYYR.MatchString(date) {
		//		fmt.Printf("[%s] is SimpleDotAmericanYYR\n", date)
		typeDate = SimpleDotAmericanYY
	}
	//dd.month.yyyy 09.Apr.2020
	SimpleDotAmericanYYYYR := regexp.MustCompile(`\b\d{2}\.[A-z]{3,9}\.\d\d\d\d\b`)
	if SimpleDotAmericanYYYYR.MatchString(date) {
		//	fmt.Printf("[%s] is SimpleDotAmericanYYYYR\n", date)
		typeDate = SimpleDotAmericanYYYY
	}
	//dd.month.yy 03.апр.20
	NumerateDotRussianYYR := regexp.MustCompile(`\b\d{2}\.[A-z]{3,8}\.\d\d\b`)
	if NumerateDotRussianYYR.MatchString(date) {
		//	fmt.Printf("[%s] is NumerateDotRussianYYR\n", date)
		typeDate = NumerateDotRussianYY
	}
	//dd.month.yyyy 03.апр.2020
	NumerateDotRussianYYYYR := regexp.MustCompile(`\b\d{2}\.[а-я]{3,8}\.\d\d\d\d\b`)
	if NumerateDotRussianYYYYR.MatchString(date) {
		//	fmt.Printf("[%s] is NumerateDotRussianYYYYR\n", date)
		typeDate = NumerateDotRussianYYYY
	}
	//dd.mm.yy 30.11.20
	SimpleDotYYR := regexp.MustCompile(`\b\d{2}\.\d\d\.\d\d\b`)
	if SimpleDotYYR.MatchString(date) {
		//	fmt.Printf("[%s] is SimpleDotYYR\n", date)
		typeDate = SimpleDotYY
	}
	//dd.mm.yyyy 30.11.2020
	SimpleDotYYYYR := regexp.MustCompile(`\b\d{2}\.\d\d\.\d\d\d\d\b`)
	if SimpleDotYYYYR.MatchString(date) {
		//	fmt.Printf("[%s] is SimpleDotYYYYR\n", date)
		typeDate = SimpleDotYYYY
	}
	//dd/mm/yy 30/11/20
	SimpleSlashYYR := regexp.MustCompile(`\b\d{2}\/\d\d\/\d\d\b`)
	if SimpleSlashYYR.MatchString(date) {
		//	fmt.Printf("[%s] is SimpleSlashYYR\n", date)
		typeDate = SimpleSlashYY
	}
	//dd/mm/yyyy 30/11/2020
	SimpleSlashYYYYR := regexp.MustCompile(`\b\d{2}\/\d\d\/\d\d\d\d\b`)
	if SimpleSlashYYYYR.MatchString(date) {
		//	fmt.Printf("[%s] is SimpleSlashYYYYR\n", date)
		typeDate = SimpleSlashYYYY
	}
	//dd month yy 27 May 20
	SpaceAmericanYYR := regexp.MustCompile(`\b\d{2}\s[A-z]{3,9}\s\d{2}\b`)
	if SpaceAmericanYYR.MatchString(date) {
		//		fmt.Printf("[%s] is SpaceAmericanYYR\n", date)
		typeDate = SpaceAmericanYY
	}
	//dd month yyyy 27 May 2020
	SpaceAmericanYYYYR := regexp.MustCompile(`\b\d{2}\s[A-z]{3,9}\s\d{4}\b`)
	if SpaceAmericanYYYYR.MatchString(date) {
		//	fmt.Printf("[%s] is SpaceAmericanYYYYR\n", date)
		typeDate = SpaceAmericanYYYY
	}
	//dd/month/yy 06/Apr/20
	SlashWordMounthYYR := regexp.MustCompile(`\b\d{2}\/[A-z]{3,9}\/\d\d\b`)
	if SlashWordMounthYYR.MatchString(date) {
		//	fmt.Printf("[%s] is NumerateDotRussianYYR\n", date)
		typeDate = SlashWordMounthYY
	}
	//dd/month/yyyy 06/Apr/2020
	SlashWordMounthYYYYR := regexp.MustCompile(`\b\d{2}\/[A-z]{3,9}\/\d\d\d\d\b`)
	if SlashWordMounthYYYYR.MatchString(date) {
		//	fmt.Printf("[%s] is NumerateDotRussianYYYYR\n", date)
		typeDate = SlashWordMounthYYYY
	}
	//dd/month/yy 06/апр/20
	SlashWordMounthRussianYYR := regexp.MustCompile(`\b\d{2}\/[А-я]{3,8}\/\d\d\b`)
	if SlashWordMounthRussianYYR.MatchString(date) {
		//	fmt.Printf("[%s] is NumerateDotRussianYYR\n", date)
		typeDate = SlashWordMounthRussianYY
	}
	//dd/month/yyyy 06/апр/2020
	SlashWordMounthRussianYYYYR := regexp.MustCompile(`\b\d{2}\/[А-я]{3,8}\/\d\d\d\d\b`)
	if SlashWordMounthRussianYYYYR.MatchString(date) {
		//	fmt.Printf("[%s] is NumerateDotRussianYYYYR\n", date)
		typeDate = SlashWordMounthRussianYYYY
	}
	switch typeDate {
	case SimpleDotYY, SimpleDotAmericanYY, NumerateDotRussianYY, SimpleDotYYYY, SimpleDotAmericanYYYY, NumerateDotRussianYYYY:
		pD := strings.Split(date, ".")
		m, err := GetMonth(pD[1])
		if err != nil {
			return time.Now(), err
		}
		y := pD[2]
		if typeDate == SimpleDotYY || typeDate == SimpleDotAmericanYY || typeDate == NumerateDotRussianYY {
			y = "20" + y
		}
		dC := fmt.Sprintf("%s-%s-%s", y, m, pD[0])
		dt, err := time.Parse(layoutISO, dC)
		if err != nil {
			return time.Now(), err
		}
		return dt, nil
	case SpaceAmericanYY, SpaceAmericanYYYY:
		pD := strings.Split(date, " ")
		m, err := GetMonth(pD[1])
		if err != nil {
			return time.Now(), err
		}
		y := pD[2]
		if typeDate == SpaceAmericanYY {
			y = "20" + y
		}
		dC := fmt.Sprintf("%s-%s-%s", y, m, pD[0])
		dt, err := time.Parse(layoutISO, dC)
		if err != nil {
			return time.Now(), err
		}
		return dt, nil
	case SimpleSlashYY, SimpleSlashYYYY:
		pD := strings.Split(date, "/")
		m, err := GetMonth(pD[0])
		if err != nil {
			return time.Now(), err
		}
		y := pD[2]
		if typeDate == SimpleSlashYY {
			y = "20" + y
		}
		dC := fmt.Sprintf("%s-%s-%s", y, m, pD[1])
		dt, err := time.Parse(layoutISO, dC)
		if err != nil {
			return time.Now(), err
		}
		return dt, nil
	case SimpleDashAmericanYY, SimpleDashAmericanYYYY:
		pD := strings.Split(date, "-")
		m, err := GetMonth(pD[1])
		if err != nil {
			return time.Now(), err
		}
		y := pD[2]
		if typeDate == SimpleDashAmericanYY {
			y = "20" + y
		}
		dC := fmt.Sprintf("%s-%s-%s", y, m, pD[0])
		dt, err := time.Parse(layoutISO, dC)
		if err != nil {
			return time.Now(), err
		}
		return dt, nil
	case SlashWordMounthYY, SlashWordMounthRussianYY, SlashWordMounthYYYY, SlashWordMounthRussianYYYY:
		pD := strings.Split(date, "/")
		m, err := GetMonth(pD[1])
		if err != nil {
			return time.Now(), err
		}
		y := pD[2]
		if typeDate == SlashWordMounthYY || typeDate == SlashWordMounthRussianYY {
			y = "20" + y
		}
		dC := fmt.Sprintf("%s-%s-%s", y, m, pD[0])
		dt, err := time.Parse(layoutISO, dC)
		if err != nil {
			return time.Now(), err
		}
		return dt, nil
	case Broken:
		return time.Now(), fmt.Errorf("broken")
	default:
		return time.Now(), fmt.Errorf("empty")
	}

}
func GetMonth(d string) (string, error) {
	date := strings.TrimSpace(strings.ToLower(d))
	switch date {
	case "january", "jan", "январь", "янв", "01":
		return "01", nil
	case "february", "feb", "февраль", "фев", "02":

		return "02", nil
	case "march", "mar", "март", "мар", "03":
		return "03", nil
	case "april", "apr", "апрель", "апр", "04":
		return "04", nil
	case "may", "май", "05":
		return "05", nil
	case "june", "jun", "июнь", "июн", "06":
		return "06", nil
	case "july", "jul", "июль", "июл", "07":
		return "07", nil
	case "august", "aug", "август", "авг", "08":
		return "08", nil
	case "september", "sep", "сентябрь", "сен", "09":
		return "09", nil
	case "october", "oct", "октябрь", "окт", "10":
		return "10", nil
	case "november", "nov", "ноябрь", "ноя", "11":
		return "11", nil
	case "december", "dec", "декабрь", "дек", "12":

		return "12", nil
	default:
		return "", fmt.Errorf("mounth not found")
	}
	//return "", fmt.Errorf("mounth not found")
}
