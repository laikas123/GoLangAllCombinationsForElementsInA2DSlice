package main


import(

	"fmt"
	"math"
	"reflect"
)


func main() {

	mySliceOriginal := [][]string{[]string{"a0", "a1", "a2"}, []string{"b0", "b1"}, []string{"c0", "c1", "c2", "c3"}}



	mySliceAsInterface := Convert2DSliceAnyTypeTo2DInterface(mySliceOriginal, len(mySliceOriginal))


	fmt.Println()
	fmt.Printf("%#v\n", mySliceAsInterface[0])

	panic("argghh")

	allCombosForElementsOfMySlice := AllCombinationsForElementsOfMultipleSlices(mySliceAsInterface)

	mySliceCombosAsOriginalType := [][]string{}

	valueInterface := [][]interface{}{[]interface{}{"hello"}}


	fmt.Printf("%#v\n", allCombosForElementsOfMySlice[0])
	fmt.Printf("%#v\n", valueInterface[0])


	panic("yippyskippy")

	fmt.Println("rowoworow", allCombosForElementsOfMySlice[0][0])
	
	

	itemcheck := valueInterface[0][0]



	// typeString := reflect.TypeOf("stir")

	value, ok := itemcheck.(string)



	switch itemcheck.(type){
		case string:
			fmt.Println("ok ok ok ")

	}


	panic("error error error ")
	fmt.Printf("%t\n", itemcheck)

	fmt.Println(ok, value)

	// panic("check check check")


	// outerType := reflect.TypeOf(mySliceAsInterface)

	// innerType := reflect.TypeOf(mySliceAsInterface[0])

	for i := 0; i < len(allCombosForElementsOfMySlice); i++ {
		
			currentElementSlice := allCombosForElementsOfMySlice[i]

			fmt.Println("row", currentElementSlice)

			convertedTypeForCurrentElements := []string{}

			for j := 0; j < len(currentElementSlice); j++ {

					
					valueInner, ok := currentElementSlice[j].(string)

					

					// typeOfString := reflect.TypeOf("string")



					innerValue := currentElementSlice[j]

					reflectValue := reflect.ValueOf(innerValue)

					fmt.Println(reflectValue.CanInterface())


					if(reflectValue.CanInterface()){
						//fmt.Println(reflectValue.Interface())
						fmt.Printf("%#v\n", reflectValue.Interface())
						// stringcast := string(reflectValue.Interface())
						// fmt.Println(stringcast)
						value, ok := reflectValue.Interface().(string)
						if(ok){
							fmt.Println("Yippeee", value)
						}else{
						//	fmt.Println(reflect.ValueOf(reflectValue.Interface()))

							// stringType := reflect.TypeOf("string")
							// interfacedVal := reflectValue.Interface()

						}
					}


					panic("good panic")

					
					reflectValue = reflect.ValueOf(innerValue)


					// reflectValueString := reflectValue.String()

					// // reflectValue = interface{}(reflectValue)

					// fmt.Println(reflectValue.SetString())

					fmt.Println(reflect.TypeOf(reflectValue))

					panic("check")

					// interfaceValue, ok0 := reflectValue.(interface{})

					// if(ok0){
					// 	reflectValue = interfaceValue.Interface()

					// 	asserted, ok1 := reflectValue.(string)

					// 	if(ok1){
					// 		fmt.Println("YIPEEE", asserted)
					// 	}
					// }
					

					// reflectValueConverted := reflectValue.Convert(typeOfString)

					// fmt.Println(reflectValueConverted)

					// /mySlice := []string{reflectValue}

					reflectType := reflect.TypeOf(reflectValue)

					// fmt.Println(reflectType)
					panic("well darn")

					switch reflectType.Kind(){
						case reflect.String:
							fmt.Println("hello hterherhekjrhekjh")
							panic("ok now tis all good") 
					}

					


					fmt.Println("elenent", currentElementSlice)

					if(ok){
						panic("ok ok ok ")
						convertedTypeForCurrentElements = append(convertedTypeForCurrentElements, valueInner)
					}else{
						fmt.Printf("%f\n", currentElementSlice[j])
						fmt.Println(reflect.ValueOf(currentElementSlice[j]))
						fmt.Println(reflect.TypeOf(valueInner))
						panic("super panic")
					}

		}

		fmt.Println(convertedTypeForCurrentElements)
		panic("errororr")
		

		mySliceCombosAsOriginalType = append(mySliceCombosAsOriginalType, convertedTypeForCurrentElements)

	}

	for i := 0; i < len(mySliceCombosAsOriginalType); i++ {

		fmt.Println(mySliceCombosAsOriginalType[i])

	}

	fmt.Println(reflect.TypeOf(mySliceCombosAsOriginalType))



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
		fmt.Println(len(allElementsToOneSlice), len(returnValsIndices))
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
			

				fmt.Println(currentAdditionToSummation)

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
		fmt.Println(horizontalCursor)
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

	fmt.Printf("%t", inputSlice)


	reflectType := reflect.TypeOf(inputSlice)

	switch reflectType.Kind() {
		case reflect.Slice:
			fmt.Println("YES")
			elementType := reflectType.Elem()
			switch elementType.Kind(){
				case reflect.Slice:
					fmt.Println("YES2")
					// fmt.Println(reflectType.Slice(0, 1))
				
					valueOf2DSlice := reflect.ValueOf(inputSlice)

					firstElement := valueOf2DSlice

					typeString := reflect.TypeOf([][]string{})

					firstElementConverted := firstElement.Convert(typeString)

					firstElementConvertedFinal := firstElementConverted.Interface().([][]string)


					returnData := [][]interface{}{}


					for i := 0; i < len(firstElementConvertedFinal); i++ {

						firstElementInner := firstElementConvertedFinal[i]

						interfaceSliceInner := []interface{}{}

						for j := 0; j < len(firstElementInner); j++ {
							interfaceSliceInner = append(interfaceSliceInner, firstElementInner[j])
						}

						returnData = append(returnData, interfaceSliceInner)
					}


					fmt.Println(firstElementConvertedFinal)



					// typeAssert, ok := firstElement.([]string)

					// if(ok){
					// 	fmt.Println(typeAssert)
					// 	panic("good good")
					// }


					fmt.Printf("%#v\n", returnData[0])
					panic("yarhoo")

					// interfaceof2dslice := valueOf2DSlice.Interface()




					// myArray := []interface{}{interfaceof2dslice}



					// fmt.Println(myArray[0][0])

					panic("earliest arghhh")


					fmt.Println(valueOf2DSlice.Slice(0, 3))

					returnData = [][]interface{}{}

					for i := 0; i < valueOf2DSlice.Len(); i++ {

						indexValue := valueOf2DSlice.Index(i)

						interfaceSlice := []interface{}{}

						fmt.Printf("%#v\n", indexValue)

						myValue := indexValue.Index(i)

						


						fmt.Printf("%#t", myValue)

						// panic("ok")


						for i := 0; i < indexValue.Len(); i++ {
							
							interfaceSlice = append(interfaceSlice, indexValue.Index(i))
						}

						returnData = append(returnData, interfaceSlice)
					}	

					fmt.Printf("%#v\n", returnData[0])

					panic("ealier arghh")

					return returnData
				default:
					panic("error, not a valid 2D slice, outer type is a slice, but inner type is not")	
			}
		default:
			panic("error, not a valid 2D slice, outermost type is not of any slice")	
			
	}

	



	return [][]interface{}{}

}











