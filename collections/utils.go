package collections



func Map[Input any, Output any](items []Input, transform func(input Input) Output) []Output {

	resut := make([]Output, len(items))

	for index, item := range items {
		resut[index] = transform(item)
	}

	return resut
}


func ForEach[Input any](items []Input, action func(item Input)) {
	for _, item := range items {
		action(item)
	}
}

func Filter[Input any](items []Input, predicate func(item Input) bool) []Input {
	result := make([]Input,0)

	for _, item := range items {
		if(predicate(item)){
			result = append(result, item)
		}
	
	}
	return result
	}


func Find[Input any](items []Input, predicate func(item Input) bool) (Input, bool) {



	for _,item := range items{
		if(predicate(item)){
			return item, true
		}	
	}
	
	var zero Input
	return zero, false
}
