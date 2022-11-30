package stepic2_4

type TestStruct struct{
	On bool
	Ammo int
	Power int
}

func (t *TestStruct) Shoot() bool {
	if t.On == false || t.Ammo <= 0 {
		return false
	}
	t.Ammo--
	return true
}

func (t *TestStruct) RideBike() bool {
	if t.On == false || t.Power <= 0 {
		return false
	}
	t.Power--
	return true
}

/*func main() {
  var a = TestStruct {true, 2,3}
  testStruct := &a*/