//This program takes an arbitrary length 2 dimensional slice of any type, and returns a 2 dimensional slice
//of the same type with all combinations for each element of each inner slice with each 
//element of the other inner slices, combinations where not all inner slices are used are included

//Example say you have a 2D slice of type [][]string{} and it looks like this:

//(NOTE for describing this code the example below is used throughout the program, but the same
//principles apply to any length 2D slice with any length inner slices)

//[][]string{[]string{"a0", "a1", "a2"}, []string{"b0", "b1"}, []string{"c0", "c1", "c2", "c3"}}

//the combinations will look something like this:

//[]string{a0} []string{a0, b0}, []string{a0, c0}, []string{a0, b0, c0}, []string{b0, c0}, []string{b0} etc...

//THIS PROGRAM DOES NOT COMBINE THE ELEMENTS OF ONE INNER SLICE WITH ANOTHER

//for instance for the above example you will never get a combination []string{a0, a1} or []string{a1, a2}

//in case the above example is too discrete take this example:

//[][]string{[]string{"dog black", "dog white", "dog brown"}, []string{"cat orange"}, []string{"fish blue", "fish red"}}

//this program would match combinations of animals, but no combination would include two cats, two dogs, or two fish
//a combination can lack a cat, fish or dog, but it can never have too of the same animal type
//the rest of the comments will now go back to the simpler example:
//[][]string{[]string{"a0", "a1", "a2"}, []string{"b0", "b1"}, []string{"c0", "c1", "c2", "c3"}}

//The idea of this program is that each inner slice is has unique elements that are not to be combined with each other
//but only to be combined with the unique elements of the other inner slices

//Throughout this code I have written comments as to where you need to rename the type of string to your slice type
//using the keyword REPLACETYPE you can search and find these lines

//If you are using a type struct for your 2D array you will need to include the struct definition in this file
//so that the code can compare to that type

//IMPORTANT as mentioned above, by default this code uses the 2D slice type [][]string
//To quick find comments where to insert your type in place of string use keyword REPLACETYPE use your 
//type of data instead of string,for custom struct you will need to add the struct definition to this file

package main


import(

	"fmt"
	"math"
	"reflect"
)


func main() {

	//REPLACETYPE use your type of data instead of string, for custom struct you will need to add the struct
	//definition to this file 
	mySliceOriginal := [][]string{[]string{"a0", "a1", "a2"}, []string{"b0", "b1", "b2"}, []string{"c0", "c1", "c2", "c3"}}

	//REPLACETYPE use your type of data instead of string, for custom struct you will need to add the struct
	//definition to this file
	mySliceAllValsToOneInnerArray := []string{}


	//the "mySliceAllValsToOneInnerArray" for the example:
	//[][]string{[]string{"a0", "a1", "a2"}, []string{"b0", "b1"}, []string{"c0", "c1", "c2", "c3"}}
	//would become 
	//[][]string{[]string{"a0", "a1", "a2", "b0", "b1", "c0", "c1", "c2", "c3"}}
	//as seen later in the program this allows the binarySlice of bits to use "breakpoints"
	//to smoothly find combinations
	for i := 0; i < len(mySliceOriginal); i++ {
		mySliceAllValsToOneInnerArray = append(mySliceAllValsToOneInnerArray, mySliceOriginal[i]...)
	}


	
	//allCombo has inner elements like this:
	//[]int{0, 0, 0, 1, 0, 1, 0, 0}
	//[]int{0, 0, 1, 0, 0, 0, 0, 0}
	//which map to the "mySliceAllValsToOneInnerArray" which as commented above is:
	//[][]string{[]string{"a0", "a1", "a2", "b0", "b1", "c0", "c1", "c2", "c3"}}
	//as seen in the function AllCombinationIndices() by using modulus operator %
	//its easy to return only indices that will never include 2 a values, 2 b values, or 2 c values
	allComboIndices := AllCombinationsIndices(mySliceOriginal) 

	//REPLACETYPE use your type of data instead of string, for custom struct you will need to add the struct
	//definition to this file
	allCombosSlice := [][]string{}

	for i := 0; i < len(allComboIndices); i++ {

		currentIndices := allComboIndices[i]

		//REPLACETYPE use your type of data instead of string, for custom struct you will need to add the struct
		//definition to this file
		comboToAppend := []string{}

		for j := 0; j < len(currentIndices); j++ {
			if(currentIndices[j] == 1){
				comboToAppend = append(comboToAppend, mySliceAllValsToOneInnerArray[j])
			}
		}

		//IMPORTANT this is slice of interest that should be returned to your main program
		allCombosSlice = append(allCombosSlice, comboToAppend)

	}
	//Print the resulting data for all combinations to check the data looks correct
	for i := 0; i < len(allCombosSlice); i++ {

		fmt.Println(allCombosSlice[i])

	}

	//TODO:

	//Here is where you can return the combination slice back to your main program



}

