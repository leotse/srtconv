package srtfix

// Resolve overlapping timestamps by merging captions
func Resolve(captions []*Caption) []*Caption {

	// empty captions?
	if len(captions) == 0 {
		return []*Caption{}
	}

	// iterate captions and merge overlapping ones
	// assumes the captions are order chronologically
	var id = 1
	var last *Caption
	res := make([]*Caption, 0)
	for _, caption := range captions {

		// overlapping captions! merge into last
		if last != nil && caption.Start < last.End {
			last.Merge(caption)
			continue
		}

		// otherwise.. just add the caption to results + update last pointer
		last = copy(caption)
		last.ID = id
		res = append(res, last)
		id++
	}

	return res
}
