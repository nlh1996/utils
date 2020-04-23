package utils

// ValidateCode 验证码字典 
type ValidateCode struct{
	code map[string]uint16
}

var validateCode *ValidateCode

func init() {
	validateCode = &ValidateCode{
		code: make(map[string]uint16),
	}
}

// GetValidateCode .
func GetValidateCode() *ValidateCode {
	if len(validateCode.code) == 0 {
		validateCode.code["a"] = 0
		validateCode.code["b"] = 1
		validateCode.code["c"] = 2
		validateCode.code["d"] = 3
		validateCode.code["e"] = 4
		validateCode.code["f"] = 5
		validateCode.code["g"] = 6
		validateCode.code["h"] = 7
		validateCode.code["i"] = 8
		validateCode.code["j"] = 9
	}
	return validateCode
}

// RandCode .
func (v *ValidateCode) RandCode(len int) []string {
	var (
		arr []string
		index int
	)
	for k := range v.code {
		arr = append(arr, k)
		index ++
		if index != len {		
			arr = append(arr, "add")
		}
		if index == len {
			break
		}
	}
	return arr 
}

// Compute 结果
func (v *ValidateCode) Compute(arr []string) uint16 {
	var sum uint16
	for _,value := range arr {
		if value != "add" {
			sum += v.code[value]
		}
	}
	return sum 
}