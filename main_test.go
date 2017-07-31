

package main

import "testing"

func TestSum(t *testing.T) {  
    total := 2
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}