package entity
import "testing"
func TestUser_Validate(t *testing.T) {
	u := &User{
		FirstName: "aa",
		LastName:  "bb",
		Email:     "cc@adf.com",
		Password:  "dd2ds1",
	}

	err := u.Validate()
	if err != nil{
		t.Log(err.Error())
	}else{
		t.Log("success")
	}
}
