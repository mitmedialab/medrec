package localRPC

import (
	"testing"
)

func TestEncryption(t *testing.T) {
	client := new(MedRecLocal)

	newUserArgs := &NewUserArgs{
		Username: "testuser1",
		Password: "pass1",
		Seed:     "3a1076bf45ab87712ad64ccb3b10217737f7faacbf2872e88fdd9a537d8fe266",
	}

	newUserReply := &NewUserReply{}
	client.NewUser(nil, newUserArgs, newUserReply)

	getSeedArgs := &GetSeedArgs{
		Username: "testuser1",
		Password: "pass1",
	}

	getSeedReply := &GetSeedReply{}
	client.GetSeed(nil, getSeedArgs, getSeedReply)

	if newUserArgs.Seed != getSeedReply.Seed {
		t.Error("the saved privatekey could not be recovered")
	}

	deleteUserArgs := &DeleteUserArgs{
		Username: "testuser1",
		Password: "pass1",
	}

	deleteUserReply := &DeleteUserReply{}
	client.DeleteUser(nil, deleteUserArgs, deleteUserReply)

	getSeedReply.Seed = ""
	client.GetSeed(nil, getSeedArgs, getSeedReply)
	if newUserArgs.Seed == getSeedReply.Seed {
		t.Error("the user's private seed was not deleted")
	}
}
