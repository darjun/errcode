package errcode

import "testing"

func TestCode(t *testing.T) {
	cases := []struct {
		errCode ErrCode
		expect  string
	}{
		{ERR_CODE_OK, "OK"},
		{ERR_CODE_INVALID_PARAMS, "无效参数"},
		{ERR_CODE_TIMEOUT, "超时"},
	}

	for _, testCase := range cases {
		if testCase.errCode.String() != testCase.expect {
			t.Errorf("error code %d description inconsistant actual:%s expect:%s", int(testCase.errCode), testCase.errCode, testCase.expect)
		}
	}
}
