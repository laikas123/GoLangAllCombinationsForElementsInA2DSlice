//this program takes a 2D slice of any kind and any length, with its inner slices being of any length
//it treats each inner slice as its own "type" and returns all combinations such that one combination 
//never has the same type twice.. 

//take the following example:

//[][]string{[]string{"dog black", "dog white", "dog brown"}, []string{"cat orange"}, []string{"fish blue", "fish red"}}

//the code will give combintions such as []string{"dog black, cat orange"}, []string{"cat orange, fish red"}

//each combination must have at least 1 animal, and can have up to 3 (or however many animals there are), but 
//the important thing there can never be two animals of any type in the same combination, so never 2 dogs, 3 fish etc..
//this program generates every possible combination of this type for any arbitrary length 2D slice, with each of its
//inner slices being of arbitrary length 

package main


import(
	"fmt"
	"math"
)

//IMPORTANT in the code it currently is set to use type [][]string as the 2D slice type by default
//to use any slice type look for the comments containing "REPLACETYPE" (no quotes) and put the desired 
//2D slice type there, note for 2D slices using a struct, add the struct definition to this program

func main() {

	//REPLACETYPE use your type of data instead of string, for custom struct you will need to add the struct
	//definition to this file 
	originalSlice := [][]string{[]string{"a0", "a1", "a2"}, []string{"b0", "b1", "b2"}, []string{"c0", "c1", "c2", "c3"}}

	//this is a 1D representation of the original slice with identical data 
	//so for the sample data above it would be
	//[]string{"a0", "a1", "a2", "b0", "b1", "b2", "c0", "c1", "c2", "c3"}
	//this allows for indexing combinations using a "binary cursor" 
	//as seen throughout the code
	originalSliceTo1DSlice := []string{}


	for i := 0; i < len(originalSlice); i++ {
		originalSliceTo1DSlice = append(originalSliceTo1DSlice, originalSlice[i]...)
	}

	//a bit slice is used to index combinations from the 1D representation
	bitCursor := []int{}
	
	//break points are where type ends into another
	//for the example slice:
	//
	//[]string{"a0", "a1", "a2", "b0", "b1", "b2", "c0", "c1", "c2", "c3"}
	//
	//indices:   0     1     2     3     4     5     6     7     8     9
	//
	//the breakpoints are 
	//0, 3, and 5 as from left to right these are where "new inner slice beginnings"
	//this is more easily seen in the original 2D representation:
	//
	//[][]string{[]string{"a0", "a1", "a2"}, []string{"b0", "b1", "b2"}, []string{"c0", "c1", "c2", "c3"}}
	//
	//indices:              0     1     2               3     4     5               6     7     8     9
	//
	//these breakpoints help the cursor ensure it never chooses a combination with two 
	//elements from the same inner slice
	breakpoints := []int{}

	breakPointTracker := 0


	for i := 0; i < len(originalSlice); i++ {

		currentInnerSlice := originalSlice[i]

		for j := 0; j < len(currentInnerSlice); j++ {

			//only append breakpoint for each new inner slice beginning
			if(j == 0 ){
				breakpoints = append(breakpoints, breakPointTracker)
			}
			bitCursor = append(bitCursor, 0)

			breakPointTracker++

		}


	}

	//this return a 2D slice of type [][]int
	//each []int contained is a binary cursor indexing the 1D representation
	//and each binary cursor ensures a unique combination is chosen
	//all combinations will be generated
	allCombinationIndices := AllCombinationsIndices(breakpoints, bitCursor, len(bitCursor))

	//REPLACETYPE use your type of data instead of string, for custom struct you will need to add the struct
	//definition to this file
	outputCombinationsSlice := [][]string{}

	for i := 0; i < len(allCombinationIndices); i++ {

		currentIndices := allCombinationIndices[i]

		combinationToAppend := []string{}

		for j := 0; j < len(currentIndices); j++ {
			if(currentIndices[j] == 1){
				combinationToAppend = append(combinationToAppend, originalSliceTo1DSlice[j])
			}
		}

		//IMPORTANT this is slice of interest that should be returned to your main program
		outputCombinationsSlice = append(outputCombinationsSlice, combinationToAppend)

	}
	//Print the resulting data for all combinations to check the data looks correct
	for i := 0; i < len(outputCombinationsSlice); i++ {

		fmt.Println(outputCombinationsSlice[i])

	}

	//TODO:

	//Here is where you can return the combination slice back to your main program



}

func AllCombinationsIndices(breakpoints []int, bitCursor []int, totalElementsToChooseFrom int) [][]int{

	//this will be the maximum value for the binary cursor
	//it is set below using the break points set to high bits
	//for a 9 bit such as the example data with break points at 
	//indices 0, 3, 5 the highest value will be:
	//100101000
	//which is 296
	maxNumber := float64(0)

	lastIndex := len(bitCursor) - 1

	for i := 0; i < len(breakpoints); i++ {

		powerToRaiseTo := float64(lastIndex - breakpoints[i])

		maxNumber = maxNumber + math.Pow(float64(2), powerToRaiseTo)


	}

	currentNumber := 1

	//unique binary cursors to return
	returnValsIndices := [][]int{}

	for (currentNumber < int(maxNumber) + 1){
		satisfied, binaryCursor := CheckModulosAreSatisfied(breakpoints, currentNumber, len(bitCursor))
	
		if(satisfied){ 

			returnValsIndices = append(returnValsIndices, binaryCursor)

		}
		currentNumber++
	}

	return returnValsIndices
}

//this functions uses a binary string to an []int slice which holds the bits of the binary string
//the function also uses the modulus operator to ensure that there are never two high bits set in between
//break points
func CheckModulosAreSatisfied(breakpoints []int, number int, totalElementsToChooseFrom int) (bool, []int) {


	//use an input number and convert it to binary string
	binaryString := fmt.Sprintf("%b", number)

	//make sure the binary string is long as the elements to choose from
	for len(binaryString) < totalElementsToChooseFrom{
		binaryString = "0" + binaryString
	}

	//convert the binary string to have its bits be held in an []int slice
	//the left most bit is the most significant
	binaryCursor := []int{}

	for i := 0; i < len(binaryString); i++ {

		currentBit := rune(binaryString[i])

		if(currentBit == '0'){
			binaryCursor = append(binaryCursor, 0)
		}else if(currentBit == '1'){
			binaryCursor = append(binaryCursor, 1)
		}else{
			panic("unkown bit CheckModulosAreSatisfied()")
		}
	}

	//used to help check the binaryCursor 
	//won't yield an invalid combination
	horizontalCursor := len(binaryCursor) - 1

	doneCheckingMods := false

	lastIndex := len(binaryCursor) - 1


	//this for loop uses the modulus operator along with the generated breakpoints
	//to ensure that there is never two high bits inbetween break points
	//this would mean that two elements of the same inner slice of the original
	//2D slice would be in a combination together violating the expected output
	for !doneCheckingMods {

		currentModDoneBeingChecked := false

		summation := float64(0)

		for !currentModDoneBeingChecked {

			currentAdditionToSummation := math.Pow(2, float64(lastIndex - horizontalCursor))


			if(binaryCursor[horizontalCursor] == 1){
		
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
					return false, binaryCursor
				}

			}else{
				horizontalCursor = horizontalCursor - 1
			}


		}
		
		if(horizontalCursor == -1){
			return true, binaryCursor
		}

	}

	return true, binaryCursor
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









