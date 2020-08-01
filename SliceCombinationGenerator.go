//This program takes a 2 dimensional slice of any type, and returns a 2 dimensional slice
//of the same type with all combinations for each element of each inner slice with each 
//element of the other inner slices, combinations where not all inner slices are used are included

//Example say you have a 2D slice of type [][]string{} and it looks like this:

//[][]string{[]string{"a0", "a1", "a2"}, []string{"b0", "b1"}, []string{"c0", "c1", "c2", "c3"}}

//the combinations will look something like this:

//[]string{a0} []string{a0, b0}, []string{a0, c0}, []string{a0, b0, c0}, []string{b0, c0}, []string{b0}

//THIS PROGRAM DOES NOT COMBINE THE ELEMENTS OF ONE INNER SLICE WITH ANOTHER

//for instance for the above example you will never get a combination []string{a0, a1} or []string{a1, a2}

//The idea of this program is that each inner slice is has unique elements that are not to be combined with each other
//but only to be combined with the unique elements of the other inner slices

//Throughout this code I have written comments as to where you need to rename the type of string to your slice type

//If you are using a type struct for your 2D array you will need to include the struct definition in this file
//so that the code can compare to that type

//To quick find my comments about where to insert your type in place of string use keyword REPLACETYPECASE

package main


import(

	"fmt"
	"math"
	"reflect"
)


func main() {

	//REPLACETYPECASE
	mySliceOriginal := [][]string{[]string{"a0", "a1", "a2"}, []string{"b0", "b1"}, []string{"c0", "c1", "c2", "c3"}}



	mySliceAsInterface := Convert2DSliceAnyTypeTo2DInterface(mySliceOriginal, len(mySliceOriginal))

	allCombosForElementsOfMySlice := AllCombinationsForElementsOfMultipleSlices(mySliceAsInterface)

	mySliceCombosAsOriginalType := [][]string{}



	innerType := reflect.TypeOf(mySliceAsInterface[0])

	for i := 0; i < len(allCombosForElementsOfMySlice); i++ {
		
			currentElementSlice := allCombosForElementsOfMySlice[i]

			//REPLACETYPECASE
			convertedTypeForCurrentElements := []string{}

			for j := 0; j < len(currentElementSlice); j++ {

				//REPLACETYPECASE	
				valueInner, ok := currentElementSlice[j].(string)

				if(ok){

					convertedTypeForCurrentElements = append(convertedTypeForCurrentElements, valueInner)
				}else{
					fmt.Println("error converting back to original data from interface, the type recognized was: \n", innerType, "\n at slice indices: \n", i, " ", j, "\n with value: \n", currentElementSlice[j])
					panic("error converting back to original data type")
				}

			}

		

		mySliceCombosAsOriginalType = append(mySliceCombosAsOriginalType, convertedTypeForCurrentElements)

	}


	//Print the resulting data for all combinations
	for i := 0; i < len(mySliceCombosAsOriginalType); i++ {

		fmt.Println(mySliceCombosAsOriginalType[i])

	}

	//TODO:

	//Here is where you can return the combination slice back to your main program



}

func AllCombinationsForElementsOfMultipleSlices(inputOptions [][]interface{})  [][]interface{}{



	bitSlice := []int{}
	breakpoints := []int{}

	breakPointTracker := 0

	for i := 0; i < len(inputOptions); i++ {

		currentInputOption := inputOptions[i]

		for j := 0; j < len(currentInputOption); j++ {

			if(j == 0 ){
				breakpoints = append(breakpoints, breakPointTracker)
			}

			bitSlice = append(bitSlice, 0)

			breakPointTracker++

		}


	}
	maxNumber := float64(0)

	lastIndex := len(bitSlice) - 1

	for i := 0; i < len(breakpoints); i++ {

		powerToRaiseTo := float64(lastIndex - breakpoints[i])

		maxNumber = maxNumber + math.Pow(float64(2), powerToRaiseTo)


	}

	currentNumber := 1

	returnValsIndices := [][]int{}

	for (currentNumber < int(maxNumber) + 1){
		satisfied, binarySlice := CheckModulosAreSatisied(breakpoints, currentNumber, len(bitSlice))
	
		if(satisfied){ 

			returnValsIndices = append(returnValsIndices, binarySlice)

		}
		currentNumber++
	}

	allElementsToOneSlice := []interface{}{}



	for i := 0; i < len(inputOptions); i++ {
		allElementsToOneSlice = append(allElementsToOneSlice, inputOptions[i]...)
	}

	if(len(allElementsToOneSlice) != len(returnValsIndices[0])){
		fmt.Println("Length Of Elements Slice ", len(allElementsToOneSlice), " Length Of Indices Slice ", len(returnValsIndices))
		panic("more indices than total elements")
	}


	returnVals := [][]interface{}{}

	for i := 0; i < len(returnValsIndices); i++ {

		currentIndices := returnValsIndices[i]

		currentCombination := []interface{}{}

		for j := 0; j < len(currentIndices); j++ {

			currentBit := currentIndices[j]

			if(currentBit == 0){
				continue
			}else{
				currentCombination = append(currentCombination, allElementsToOneSlice[j])
			}

		}

		returnVals = append(returnVals, currentCombination)

	}


	return returnVals
	
}


func CheckModulosAreSatisied(breakpoints []int, number int, numberRequiredBinaryLength int) (bool, []int) {



	binaryString := fmt.Sprintf("%b", number)

	for len(binaryString) < numberRequiredBinaryLength{
		binaryString = "0" + binaryString
	}

	binarySlice := []int{}

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
	horizontalCursor := len(binarySlice) - 1

	doneCheckingMods := false

	lastIndex := len(binarySlice) - 1

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





func IsABreakPoint(testVal int, breakpoints []int) bool {

	for i := 0; i < len(breakpoints); i++ {
		if(breakpoints[i] == testVal){
			return true
		}
	}
	return false
}



func Convert2DSliceAnyTypeTo2DInterface(inputSlice interface{}, lengthReturnSlice int) [][]interface{}{

	
	reflectType := reflect.TypeOf(inputSlice)

	switch reflectType.Kind() {
		case reflect.Slice:
			elementType := reflectType.Elem()
			switch elementType.Kind(){
				case reflect.Slice:
				
					valueOf2DSlice := reflect.ValueOf(inputSlice)

					//REPLACETYPECASE
					typeString := reflect.TypeOf([][]string{})

					firstConversion2DSlice := valueOf2DSlice.Convert(typeString)

					//REPLACETYPECASE
					finalConversion2DSlice := firstConversion2DSlice.Interface().([][]string)


					returnData := [][]interface{}{}


					for i := 0; i < len(finalConversion2DSlice); i++ {

						firstElementInner := finalConversion2DSlice[i]

						interfaceSliceInner := []interface{}{}

						for j := 0; j < len(firstElementInner); j++ {
							interfaceSliceInner = append(interfaceSliceInner, firstElementInner[j])
						}

						returnData = append(returnData, interfaceSliceInner)
					}

					return returnData
				default:
					panic("error, not a valid 2D slice, outer type is a slice, but inner type is not")	
			}
		default:
			panic("error, not a valid 2D slice, outermost type is not of any slice")	
			
	}

	



	return [][]interface{}{}

}











