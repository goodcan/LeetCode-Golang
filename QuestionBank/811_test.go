/*
 * @Time     : 2019/11/12 10:31
 * @Author   : cancan
 * @File     : 811_test.go
 * @Function :
 */

package QuestionBank

import (
	"sort"
	"testing"

	"LeetCode-Golang/utils"
)

func Test_811(t *testing.T) {
	tests := []struct {
		cpdomains []string
		ans       []string
	}{
		{
			[]string{"9001 discuss.leetcode.com"},
			[]string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"},
		},
		{
			[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
			[]string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		},
	}

	for _, test := range tests {
		ans := subdomainVisits(test.cpdomains)
		sort.Strings(ans)
		sort.Strings(test.ans)
		if !utils.SliceEqual(ans, test.ans) {
			t.Errorf("failure cpdomains %v ans %v", test.cpdomains, test.ans)
		}
	}
}
