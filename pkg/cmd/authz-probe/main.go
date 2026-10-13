// Command authz-probe asks the authz service whether a user can read a dashboard.
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	user := flag.String("user", "", "user UID")
	dashboard := flag.String("dashboard", "", "dashboard UID")
	flag.Parse()
	if *user == "" || *dashboard == "" {
		fmt.Fprintln(os.Stderr, "usage: authz-probe -user <uid> -dashboard <uid>")
		os.Exit(2)
	}
	fmt.Printf("checking dashboards:uid:%s for user %s\n", *dashboard, *user)
}
