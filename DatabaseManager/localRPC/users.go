package localRPC

import (
	"../params"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/syndtr/goleveldb/leveldb/util"
	"golang.org/x/crypto/scrypt"
)

type GetUsernamesReply struct {
	Usernames []string
	Passwords []string
}

// GetUsernames queries the sql databse for the list of local users
func (client *MedRecLocal) GetUsernames(r *http.Request, args *params.NoArgs, reply *GetUsernamesReply) error {
	tab := instantiateLookupTable()
	defer tab.Close()

	rows := tab.NewIterator(util.BytesPrefix([]byte("username-")), nil)
	rows.Release()

	reply.Usernames = []string{}
	for rows.Next() {
		reply.Usernames = append(reply.Usernames, string(rows.Value()))
	}

	return nil
}

type UserDetailsArgs struct {
	Username string
}

type GetUserDetailsReply struct {
	FirstName string
	LastName  string
}

// GetUserDetails queries the sql databse for the list of local users
func (client *MedRecLocal) GetUserDetails(r *http.Request, args *UserDetailsArgs, reply *GetUserDetailsReply) error {
	tab := instantiateLookupTable()
	defer tab.Close()

	firstName, _ := tab.Get([]byte(args.Username+"-firstName"), nil)
	lastName, _ := tab.Get([]byte(args.Username+"-lastName"), nil)

	reply.FirstName = string(firstName)
	reply.LastName = string(lastName)

	return nil
}

type NewUserArgs struct {
	FirstName string
	LastName  string
	Username  string
	Password  string
	Seed      string
}

type NewUserReply struct {
	Error string
}

// NewUser adds a new user to the database
func (client *MedRecLocal) NewUser(r *http.Request, args *NewUserArgs, reply *NewUserReply) error {
	tab := instantiateLookupTable()
	defer tab.Close()

	salt := make([]byte, ScryptSaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	}
	//use the salt to encrypt the password
	hashedPassword, err := scrypt.Key([]byte(args.Password), salt, ScryptN, ScryptR, ScryptP, ScryptKeyLen)
	if err != nil {
		log.Fatal(err)
	}
	//encode it for storage
	derivedPassword := base64.StdEncoding.EncodeToString(hashedPassword)

	//use the encrypted password to reversibly encrypt the privatekey
	//fist get the first aes encryption block
	block, err := aes.NewCipher([]byte(derivedPassword[0:AesKeyLen]))
	if err != nil {
		log.Fatal(err)
	}

	//generate a placeholder byte array for the future ciphertext
	ciphertext := make([]byte, aes.BlockSize+len(args.Seed))
	//we need a pseudorandom iv to ensure the first block of aes is secure
	//https://stackoverflow.com/questions/9049789/aes-encryption-key-versus-iv
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	//encrypt each block of the aes key in CFB mode
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(args.Seed))
	//encode the ciphertext for storage
	derivedSeed := base64.StdEncoding.EncodeToString(ciphertext)

	//we can't directly store the encrypted password since we used it to encrypt the private key
	//so encrypt it again before storage
	hashedPassword, err = scrypt.Key(hashedPassword, salt, ScryptN, ScryptR, ScryptP, ScryptKeyLen)
	if err != nil {
		log.Fatal(err)
	}
	derivedPassword = base64.StdEncoding.EncodeToString(hashedPassword)
	derivedPassword += ":" + base64.StdEncoding.EncodeToString(salt)

	tab.Put([]byte("username-"+args.Username), []byte(args.Username), nil)
	tab.Put([]byte(args.Username+"-firstName"), []byte(args.FirstName), nil)
	tab.Put([]byte(args.Username+"-lastName"), []byte(args.LastName), nil)
	tab.Put([]byte(args.Username+"-password"), []byte(derivedPassword), nil)
	tab.Put([]byte(args.Username+"-privateKey"), []byte(derivedSeed), nil)

	if err != nil {
		reply.Error = err.Error()
	}
	return nil
}

type GetSeedArgs struct {
	Username string
	Password string
}

type GetSeedReply struct {
	Seed  string
	Error string
}

// GetSeed ecrypts and retrives a user's private key seed
func (client *MedRecLocal) GetSeed(r *http.Request, args *GetSeedArgs, reply *GetSeedReply) error {
	tab := instantiateLookupTable()
	defer tab.Close()

	storedPassword, _ := tab.Get([]byte(args.Username+"-password"), nil)
	derivedSeed, _ := tab.Get([]byte(args.Username+"-privateKey"), nil)

	if string(storedPassword) != "" {
		saltedPass := strings.Split(string(storedPassword), ":")

		salt, _ := base64.StdEncoding.DecodeString(saltedPass[1])

		hashedPassword, err := scrypt.Key([]byte(args.Password), salt, ScryptN, ScryptR, ScryptP, ScryptKeyLen)
		if err != nil {
			log.Fatal(err)
		}
		hashedPassword2, err := scrypt.Key(hashedPassword, salt, ScryptN, ScryptR, ScryptP, ScryptKeyLen)
		if err != nil {
			log.Fatal(err)
		}
		derivedPassword := base64.StdEncoding.EncodeToString(hashedPassword2)

		if saltedPass[0] != derivedPassword {
			reply.Error = "Invalid password"
			return nil
		}

		// generate the first block of aes
		derivedPassword = base64.StdEncoding.EncodeToString(hashedPassword)
		block, err := aes.NewCipher([]byte(derivedPassword[0:AesKeyLen]))
		if err != nil {
			log.Fatal(err)
		}

		//decode the encrypted key from storage
		hashedSeed, _ := base64.StdEncoding.DecodeString(string(derivedSeed))
		//split up the ciphertext and iv
		iv := hashedSeed[:aes.BlockSize]
		ciphertext := hashedSeed[aes.BlockSize:]

		//run the decryption and set the privatekey in the reply
		stream := cipher.NewCFBDecrypter(block, iv)
		stream.XORKeyStream(ciphertext, ciphertext)
		reply.Seed = string(ciphertext)

	} else {
		reply.Error = "the selected username could not be found"
	}

	return nil
}

type DeleteUserArgs struct {
	Username string
	Password string
}

type DeleteUserReply struct {
	Error string
}

// DeleteUser retrieves a user's decrypts and retrives a user's private key
func (client *MedRecLocal) DeleteUser(r *http.Request, args *DeleteUserArgs, reply *DeleteUserReply) error {
	tab := instantiateLookupTable()
	defer tab.Close()

	storedPassword, _ := tab.Get([]byte(args.Username+"-password"), nil)
	saltedPass := strings.Split(string(storedPassword), ":")

	salt, _ := base64.StdEncoding.DecodeString(saltedPass[1])

	hashedPassword, err := scrypt.Key([]byte(args.Password), salt, ScryptN, ScryptR, ScryptP, ScryptKeyLen)
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword2, err := scrypt.Key(hashedPassword, salt, ScryptN, ScryptR, ScryptP, ScryptKeyLen)
	if err != nil {
		log.Fatal(err)
	}
	derivedPassword := base64.StdEncoding.EncodeToString(hashedPassword2)

	if saltedPass[0] != derivedPassword {
		reply.Error = "Password mismatch"
		return nil
	}

	tab.Delete([]byte("username-"+args.Username), nil)
	tab.Delete([]byte(args.Username+"-firstName"), nil)
	tab.Delete([]byte(args.Username+"-lastName"), nil)
	tab.Delete([]byte(args.Username+"-password"), nil)
	tab.Delete([]byte(args.Username+"-privateKey"), nil)

	return nil
}
