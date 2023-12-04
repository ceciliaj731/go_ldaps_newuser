package main

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
	"golang.org/x/text/encoding/unicode"
)

func main() {
	// Connect to the LDAP server
	//      l, err := ldap.Dial("tcp", "10.0.0.110:389")
	l, err := ldap.DialTLS("tcp", "10.0.0.110:636", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	//	if err != nil {
	//		fmt.Println("Error starting TLS:", err)
	//		return
	//	}
	// Bind with admin credentials
	err = l.Bind("admin", "J8F91n24b!")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new user

	addReq := ldap.NewAddRequest("cn=rose,ou=testing,dc=testing,dc=example,dc=com", nil)
	addReq.Attribute("objectClass", []string{"top", "person", "organizationalPerson", "user"})
	addReq.Attribute("displayName", []string{"rose"})
	addReq.Attribute("sAMAccountName", []string{"rose"})

	err = l.Add(addReq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User created successfully")

	//encode psw using littleendian
	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, _ := utf16.NewEncoder().String("\"J8F91n24b!\"")

	//modreq1,change psw
	modReq := ldap.NewModifyRequest("cn=rose,ou=testing,dc=testing,dc=example,dc=com", nil)
	modReq.Add("unicodePwd", []string{pwdEncoded})

	if err := l.Modify(modReq); err != nil {
		log.Fatal("error setting user password:", modReq, err)
	}

	//modreq2,change user status to normal
	modReq2 := ldap.NewModifyRequest("cn=rose,ou=testing,dc=testing,dc=example,dc=com", nil)
	modReq2.Replace("userAccountControl", []string{"512"})
	if err := l.Modify(modReq2); err != nil {
		log.Fatal("error setting to active", modReq2, err)
	}
}
