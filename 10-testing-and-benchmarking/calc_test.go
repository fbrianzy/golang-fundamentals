package calc

import "testing"

func TestSum(t *testing.T){
    got := Sum(1,2,3)
    if got != 6 { t.Fatalf("want 6 got %d", got) }
}

func BenchmarkSum(b *testing.B){
    for i:=0;i<b.N;i++ { Sum(1,2,3,4,5,6,7,8,9) }
}
