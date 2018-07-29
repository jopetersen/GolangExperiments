package main
//https://www.youtube.com/watch?v=i3o4G4bmqPc
import ("fmt")

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

// making a method to calculate km per hour

func (c car) kmh() float64{
// if adding () after func, will associate it with a struct, making it a method
  return float64(c.gas_pedal)* (c.top_speed_kmh/usixteenbitmax)
    // converting to float 64


}




type car struct{
  gas_pedal uint16
  brake_pedal uint16
  steering_wheel uint16 // -32k - +32k
  top_speed_kmh float64
}

func main(){
  a_car := car{
                gas_pedal: 22341,
                brake_pedal: 0,
                steering_wheel: 12561,
                top_speed_kmh: 225.0}
  fmt.Println(a_car.gas_pedal)
  fmt.Println(a_car.kmh())
}
