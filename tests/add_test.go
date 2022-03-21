package main

import "testing"

func TestAdd(t *testing.T) {

    t.Run("adding two positive numbers", func(t *testing.T) {
        sum := addition(2, 2)
        expected := 4
        
        if sum != expected {
            t.Errorf("Expected %d; but got %d", expected, sum)
        }
    })
    
    t.Run("adding two negative numbers", func(t *testing.T) {
        sum := addition(-3, -4)
        expected := -7

        if sum != expected {
            t.Errorf("Expected %d; but got %d", expected, sum)
        }
    })

    t.Run("adding one positive and one negative integer", func(t *testing.T) {
        sum := addition(1, -3)
        expected := -2

        if sum != expected {
            t.Errorf("Expected %d; but got %d", expected, sum)
        }
    })
    
}
