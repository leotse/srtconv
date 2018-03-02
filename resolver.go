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

		// first caption - just update state and continue
		if last == nil {
			last = copy(caption)
			last.ID = id
			res = append(res, last)
			id++
			continue
		}

		// overlapping captions! merge into one
		if caption.Start < last.End {
			last.Merge(caption)
			continue
		}

		// no overlapping, add to resulting caption
		addme := copy(caption)
		addme.ID = id
		res = append(res, addme)
		id++
	}

	return res
}