func AllCombinationsIndices(inputSlice interface{})  [][]int{


	//since the program should be able to use any type of 2D slice use
	//reflect to just ensure that its any 2D slice
	//again in the main function change the input and output slices to match your data type
	//if you are using a custom struct add that struct definition to this file so the program will
	//work
	reflectType := reflect.TypeOf(inputSlice)

	//this will be the inputSlice parameter to this function
	//but temporarily in interface form
	inputOptions := [][]interface{}{}

	switch reflectType.Kind() {
		case reflect.Slice:
			elementType := reflectType.Elem()
			switch elementType.Kind(){
				case reflect.Slice:
				
					valueOf2DSlice := reflect.ValueOf(inputSlice)

					//REPLACETYPE use your type of data instead of string, for custom struct you will need to add the struct
					//definition to this file
					typeString := reflect.TypeOf([][]string{})

					firstConversion2DSlice := valueOf2DSlice.Convert(typeString)

					//REPLACETYPE use your type of data instead of string, for custom struct you will need to add the struct
					//definition to this file
					finalConversion2DSlice := firstConversion2DSlice.Interface().([][]string)



					for i := 0; i < len(finalConversion2DSlice); i++ {

						firstElementInner := finalConversion2DSlice[i]

						interfaceSliceInner := []interface{}{}

						for j := 0; j < len(firstElementInner); j++ {
							interfaceSliceInner = append(interfaceSliceInner, firstElementInner[j])
						}

						inputOptions = append(inputOptions, interfaceSliceInner)
					}

					
				default:
					panic("error, not a valid 2D slice, outer type is a slice, but inner type is not")	
			}
		default:
			panic("error, not a valid 2D slice, outermost type is not of any slice")	
			
	}

	//the bit slice is a slice that matches the length of 
	//all the original elements of the original slice input to the entire program
	//combined into one slice 
	//sticking to the original input example: 
	//[][]string{[]string{"a0", "a1", "a2"}, []string{"b0", "b1"}, []string{"c0", "c1", "c2", "c3"}}
	//all of its elements in one slice looks like:
	//[][]string{[]string{"a0", "a1", "a2", "b0", "b1", "c0", "c1", "c2", "c3"}}
	//the bit slice indexes the latter slice above in a careful way to ensure
	//no two types for example a, b,  or c are included in the same combination 
	bitSlice := []int{}
	

	//following the example:
	//[][]string{[]string{"a0", "a1", "a2", "b0", "b1", "c0", "c1", "c2", "c3"}}
	//with indices ---->    0     1     2     3     4     5     6     7     8
	//the break points are 0, 3, and 5
	//these are points where the data type changes from one to the next
	breakpoints := []int{}

	//this value is used to calculate where breakpoints are as shown in the comment above
	//the for loop below uses it
	breakPointTracker := 0


	//input options would be the entire parent slice
	//[][]string{[]string{"a0", "a1", "a2"}, []string{"b0", "b1"}, []string{"c0", "c1", "c2", "c3"}}
	for i := 0; i < len(inputOptions); i++ {

		//the first option at index 0 is []string{"a0", "a1", "a2"}
		//the second option at index 1 is []string{"b0", "b1"}
		//the third option at index 2 is []string{"c0", "c1", "c2", "c3"}
		currentInputOption := inputOptions[i]

		for j := 0; j < len(currentInputOption); j++ {

			//only append breakpoint for each new inner slice beginning
			if(j == 0 ){
				breakpoints = append(breakpoints, breakPointTracker)
			}

			//start with all 0 bits in the bit slice
			bitSlice = append(bitSlice, 0)

			//increment breakPoint tracker for each new element
			breakPointTracker++

		}


	}

	//max number is the max value the bitslice can be
	//one thing to point out
	//for the slice:
	//[][]string{[]string{"a0", "a1", "a2", "b0", "b1", "c0", "c1", "c2", "c3"}}
	//the maximum techincal value for an 9 bit cursor is obviously
	//111111111
	//however selecting all of the indices would break the rule because that would 
	//make a combination with multiple a, b, and c values which is not the goal
	//the maximum value for this program is
	//100101000
	//which would yield the combination:
	//[]string{"a0", "b0", "c0"}
	maxNumber := float64(0)

	lastIndex := len(bitSlice) - 1

	//this for loop calculates the binary value at each breakpoint
	//for the slice
	//[][]string{[]string{"a0", "a1", "a2", "b0", "b1", "c0", "c1", "c2", "c3"}}
	//with indices ---->   0      1     2     3     4     5     6     7     8 
	//the breakpoints are 0, 3, and 5
	//imagining a 9 bit string with these bits set to 1 you would get:
	//100101000
	//the 1 bits would be equal to 256, 32, and 8 respectively
	//this becomes useful in the CheckModulosAreSatisfied() function
	//the following for loop just caluclates these values and adds them together
	//to give the maximum allowed value for the bit slice
	for i := 0; i < len(breakpoints); i++ {

		powerToRaiseTo := float64(lastIndex - breakpoints[i])

		maxNumber = maxNumber + math.Pow(float64(2), powerToRaiseTo)


	}

	//this number starts the binary string at 000000001
	//and the for loop increments to max val  100101000
	currentNumber := 1

	returnValsIndices := [][]int{}

	for (currentNumber < int(maxNumber) + 1){
		satisfied, binarySlice := CheckModulosAreSatisfied(breakpoints, currentNumber, len(bitSlice))
	
		if(satisfied){ 

			returnValsIndices = append(returnValsIndices, binarySlice)

		}
		currentNumber++
	}

	//this slice contains all possible bit slices
	//that will be used to get all unqiue combinations
	return returnValsIndices

}


