package solution

import "testing"

func TestClimbingStairs(t *testing.T) {
	for _, test := range climbingStairsTestCases {
		actual := climbStairs(test.input)
		if actual != test.expected {
			t.Errorf("climbing stairs step %d, expected %d, actual %d", test.input, test.expected, actual)
		}
	}
}

func BenchmarkClimbingStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range climbingStairsTestCases {
			climbStairs(test.input)
		}
	}
}
