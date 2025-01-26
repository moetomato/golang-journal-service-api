package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"
)

const idToken = `xxxxxxxxxx`

func main() {
	// construct public key with vals gotten via https://www.googleapis.com/oauth2/v3/certs
	N := "4gQqqklPFAI4AKTr0HPsxjHsJ3mAaPejrJ_aplDZsYUyH3bvEZ0vddQ7VYRy-Hozt-4lNjaw-T3fosSATtSGrQ2UtAkrxsS3_oeOgHyQ1Xt-OH3Pzgq1HZVMXf_xxCxOzhBffnCehI5eXZ2GxLn_1Xz-FNw2SJqNGudrxD4HodkhGsHvhbelvfE9-tozoFxlT7rIK8fWpR4SpZwQjbMhHYKjSAbuVjbZoF7wL0cqWYo3zT9OHp8XbfLqduabPgYN1CVuNYMomWIHdQO3SKdNXdgLbOqhkQ5xAbEo75C2zYcBHWfPuiVZclpClVPR7rN_sJPz7s6MWGQvMw3FpqcQyw"
	E := "AQAB"
	dn, _ := base64.RawURLEncoding.DecodeString(N)
	de, _ := base64.RawURLEncoding.DecodeString(E)

	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(dn),
		E: int(new(big.Int).SetBytes(de).Int64()),
	}

	// deconstruct id token
	arr := strings.Split(idToken, ".")
	header, payload, sig := arr[0], arr[1], arr[2]

	// header+payload -> sha256 encode(signed) -> message digest
	md := sha256.Sum256([]byte(header + "." + payload))

	// sig -> base64 decode -> public key decode -> message digest
	dsig, err := base64.RawURLEncoding.DecodeString(sig)
	if err != nil {
		fmt.Println("error sig:", err)
		return
	}

	if err := rsa.VerifyPKCS1v15(pk, crypto.SHA256, md[:], dsig); err != nil {
		fmt.Println("invalid token")
	} else {
		fmt.Println("valid token")
		headerData, err := base64.RawURLEncoding.DecodeString(header)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		payloadData, err := base64.RawURLEncoding.DecodeString(payload)
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println("header: ", string(headerData))
		fmt.Println("payload: ", string(payloadData))
	}
}

// func main() {
// 	var (
// 		dbUser = os.Getenv("DB_USER")
// 		dbPass = os.Getenv("DB_PASS")
// 		dbName = os.Getenv("DB_NAME")
// 		// [temporal] work on localhost とりあえず
// 		dbConn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPass, dbName)
// 	)

// 	db, err := sql.Open("mysql", dbConn)
// 	if err != nil {
// 		log.Println("failed to connect DataBase")
// 		return
// 	}
// 	r := api.NewRouter(db)

// 	log.Println("journal api server starting at port 8080...")
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }
