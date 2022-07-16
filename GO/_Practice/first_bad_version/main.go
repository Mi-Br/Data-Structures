package main

//https://leetcode.com/problems/first-bad-version/
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l := 1

	if isBadVersion(l) {
		return l
	}

	for l < n {
		m := l + int((n-l+1)/2) // check with upper bound because if we have even elements (ie 2 ) we need to break the loop otherwise it will always be l < n

		if !isBadVersion(m) {
			l = m
		} else {
			if !isBadVersion(m - 1) {
				return m
			} else {
				n = m
			}

		}
	}
	return l
}

func main() {

}
