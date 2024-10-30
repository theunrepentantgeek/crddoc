package functions

// asList is a function that wraps an existing slice with metadata to make it
// easier to iterate over in templates.
func asList[I any](slice []I) []ListItem[I] {
	result := make([]ListItem[I], len(slice))
	for i, item := range slice {
		result[i] = ListItem[I]{
			Item:  item,
			First: i == 0,
			Last:  i == len(slice)-1,
		}
	}

	return result
}

type ListItem[I any] struct {
	// Item is the item in the list.
	Item I

	// First indicates whether this is the first item in the list.
	First bool

	// Last indicates whether this is the last item in the list.
	Last bool
}
