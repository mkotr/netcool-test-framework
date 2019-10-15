package service

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/tealeg/xlsx"
)

const (
	TestsSheet           = "TESTS"
	TrapsSheet           = "TRAPS"
	VarbindsSheet        = "TRAPVARBINDS"
	ExpectedResultsSheet = "EXPECTEDRESULTS"
)

var expectedHeaders = map[string][]string{
	TestsSheet:           {"TEST_NAME", "SLEEP_TIME", "NETWORK_DOMAIN", "EQUIPMENT_TYPE", "EQUIPMENT_ROLE", "PROBE_TYPE", "TEST_GROUP_NAME", "ACTION"},
	TrapsSheet:           {"TEST_NAME", "TRAP_NO", "TRAP_OID", "TRAP_NAME", "TRAP_DESC", "PROBE_DESC", "DEVICE_NAME", "TEST_TYPE"},
	VarbindsSheet:        {"TRAP_NO", "VARBIND_NAME", "VARBIND_VALUE"},
	ExpectedResultsSheet: {"TRAP_NO", "COL_NAME", "EXPECTED_VALUE"},
}

type TestSheet struct {
	TestName      string `xlsx:"0" json:"name"`
	SleepTime     int    `xlsx:"1" json:"sleepTime"`
	NetworkDomain string `xlsx:"2" json:"networkDomain"`
	EquipmentType string `xlsx:"3" json:"equipmentType"`
	EquipmentRole string `xlsx:"4" json:"equipmentRole"`
	ProbeType     string `xlsx:"5" json:"probeType"`
	TestGroupName string `xlsx:"6" json:"testGroupName"`
	Action        string `xlsx:"7" json:"action"`
}

type TrapSheet struct {
	TestName   string `xlsx:"0" json:"testName"`
	TrapNo     string `xlsx:"1" json:"trapNo"`
	TrapOID    string `xlsx:"2" json:"oid"`
	TrapName   string `xlsx:"3" json:"name"`
	TrapDesc   string `xlsx:"4" json:"trapDesc"`
	ProbeDesc  string `xlsx:"5" json:"probeDesc"`
	DeviceName string `xlsx:"6" json:"deviceName"`
	TestType   string `xlsx:"7" json:"testType"`
}

type VarbindSheet struct {
	TrapNo       string `xlsx:"0" json:"trapNo"`
	VarbindName  string `xlsx:"1" json:"name"`
	VarbindValue string `xlsx:"2" json:"value"`
}

type ExpectedResultSheet struct {
	TrapNo        string `xlsx:"0" json:"trapNo"`
	ColName       string `xlsx:"1" json:"name"`
	ExpectedValue string `xlsx:"2" json:"value"`
}

type ParsedFile struct {
	Tests           []TestSheet           `json:"tests"`
	Traps           []TrapSheet           `json:"traps"`
	Varbinds        []VarbindSheet        `json:"varbinds"`
	ExpectedResults []ExpectedResultSheet `json:"expectedResults"`
}

type FileParser struct {
	file multipart.File
}

func (f *FileParser) parseFile() (ParsedFile, error) {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, f.file); err != nil {
		return ParsedFile{}, err
	}

	xlFile, err := xlsx.OpenBinary(buf.Bytes())
	if err != nil {
		return ParsedFile{}, err
	}

	//We do a quick check here to see if the headers are what we expect.
	if valid := f.validateSheets(xlFile); !valid {
		return ParsedFile{}, err
	}

	parsedResult, err := f.parseExcel(xlFile)
	if err != nil {
		return ParsedFile{}, err
	}
	return parsedResult, nil
}

func (f *FileParser) validateSheets(xlFile *xlsx.File) bool {
	isValid := true
	for _, sheet := range xlFile.Sheets {
		headerRow := []string{}
		for _, cell := range sheet.Row(0).Cells {
			headerRow = append(headerRow, cell.String())
		}
		isValid = f.validateHeaders(headerRow, sheet.Name)
	}

	return isValid
}

func (f *FileParser) validateHeaders(actualHeaders []string, sheetName string) bool {
	expectedSheetHeaders := expectedHeaders[sheetName]

	if len(actualHeaders) != len(expectedSheetHeaders) {
		return false
	}

	for i, v := range actualHeaders {
		if v != expectedSheetHeaders[i] {
			return false
		}
	}
	return true
}

func (f *FileParser) parseExcel(xlFile *xlsx.File) (ParsedFile, error) {
	var tests []TestSheet
	var traps []TrapSheet
	var varbinds []VarbindSheet
	var expectedResults []ExpectedResultSheet

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows[1:] {
			switch sheet.Name {
			case TestsSheet:
				testInfo := TestSheet{}
				err := row.ReadStruct(&testInfo)
				if err != nil {
					return ParsedFile{}, err
				}
				tests = append(tests, testInfo)

			case TrapsSheet:
				trapInfo := TrapSheet{}
				err := row.ReadStruct(&trapInfo)
				if err != nil {
					return ParsedFile{}, err
				}
				traps = append(traps, trapInfo)

			case VarbindsSheet:
				varbindInfo := VarbindSheet{}
				err := row.ReadStruct(&varbindInfo)
				if err != nil {
					return ParsedFile{}, err
				}
				varbinds = append(varbinds, varbindInfo)

			case ExpectedResultsSheet:
				resultInfo := ExpectedResultSheet{}
				err := row.ReadStruct(&resultInfo)
				if err != nil {
					return ParsedFile{}, err
				}
				expectedResults = append(expectedResults, resultInfo)
			}
		}
	}
	return ParsedFile{tests, traps, varbinds, expectedResults}, nil
}
