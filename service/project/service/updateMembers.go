package service

import (
	"context"
	pb "muxi-workbench-project/proto"
)

// UpdateMembers ... 更新项目成员列表
func (s *Service) UpdateMembers(ctx context.Context, req *pb.UpdateMemberRequest, res *pb.Response) error {

	// 这里需要 diff 老的项目成员列表，出 add list 和 delete list 然后分别操作。
	// user2project 需要加唯一性索引，在 db 层面保证记录不重复
	// alter table t_aa add unique index(aa,bb);
	return nil
}
