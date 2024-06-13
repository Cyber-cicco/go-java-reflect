package utils

import "strings"


func  GetFileNameFromUrl(u string) string {
    a := strings.Split(u, "/")
    return a[len(a)-1]
}
