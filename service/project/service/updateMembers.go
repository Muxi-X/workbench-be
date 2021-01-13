package service

import (
	"context"

	errno "muxi-workbench-project/errno"
	"muxi-workbench-project/model"
	pb "muxi-workbench-project/proto"
	m "muxi-workbench/model"
	e "muxi-workbench/pkg/err"
)

// UpdateMembers ... 更新项目成员列表
func (s *Service) UpdateMembers(ctx context.Context, req *pb.UpdateMemberRequest, res *pb.ProjectNameAndIDResponse) error {

	// 这里需要 diff 老的项目成员列表，出 add list 和 delete list 然后分别操作。
	// user2project 需要加唯一性索引，在 db 层面保证记录不重复
	// alter table t_aa add unique index(aa,bb);
	list, err := model.GetUserListByProject(req.Id)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// req.List 和 list 做 diff
	// 先排序，后遍历，时间复杂度：O(nlgn+m+n)
	quickSort(req.List, 0, len(req.List)-1)

	var addList []uint32
	var delList []uint32

	flag := 0
	indexNow := 0
	indexNew := 0

	// req 小于 list -> req 往后走，写入 addList
	// req 大于 list -> list 往后走，写入 delList
	// 直到 req 等于 list -> 一起往后走，直到一方走完或者不等
	for indexNew < len(req.List) && indexNow < len(list) {
		new := req.List[indexNew]
		now := list[indexNow].ID
		if new < now && flag == 0 {
			addList = append(addList, new)
			indexNew++
		} else if new > now && flag == 0 {
			delList = append(delList, now)
			indexNow++
		} else if new == now {
			flag++
			indexNow++
			indexNew++
		} else {
			break
		}
	}

	// 写入剩余结果
	addList = append(addList, req.List[indexNew:]...)

	for indexNow < len(list) {
		delList = append(delList, list[indexNow].ID)
		indexNow++
	}

	// 事务分步执行
	err = model.UpdateMembers(m.DB.Self, req.Id, addList, delList)
	if err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	// 找 projectName
	var name string
	if name, err = model.GetProjectName(req.Id); err != nil {
		return e.ServerErr(errno.ErrDatabase, err.Error())
	}

	res.Name = name

	return nil
}

// 交换
func swap(m, n *uint32) {
	tmp := *m
	*m = *n
	*n = tmp
}

// 默认以最后一个元素作为中枢
// 比中枢大的放右边，比中枢小的放左边。一直递归求解即可。
func partition(arr []uint32, low, high int) int {
	pivot := arr[high]
	i := low - 1

	// 采用了双指针比较，一遍循环结束
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			swap(&arr[j], &arr[i])
		}
	}
	swap(&arr[i+1], &arr[high])
	return (i + 1)
}

func quickSort(arr []uint32, low, high int) {
	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}
