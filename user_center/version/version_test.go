// @Author: Ciusyan 2023/1/25
package version_test

import (
	"Go-To-Byte/DouSheng/user_center/version"
	"fmt"
	"testing"
)

func TestFullVersion(t *testing.T) {
	fmt.Println(version.FullVersion())
}