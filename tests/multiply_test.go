package main

import "testing"

func TestMultiply(t *testing.T) {

    t.Run("multiplying two positive numbers", func(t *testing.T) {
        sum := Multiplication(2, 2)
        expected := 4
        
        if sum != expected {
            t.Errorf("Expected %d; but got %d", expected, sum)
        }
    })
    
    t.Run("multiplying two negative numbers", func(t *testing.T) {
        sum := Multiplication(-3, -4)
        expected := 12

        if sum != expected {
            t.Errorf("Expected %d; but got %d", expected, sum)
        }
    })

    t.Run("multiplying one positive and one negative integer", func(t *testing.T) {
        sum := Multiplication(1, -3)
        expected := -3

        if sum != expected {
            t.Errorf("Expected %d; but got %d", expected, sum)
        }
    })
    
}