func CheckModulosAreSatisfied(breakpoints []int, number int, numberRequiredBinaryLength int) (bool, []int) {


	//the binary string is a string of bits
	//the length is the numberRequiredBinaryLength parameter to the function
	//which is the summation of the total elements from the 2D slice input to the program
	//its job is to index the 2D slice and never index two elements from the same inner slice
	//it creates unique combinations by incrementing from a binary 1 to the maximum allowed
	//number which is touched upon later
	binaryString := fmt.Sprintf("%b", number)


	//append 0 to the binary string so it meets length requirements
	for len(binaryString) < numberRequiredBinaryLength{
		binaryString = "0" + binaryString
	}

	//the binary slice is essentially the binary string
	//but in an int format, so the process is increment an integer
	//get its binary string, convert that string to a binary slice
	//which is really just the list of bits from the binary string
	binarySlice := []int{}

	//here is where the binary string gets converted to its slice representation
	//the leftmost bit (index 0) is most significant
	for i := 0; i < len(binaryString); i++ {

		currentBit := rune(binaryString[i])

		if(currentBit == '0'){
			binarySlice = append(binarySlice, 0)
		}else if(currentBit == '1'){
			binarySlice = append(binarySlice, 1)
		}else{
			panic("what kind of madness bit is this")
		}
	}

	//the horizontal cursor helps move through the binary slice smoothly
	//and keep track of break points
	//break points represent points in the original slice where inner slices end and start
	//take the example 
	//[][]string{[]string{"a0", "a1", "a2"}, []string{"b0", "b1"}, []string{"c0", "c1", "c2", "c3"}}
	//with indices ----->  0      1     2               3     4               5     6     7     8
	//the break points are 0, 3, 5
	//this is because the horizontal cursor moves from right to left
	horizontalCursor := len(binarySlice) - 1

	doneCheckingMods := false

	lastIndex := len(binarySlice) - 1

	//THIS FOLLOWING FOR LOOP IS THE IMPORTANT PART..The magic..
	//the for loop below does most of the magic
	//it starts at 000000001 and increments up to 100101000
	//but doesn't accept values that might yield a combination with two 
	//elements of the same type
	//for example an ok bit slice would be
	//0  0  1  0  1  0  0  0  0
	//for our example:
	//[][]string{[]string{"a0", "a1", "a2", "b0", "b1", "c0", "c1", "c2", "c3"}}
	//this would get the values
	//a2, and b1 
	//a bit slice that is not ok is:
	//1  0  0  1  1  0  0  0  0
	//this would get 
	//a0, b0, and b1 which would include 2 b value types and break the rule
	//how does the for loop recognize this?
	//via the modulus operator
	//take the above example bit slice
	//1  0  0  1  1  0  0  0  0
	//the high bits are 256, 32, and 16 respectively
	//this is where the breakpoints help
	//one of the breakpoints was defined at index 3
	//aka the high bit equal to 32
	//the next break point would be at index 5
	//what the for loop does it look at the values
	//for  3 <= indices < 5
	//in this case indices 3 and 4 
	//it adds up the high bits, so 32 and 16 which equals 48
	//then perform a modulus for the value of the break point
	//in this case the break point at index 3 is 32
	//so 32%48 = 32
	//this fails the test, only a 0 remainder is a pass
	//to see why this works take the example where only 32 is high
	//or only 16 is high 
	//32%32 = 0 32%16 = 0 
	//this even works for trickier examples....
	//what it does is checks that only 1 bit is high for between the breakpoints
	//and this is what gaurantess unique combinations are acieved
	for !doneCheckingMods {

		currentModDoneBeingChecked := false

		summation := float64(0)

		for !currentModDoneBeingChecked {

			currentAdditionToSummation := math.Pow(2, float64(lastIndex - horizontalCursor))


			if(binarySlice[horizontalCursor] == 1){
		
				summation = summation + currentAdditionToSummation

			}

			if(IsABreakPoint(horizontalCursor, breakpoints)){

				intSummation := int(summation)

				intCurrentAdditionToSummation := int(currentAdditionToSummation)

				if(intSummation == 0){
					horizontalCursor = horizontalCursor - 1
					currentModDoneBeingChecked = true
				}else if( (intCurrentAdditionToSummation % intSummation == 0)){
					horizontalCursor = horizontalCursor - 1

					currentModDoneBeingChecked = true

				}else{
					return false, binarySlice
				}

			}else{
				horizontalCursor = horizontalCursor - 1
			}


		}
		
		if(horizontalCursor == -1){
			return true, binarySlice
		}

	}

	return true, binarySlice
}




//simple function test if current value is in breakpoints slice
func IsABreakPoint(testVal int, breakpoints []int) bool {

	for i := 0; i < len(breakpoints); i++ {
		if(breakpoints[i] == testVal){
			return true
		}
	}
	return false
}









