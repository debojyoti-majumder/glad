# glad

This a go module which would have multiple client and multiple server ( which acts as a cluster ).
Able to run the application using command

`go run main.go`

To run the application is server mode
`./glad server --port 3001`

In client mode
`./glad client --port 3001`

LDAP Client code can be made using the following command

`ldapsearch -D "cn=Directory Manager" -w "pass" -p 389 -h server.example.com -b "dc=example,dc=com" -s sub -x "(objectclass=*)"`