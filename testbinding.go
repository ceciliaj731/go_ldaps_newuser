package main

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

func main() {
	// LDAP server connection parameters
	// server := "ldap://ldap.example.com:389"
	// bindDN := "cn=admin,dc=example,dc=com"
	//  bindPassword := "admin_password"

	// User credentials
	userDN := "cn=rose,ou=testing,dc=testing,dc=example,dc=com"
	// userPassword := "F7p53n14b!"

	// Create an LDAP client connection
	l, err := ldap.Dial("tcp", "10.0.0.110:389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// Attempt to bind as the user
	err = l.Bind(userDN, "J8F91n24b!")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User successfully bound.")
}
