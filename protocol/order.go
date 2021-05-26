package protocol

//无需定义接口的排序发方法
func Order_dept(list []*MSG_USER_Dept_cache, f func(a, b *MSG_USER_Dept_cache) bool) {
	max_len := len(list)
	tmp := make([]*MSG_USER_Dept_cache, max_len)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if f(list[i+1], list[i]) {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if f(list[i+2], list[i]) {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if f(list[i+3], list[i+1]) {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if f(list[i+2], list[i]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r int
	step = 4
	for step < max_len {
		step <<= 1
		for i := 0; i < max_len; i += step {
			l, r, max = i, i+step/2, i+step
			if max > max_len {
				max = max_len
			}
			for index := i; index < max; index++ {
				if l == step/2+i || (r < max && f(list[r], list[l])) {
					tmp[index] = list[r]
					r++
				} else {
					tmp[index] = list[l]
					l++
				}
			}
		}
		if step < max_len {
			for i := 0; i < max_len; i += step {
				l, r, max = i, i+step/2, i+step
				if max > max_len {
					max = max_len
				}
				for index := i; index < max; index++ {
					if l == step/2+i || (r < max && f(tmp[r], tmp[l])) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
				}
			}
		} else {
			copy(list, tmp)
		}
	}
}

//无需定义接口的排序发方法
func Order_product(list []*MSG_PROJECT_product_cache, f func(a, b *MSG_PROJECT_product_cache) bool) {
	max_len := len(list)
	tmp := make([]*MSG_PROJECT_product_cache, max_len)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if f(list[i+1], list[i]) {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if f(list[i+2], list[i]) {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if f(list[i+3], list[i+1]) {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if f(list[i+2], list[i]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r int
	step = 4
	for step < max_len {
		step <<= 1
		for i := 0; i < max_len; i += step {
			l, r, max = i, i+step/2, i+step
			if max > max_len {
				max = max_len
			}
			for index := i; index < max; index++ {
				if l == step/2+i || (r < max && f(list[r], list[l])) {
					tmp[index] = list[r]
					r++
				} else {
					tmp[index] = list[l]
					l++
				}
			}
		}
		if step < max_len {
			for i := 0; i < max_len; i += step {
				l, r, max = i, i+step/2, i+step
				if max > max_len {
					max = max_len
				}
				for index := i; index < max; index++ {
					if l == step/2+i || (r < max && f(tmp[r], tmp[l])) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
				}
			}
		} else {
			copy(list, tmp)
		}
	}
}
func Order_tree(list []*MSG_PROJECT_tree_cache, f func(a, b *MSG_PROJECT_tree_cache) bool) {
	max_len := len(list)
	tmp := make([]*MSG_PROJECT_tree_cache, max_len)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if f(list[i+1], list[i]) {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if f(list[i+2], list[i]) {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if f(list[i+3], list[i+1]) {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if f(list[i+2], list[i]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r int
	step = 4
	for step < max_len {
		step <<= 1
		for i := 0; i < max_len; i += step {
			l, r, max = i, i+step/2, i+step
			if max > max_len {
				max = max_len
			}
			for index := i; index < max; index++ {
				if l == step/2+i || (r < max && f(list[r], list[l])) {
					tmp[index] = list[r]
					r++
				} else {
					tmp[index] = list[l]
					l++
				}
			}
		}
		if step < max_len {
			for i := 0; i < max_len; i += step {
				l, r, max = i, i+step/2, i+step
				if max > max_len {
					max = max_len
				}
				for index := i; index < max; index++ {
					if l == step/2+i || (r < max && f(tmp[r], tmp[l])) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
				}
			}
		} else {
			copy(list, tmp)
		}
	}

}
func Order_project(list []*MSG_PROJECT_project_cache, f func(a, b *MSG_PROJECT_project_cache) bool) {
	if f == nil {
		f = func(a, b *MSG_PROJECT_project_cache) bool {
			if a.Order == b.Order {
				return a.Id > b.Id
			}
			return a.Order > b.Order
		}
	}
	max_len := len(list)
	tmp := make([]*MSG_PROJECT_project_cache, max_len)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if f(list[i+1], list[i]) {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if f(list[i+2], list[i]) {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if f(list[i+3], list[i+1]) {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if f(list[i+2], list[i]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r int
	step = 4
	for step < max_len {
		step <<= 1
		for i := 0; i < max_len; i += step {
			l, r, max = i, i+step/2, i+step
			if max > max_len {
				max = max_len
			}
			for index := i; index < max; index++ {
				if l == step/2+i || (r < max && f(list[r], list[l])) {
					tmp[index] = list[r]
					r++
				} else {
					tmp[index] = list[l]
					l++
				}
			}
		}
		if step < max_len {
			for i := 0; i < max_len; i += step {
				l, r, max = i, i+step/2, i+step
				if max > max_len {
					max = max_len
				}
				for index := i; index < max; index++ {
					if l == step/2+i || (r < max && f(tmp[r], tmp[l])) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
				}
			}
		} else {
			copy(list, tmp)
		}
	}
}
func Order_ascInt32(list []int32) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] > list[i+2] {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] > list[i+3] {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, index, n int
	tmp := make([]int32, max_len)
	step = 4
	for step < max_len {
		step <<= 1
		n++
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && list[l] > list[r]) {
						tmp[index] = list[r]
						r++
					} else {
						tmp[index] = list[l]
						l++
					}
				}
			}
		} else {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && tmp[l] > tmp[r]) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
				}
			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}
func Order_group(list []*MSG_USER_Group_cache, f func(a, b *MSG_USER_Group_cache) bool) {
	if f == nil {
		f = func(a, b *MSG_USER_Group_cache) bool {
			return a.Id > b.Id
		}
	}
	max_len := len(list)
	tmp := make([]*MSG_USER_Group_cache, max_len)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if f(list[i+1], list[i]) {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if f(list[i+2], list[i]) {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if f(list[i+3], list[i+1]) {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if f(list[i+2], list[i]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r int
	step = 4
	for step < max_len {
		step <<= 1
		for i := 0; i < max_len; i += step {
			l, r, max = i, i+step/2, i+step
			if max > max_len {
				max = max_len
			}
			for index := i; index < max; index++ {
				if l == step/2+i || (r < max && f(list[r], list[l])) {
					tmp[index] = list[r]
					r++
				} else {
					tmp[index] = list[l]
					l++
				}
			}
		}
		if step < max_len {
			for i := 0; i < max_len; i += step {
				l, r, max = i, i+step/2, i+step
				if max > max_len {
					max = max_len
				}
				for index := i; index < max; index++ {
					if l == step/2+i || (r < max && f(tmp[r], tmp[l])) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
				}
			}
		} else {
			copy(list, tmp)
		}
	}
}
func Order_branch(list []*MSG_PROJECT_branch_info, f func(a, b *MSG_PROJECT_branch_info) bool) {
	if f == nil {
		f = func(a, b *MSG_PROJECT_branch_info) bool {
			if a.Order == b.Order {
				return a.Id > b.Id
			}
			return a.Order > b.Order
		}
	}
	max_len := len(list)
	tmp := make([]*MSG_PROJECT_branch_info, max_len)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if f(list[i+1], list[i]) {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if f(list[i+2], list[i]) {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if f(list[i+3], list[i+1]) {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if f(list[i+2], list[i]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r int
	step = 4
	for step < max_len {
		step <<= 1
		for i := 0; i < max_len; i += step {
			l, r, max = i, i+step/2, i+step
			if max > max_len {
				max = max_len
			}
			for index := i; index < max; index++ {
				if l == step/2+i || (r < max && f(list[r], list[l])) {
					tmp[index] = list[r]
					r++
				} else {
					tmp[index] = list[l]
					l++
				}
			}
		}
		if step < max_len {
			for i := 0; i < max_len; i += step {
				l, r, max = i, i+step/2, i+step
				if max > max_len {
					max = max_len
				}
				for index := i; index < max; index++ {
					if l == step/2+i || (r < max && f(tmp[r], tmp[l])) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
				}
			}
		} else {
			copy(list, tmp)
		}
	}
}
func order_diff(list []diff) {
	f := func(a, b diff) bool {
		return a.key < b.key
	}
	max_len := len(list)
	tmp := make([]diff, max_len)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if f(list[i+1], list[i]) {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if f(list[i+2], list[i]) {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if f(list[i+3], list[i+1]) {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if f(list[i+2], list[i]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if f(list[i+2], list[i+1]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r int
	step = 4
	for step < max_len {
		step <<= 1
		for i := 0; i < max_len; i += step {
			l, r, max = i, i+step/2, i+step
			if max > max_len {
				max = max_len
			}
			for index := i; index < max; index++ {
				if l == step/2+i || (r < max && f(list[r], list[l])) {
					tmp[index] = list[r]
					r++
				} else {
					tmp[index] = list[l]
					l++
				}
			}
		}
		if step < max_len {
			for i := 0; i < max_len; i += step {
				l, r, max = i, i+step/2, i+step
				if max > max_len {
					max = max_len
				}
				for index := i; index < max; index++ {
					if l == step/2+i || (r < max && f(tmp[r], tmp[l])) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
				}
			}
		} else {
			copy(list, tmp)
		}
	}
}
