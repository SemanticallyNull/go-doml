
//line doml.rl:1
package doml

import "fmt"


//line doml.go:9
const doml_start int = 1
const doml_first_final int = 46
const doml_error int = 0

const doml_en_main int = 1


//line doml.rl:8


func Parse(data string) (map[string]any, error) {
  cs, p, pe := 0, 0, len(data)
  val := make(map[string]any, 0)
  var key, current string
  currentArr := make([]string, 0)

  
//line doml.go:27
	{
	cs = doml_start
	}

//line doml.go:32
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	}
	goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 32:
			goto st1
		case 240:
			goto st2
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st1
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 159 {
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 143 {
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 129 {
			goto st5
		}
		goto st0
tr28:
//line doml.rl:26
 val[key] = currentArr; currentArr = make([]string, 0) 
	goto st5
tr50:
//line doml.rl:24
 val[key] = current; current = "" 
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line doml.go:193
		switch data[p] {
		case 32:
			goto st5
		case 226:
			goto st6
		case 240:
			goto st43
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st5
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 156 {
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 140 {
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 239 {
			goto st9
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 184 {
			goto st10
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 143 {
			goto st11
		}
		goto st0
tr13:
//line doml.rl:22
 current = current + string(data[p]) 
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line doml.go:260
		switch data[p] {
		case 32:
			goto st0
		case 128:
			goto st0
		case 140:
			goto st0
		case 143:
			goto st0
		case 145:
			goto st0
		case 159:
			goto st0
		case 182:
			goto st0
		case 184:
			goto st0
		case 226:
			goto st12
		}
		switch {
		case data[p] < 150:
			if 9 <= data[p] && data[p] <= 10 {
				goto st0
			}
		case data[p] > 151:
			switch {
			case data[p] > 156:
				if 239 <= data[p] && data[p] <= 240 {
					goto st0
				}
			case data[p] >= 155:
				goto st0
			}
		default:
			goto st0
		}
		goto tr13
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 156 {
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 140 {
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 239 {
			goto st15
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 184 {
			goto st16
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if data[p] == 143 {
			goto tr19
		}
		goto st0
tr19:
//line doml.rl:23
 key = current; current = "" 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line doml.go:353
		switch data[p] {
		case 32:
			goto st17
		case 226:
			goto st18
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st17
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 143:
			goto st19
		case 150:
			goto st33
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 169 {
			goto st20
		}
		goto st0
tr39:
//line doml.rl:25
 currentArr = append(currentArr, current); current = "" 
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line doml.go:394
		switch data[p] {
		case 32:
			goto st20
		case 226:
			goto st21
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st20
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 143:
			goto st22
		case 150:
			goto st23
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 170 {
			goto tr28
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 182 {
			goto st24
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 239 {
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 184 {
			goto st26
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 143 {
			goto st27
		}
		goto st0
tr33:
//line doml.rl:22
 current = current + string(data[p]) 
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line doml.go:471
		switch data[p] {
		case 32:
			goto st0
		case 128:
			goto st0
		case 140:
			goto st0
		case 143:
			goto st0
		case 145:
			goto st0
		case 159:
			goto st0
		case 182:
			goto st0
		case 184:
			goto st0
		case 226:
			goto st28
		}
		switch {
		case data[p] < 150:
			if 9 <= data[p] && data[p] <= 10 {
				goto st0
			}
		case data[p] > 151:
			switch {
			case data[p] > 156:
				if 239 <= data[p] && data[p] <= 240 {
					goto st0
				}
			case data[p] >= 155:
				goto st0
			}
		default:
			goto st0
		}
		goto tr33
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if data[p] == 151 {
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 128 {
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if data[p] == 239 {
			goto st31
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if data[p] == 184 {
			goto st32
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 143 {
			goto tr39
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 182 {
			goto st34
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 239 {
			goto st35
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if data[p] == 184 {
			goto st36
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if data[p] == 143 {
			goto st37
		}
		goto st0
tr44:
//line doml.rl:22
 current = current + string(data[p]) 
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line doml.go:600
		switch data[p] {
		case 32:
			goto st0
		case 128:
			goto st0
		case 140:
			goto st0
		case 143:
			goto st0
		case 145:
			goto st0
		case 159:
			goto st0
		case 182:
			goto st0
		case 184:
			goto st0
		case 226:
			goto st38
		}
		switch {
		case data[p] < 150:
			if 9 <= data[p] && data[p] <= 10 {
				goto st0
			}
		case data[p] > 151:
			switch {
			case data[p] > 156:
				if 239 <= data[p] && data[p] <= 240 {
					goto st0
				}
			case data[p] >= 155:
				goto st0
			}
		default:
			goto st0
		}
		goto tr44
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if data[p] == 151 {
			goto st39
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if data[p] == 128 {
			goto st40
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		if data[p] == 239 {
			goto st41
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if data[p] == 184 {
			goto st42
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 143 {
			goto tr50
		}
		goto st0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		if data[p] == 159 {
			goto st44
		}
		goto st0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		if data[p] == 155 {
			goto st45
		}
		goto st0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		if data[p] == 145 {
			goto st46
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		if data[p] == 32 {
			goto st46
		}
		if 9 <= data[p] && data[p] <= 10 {
			goto st46
		}
		goto st0
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof

	_test_eof: {}
	_out: {}
	}

//line doml.rl:36


  if cs < doml_first_final {
    return map[string]any{}, fmt.Errorf("parse error: %s", data)
  }

  return val, nil
}
